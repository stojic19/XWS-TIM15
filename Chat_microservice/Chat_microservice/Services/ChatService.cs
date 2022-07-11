using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection.Metadata.Ecma335;
using System.Threading.Tasks;
using AutoMapper;
using Chat_microservice.Configuration;
using Chat_microservice.model;
using Chat_microservice.Protos;
using Chat_microservice.Repository;
using Google.Protobuf.Collections;
using Grpc.Core;
using Grpc.Net.Client;
using Microsoft.AspNetCore.Authorization;
using OpenTracing;


namespace Chat_microservice.Services
{
    public class ChatService : ChatServiceGrpc.ChatServiceGrpcBase
    {
        private readonly IChatRepository _chatRepository;
        private readonly IMapper _mapper;
        private readonly ITracer _tracer;

        public ChatService(IChatRepository chatRepository, IMapper mapper, ITracer tracer)
        {
            _chatRepository = chatRepository;
            _mapper = mapper;
            _tracer = tracer;
        }

        public override Task<ChatsMsg> Get(GetRequest request, ServerCallContext context)
        {
            var scope = _tracer.BuildSpan("Get").StartActive(true);
            try
            {
                var chats = _chatRepository.GetAll();
                scope.Span.Finish();
                return Task.FromResult(_mapper.Map<ChatsMsg>(chats));
            }
            catch
            {
                scope.Span.Finish();
                throw;
            }
        }

        public override Task<ChatsMsg> GetForUser(IdMessage id, ServerCallContext context)
        {
            var scope = _tracer.BuildSpan("GetForUser").StartActive(true);
            try
            {
                var chats = _chatRepository.GetForUser(Guid.Parse(id.Id)).ToList();
                scope.Span.Finish();
                return Task.FromResult(_mapper.Map<ChatsMsg>(chats));
            }
            catch
            {
                scope.Span.Finish();
                throw;
            }
        }

        public override Task<ChatMsg> Add(NewMessage message, ServerCallContext context)
        {
            var metadata = context.RequestHeaders;
            var scope = _tracer.BuildSpan("Add").StartActive(true);
            metadata.Add(new Metadata.Entry("uber-trace-id", scope.Span.Context.ToString()));
            try
            {
                Authorize(context);
                var chat = _chatRepository.GetByParticipants(Guid.Parse(message.SenderId),
                    Guid.Parse(message.ReceiverId));
                if (chat == null)
                {
                    chat = _mapper.Map<Chat>(message);
                    _chatRepository.Add(chat);
                    SendNotification(message.ReceiverId, message.SenderId, chat, metadata);
                    scope.Span.Finish();
                    return Task.FromResult(_mapper.Map<ChatMsg>(chat));
                }

                var chatMessage = _mapper.Map<ChatMessage>(message);
                chat.Messages = chat.Messages.Append(chatMessage);
                _chatRepository.Update(chat);
                if(!chat.IsBlocked(message.ReceiverId))SendNotification(message.ReceiverId, message.SenderId, chat, metadata);
                scope.Span.Finish();
                return Task.FromResult(_mapper.Map<ChatMsg>(chat));
            }
            catch
            {
                scope.Span.Finish();
                throw;
            }
        }

        private static void SendNotification(string receiverId, string senderId, Chat chat, Metadata metadata)
        {
            var config = new EnvironmentConfiguration();
            var channel = GrpcChannel.ForAddress("http://" + config.NotificationsHost + ":" + config.NotificationsPort);
            var client = new NotificationsService.NotificationsServiceClient(channel);
            client.SaveNotificationAsync(new SaveNotificationRequest
            {
                Notification = new Notification
                {
                    UserId = receiverId,
                    MessagesId = chat.Id,
                    Time = chat.Messages.First().CreatedDate.ToShortTimeString(),
                    Type = "message",
                    FollowerId = senderId,
                    Action = "like"
                }
            }, metadata);
        }

        private void Authorize(ServerCallContext context)
        {
            string sub = context.RequestHeaders.Get("sub").Value;
            if (sub is null or "")
                throw new RpcException(new Status(StatusCode.Unauthenticated, "Unauthorized"));
        }
    }
}

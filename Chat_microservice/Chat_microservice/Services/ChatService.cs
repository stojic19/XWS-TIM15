using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection.Metadata.Ecma335;
using System.Threading.Tasks;
using AutoMapper;
using Chat_microservice.model;
using Chat_microservice.Protos;
using Chat_microservice.Repository;
using Google.Protobuf.Collections;
using Grpc.Core;
using Microsoft.AspNetCore.Authorization;


namespace Chat_microservice.Services
{
    public class ChatService : ChatServiceGrpc.ChatServiceGrpcBase
    {
        private readonly IChatRepository _chatRepository;
        private readonly IMapper _mapper;

        public ChatService(IChatRepository chatRepository, IMapper mapper)
        {
            _chatRepository = chatRepository;
            _mapper = mapper;
        }

        public override Task<ChatsMsg> Get(GetRequest request, ServerCallContext context)
        {
            var chats = _chatRepository.GetAll();
            return Task.FromResult(_mapper.Map<ChatsMsg>(chats));
        }

        public override Task<ChatsMsg> GetForUser(IdMessage id, ServerCallContext context)
        {
            var chats = _chatRepository.GetForUser(Guid.Parse(id.Id)).ToList();
            return Task.FromResult(_mapper.Map<ChatsMsg>(chats));
        }

        public override Task<ChatMsg> Add(NewMessage message, ServerCallContext context)
        {
            Authorize(context);
            var chat = _chatRepository.GetByParticipants(Guid.Parse(message.SenderId),
                Guid.Parse(message.ReceiverId));
            if (chat == null)
            {
                chat = _mapper.Map<Chat>(message);
                _chatRepository.Add(chat);
                return Task.FromResult(_mapper.Map<ChatMsg>(chat));
            }

            var chatMessage = _mapper.Map<ChatMessage>(message); 
            chat.Messages = chat.Messages.Append(chatMessage);
            _chatRepository.Update(chat);
            return Task.FromResult(_mapper.Map<ChatMsg>(chat));
        }

        private void Authorize(ServerCallContext context)
        {
            string sub = context.RequestHeaders.Get("sub").Value;
            if (sub is null or "")
                throw new RpcException(new Status(StatusCode.Unauthenticated, "Unauthorized"));
        }
    }
}

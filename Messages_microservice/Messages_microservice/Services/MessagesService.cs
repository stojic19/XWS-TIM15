using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection.Metadata.Ecma335;
using System.Threading.Tasks;
using Grpc.Core;
using Messages_microservice.model;
using Messages_microservice.Protos;
using Messages_microservice.Repository;
using Message = Messages_microservice.Protos.Message;

namespace Messages_microservice.Services
{
    public class MessagesService : Messages.MessagesBase
    {
        private readonly IMessageRepository _messageRepository;

        public MessagesService(IMessageRepository messageRepository)
        {
            _messageRepository = messageRepository;
        }

        public override Task<GetResponse> Get(GetRequest request, ServerCallContext context)
        {
            //throw new RpcException(new Status(StatusCode.Unauthenticated, "Unauthorized!"));
            /*var msgs = _messageRepository.GetAll();
            var retList = new List<ChatMessage>();
            foreach (var msg in msgs)retList.Add(new ChatMessage
            {
                Id = msg.Id.ToString(),
                Text = msg.Text
            });
            var retVal = new GetResponse
            {
                ChatMessage = { retList }
            };
            return Task.FromResult(retVal);*/
            return null;
        }

        public override Task<Message> Add(NewMessage message, ServerCallContext context)
        {
            var chat = _messageRepository.GetByParticipants(Guid.Parse(message.SenderId),
                Guid.Parse(message.ReceiverId));
            if (chat == null)
            {
                chat = new Chat
                {
                    FirstParticipant = new ChatParticipant
                    {
                        UserId = Guid.Parse(message.SenderId),
                        BlockedChat = false
                    },
                    SecondParticipant = new ChatParticipant
                    {
                        UserId = Guid.Parse(message.ReceiverId),
                        BlockedChat = false
                    },
                    Messages = new List<ChatMessage>{new()
                    {
                        CreatedDate = DateTime.Now,
                        Text = message.Text,
                        UserId = Guid.Parse(message.SenderId)
                    }}
                };
                _messageRepository.Add(chat);
                return Task.FromResult(new Message
                {
                    SenderId = chat.Messages.Last().UserId.ToString(),
                    ReceiverId = chat.FirstParticipant.UserId == Guid.Parse(message.ReceiverId) ? 
                        chat.SecondParticipant.UserId.ToString() : chat.FirstParticipant.UserId.ToString(),
                    Text = chat.Messages.Last().Text
                });
            }

            var chatMessage = new ChatMessage
            {
                CreatedDate = DateTime.Now,
                Text = message.Text,
                UserId = Guid.Parse(message.SenderId)
            }; 
            chat.Messages = chat.Messages.Append(chatMessage);
            _messageRepository.Update(chat);
            return Task.FromResult(new Message
            {
                SenderId = chat.Messages.Last().UserId.ToString(),
                ReceiverId = chat.FirstParticipant.UserId == Guid.Parse(message.ReceiverId) ?
                    chat.SecondParticipant.UserId.ToString() : chat.FirstParticipant.UserId.ToString(),
                Text = chat.Messages.Last().Text
            });
        }
    }
}

using System.Collections.Generic;
using System.Reflection.Metadata.Ecma335;
using System.Threading.Tasks;
using Grpc.Core;
using Messages_microservice.Protos;
using Messages_microservice.Repository;

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
            var msgs = _messageRepository.GetAll();
            var retList = new List<Message>();
            foreach (var msg in msgs)retList.Add(new Message
            {
                Id = msg.Id.ToString(),
                Text = msg.Text
            });
            var retVal = new GetResponse
            {
                Message = { retList }
            };
            return Task.FromResult(retVal);
        }

        public override Task<Message> Add(NewMessage message, ServerCallContext context)
        {
            var msg = _messageRepository.Add(new model.Message
            {
                Text = message.Text
            });
            return Task.FromResult(new Message
            {
                Id = msg.Id.ToString(),
                Text = msg.Text
            });
        }
    }
}

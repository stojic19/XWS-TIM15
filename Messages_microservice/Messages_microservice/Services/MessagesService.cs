using System.Reflection.Metadata.Ecma335;
using System.Threading.Tasks;
using Grpc.Core;
using Messages_microservice.Protos;

namespace Messages_microservice.Services
{
    public class MessagesService : Messages.MessagesBase
    {
        public override Task<GetResponse> Get(GetRequest request, ServerCallContext context)
        {
            return Task.FromResult(new GetResponse
            {
                Message = "Alo"
            });
        }
    }
}

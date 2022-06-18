using System.Collections;
using System.Collections.Generic;
using Messages_microservice.model;

namespace Messages_microservice.Repository
{
    public interface IMessageRepository
    {
        public IEnumerable<Message> GetAll();
        public Message Add(Message message);
    }
}

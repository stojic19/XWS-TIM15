using System;
using System.Collections;
using System.Collections.Generic;
using Messages_microservice.model;

namespace Messages_microservice.Repository
{
    public interface IMessageRepository
    {
        public IEnumerable<Chat> GetAll();
        public Chat GetByParticipants(Guid first, Guid second);
        public Chat Add(Chat chat);
        public Chat Update(Chat chat);
    }
}

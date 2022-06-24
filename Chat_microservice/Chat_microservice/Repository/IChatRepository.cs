using System;
using System.Collections;
using System.Collections.Generic;
using Chat_microservice.model;

namespace Chat_microservice.Repository
{
    public interface IChatRepository
    {
        public IEnumerable<Chat> GetAll();
        public Chat GetByParticipants(Guid first, Guid second);
        public IEnumerable<Chat> GetForUser(Guid userId);
        public Chat Add(Chat chat);
        public Chat Update(Chat chat);
    }
}

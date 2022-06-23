using System;
using Messages_microservice.model;
using MongoDB.Driver;
using System.Collections.Generic;
using System.Threading;
using MongoDB.Bson;

namespace Messages_microservice.Repository
{
    public class MessageRepository : IMessageRepository
    {
        private readonly IMongoCollection<Chat> _chats;

        public MessageRepository()
        {
            var mongoClient = new MongoClient("mongodb://" +
                                              Environment.GetEnvironmentVariable("MESSAGES_DB_HOST") + ":" +
                                              Environment.GetEnvironmentVariable("MESSAGES_DB_PORT"));
            var mongoDatabase = mongoClient.GetDatabase("chats");
            _chats = mongoDatabase.GetCollection<Chat>("chats");
        }

        public IEnumerable<Chat> GetAll()
        { 
            return _chats.Find(_ => true).ToList();
        }

        public Chat GetByParticipants(Guid first, Guid second)
        {
            var filter11 = Builders<Chat>.Filter.Eq(c => c.FirstParticipant.UserId, first);
            var filter12 = Builders<Chat>.Filter.Eq(c => c.SecondParticipant.UserId, second);

            var filter21 = Builders<Chat>.Filter.Eq(c => c.FirstParticipant.UserId, second);
            var filter22 = Builders<Chat>.Filter.Eq(c => c.SecondParticipant.UserId, first);

            var filter1 = Builders<Chat>.Filter.And(filter11, filter12);
            var filter2 = Builders<Chat>.Filter.And(filter21, filter22);

            var chat = _chats.Find(Builders<Chat>.Filter.Or(filter1, filter2)).FirstOrDefault();
            return chat;
        }

        public Chat Add(Chat chat)
        { 
            _chats.InsertOne(chat);
            return chat;
        }

        public Chat Update(Chat chat)
        {
            var filter = Builders<Chat>.Filter.Eq(c => c.Id, chat.Id);
            _chats.ReplaceOne(filter, chat);
            return chat;
        }
    }
}

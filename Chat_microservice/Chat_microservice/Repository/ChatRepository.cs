using System;
using Chat_microservice.model;
using MongoDB.Driver;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using Chat_microservice.Configuration;
using MongoDB.Bson;

namespace Chat_microservice.Repository
{
    public class ChatRepository : IChatRepository
    {
        private readonly IMongoCollection<Chat> _chats;

        public ChatRepository()
        {
            var cfg = new EnvironmentConfiguration();
            var mongoClient = new MongoClient("mongodb://" +
                                              cfg.ChatDbHost + ":" +
                                              cfg.ChatDbPort);
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

        public IEnumerable<Chat> GetForUser(Guid userId)
        {
            var filter11 = Builders<Chat>.Filter.Eq(c => c.FirstParticipant.UserId, userId);
            var filter12 = Builders<Chat>.Filter.Eq(c => c.FirstParticipant.BlockedChat, false);

            var filter21 = Builders<Chat>.Filter.Eq(c => c.SecondParticipant.UserId, userId);
            var filter22 = Builders<Chat>.Filter.Eq(c => c.SecondParticipant.BlockedChat, false);

            var filter1 = Builders<Chat>.Filter.And(filter11, filter12);
            var filter2 = Builders<Chat>.Filter.And(filter21, filter22);

            var filter3 = Builders<Chat>.Filter.SizeGt(c => c.Messages, 0);

            var filter4 = Builders<Chat>.Filter.Or(filter1, filter2);

            return _chats.Find(Builders<Chat>.Filter.And(filter3, filter4)).ToEnumerable();
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

        public Chat Delete(Chat chat)
        {
            var filter = Builders<Chat>.Filter.Eq(c => c.Id, chat.Id);
            _chats.DeleteOne(filter);
            return chat;
        }
    }
}

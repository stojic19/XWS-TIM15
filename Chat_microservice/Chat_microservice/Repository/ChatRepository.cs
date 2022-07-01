using System;
using Chat_microservice.model;
using MongoDB.Driver;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using Chat_microservice.Configuration;
using MongoDB.Bson;
using OpenTracing;

namespace Chat_microservice.Repository
{
    public class ChatRepository : IChatRepository
    {
        private readonly ITracer _tracer;

        private readonly IMongoCollection<Chat> _chats;

        public ChatRepository(ITracer tracer)
        {
            _tracer = tracer;
            var cfg = new EnvironmentConfiguration();
            var mongoClient = new MongoClient("mongodb://" +
                                              cfg.ChatDbHost + ":" +
                                              cfg.ChatDbPort);
            var mongoDatabase = mongoClient.GetDatabase("chats");
            _chats = mongoDatabase.GetCollection<Chat>("chats");
        }

        public IEnumerable<Chat> GetAll()
        {
            var span1 = _tracer.BuildSpan("MongoGetAll").Start();
            IEnumerable<Chat> retVal;
            try
            {
                retVal = _chats.Find(_ => true).ToList();
            }
            catch
            {
                span1.Finish();
                throw;
            }
            span1.Finish();
            return retVal;
        }

        public Chat GetByParticipants(Guid first, Guid second)
        {
            var filter11 = Builders<Chat>.Filter.Eq(c => c.FirstParticipant.UserId, first);
            var filter12 = Builders<Chat>.Filter.Eq(c => c.SecondParticipant.UserId, second);

            var filter21 = Builders<Chat>.Filter.Eq(c => c.FirstParticipant.UserId, second);
            var filter22 = Builders<Chat>.Filter.Eq(c => c.SecondParticipant.UserId, first);

            var filter1 = Builders<Chat>.Filter.And(filter11, filter12);
            var filter2 = Builders<Chat>.Filter.And(filter21, filter22);
            Chat chat;
            var span1 = _tracer.BuildSpan("MongoGetByParticipants").Start();
            try
            {
                chat = _chats.Find(Builders<Chat>.Filter.Or(filter1, filter2)).FirstOrDefault();
            }
            catch
            {
                span1.Finish();
                throw;
            }
            span1.Finish();
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

            var span1 = _tracer.BuildSpan("MongoGetForUser").Start();
            IEnumerable<Chat> chats;
            try
            {
                chats = _chats.Find(Builders<Chat>.Filter.And(filter3, filter4)).ToEnumerable();
            }
            catch
            {
                span1.Finish();
                throw;
            }
            span1.Finish();
            return chats;
        }

        public Chat Add(Chat chat)
        {
            var span1 = _tracer.BuildSpan("MongoAdd").Start();
            try
            {
                _chats.InsertOne(chat);
            }
            catch
            {
                span1.Finish();
                throw;
            }
            span1.Finish();
            return chat;
        }

        public Chat Update(Chat chat)
        {
            var filter = Builders<Chat>.Filter.Eq(c => c.Id, chat.Id);
            var span1 = _tracer.BuildSpan("MongoUpdate").Start();
            try
            {
                _chats.ReplaceOne(filter, chat);
            }
            catch
            {
                span1.Finish();
                throw;
            }
            span1.Finish();
            return chat;
        }

        public Chat Delete(Chat chat)
        {
            var filter = Builders<Chat>.Filter.Eq(c => c.Id, chat.Id);
            var span1 = _tracer.BuildSpan("MongoDelete").Start();
            try
            {
                _chats.DeleteOne(filter);
            }
            catch
            {
                span1.Finish();
                throw;
            }
            span1.Finish();
            return chat;
        }
    }
}

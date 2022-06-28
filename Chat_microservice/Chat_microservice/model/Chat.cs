using System;
using System.Collections;
using System.Collections.Generic;
using Chat_microservice.Nats.Messages;
using Microsoft.Extensions.Configuration;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace Chat_microservice.model
{
    public class Chat
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }
        public ChatParticipant FirstParticipant { get; set; }
        public ChatParticipant SecondParticipant { get; set; }
        public IEnumerable<ChatMessage> Messages { get; set; }

        public bool SetToBlocked(Guid id)
        {
            if (FirstParticipant.UserId == id)
            {
                FirstParticipant.BlockedChat = true;
                return true;
            }

            if (SecondParticipant.UserId == id)
            {
                SecondParticipant.BlockedChat = true;
                return true;
            }
            return false;
        }

        public bool SetToUnblocked(Guid id)
        {
            if (FirstParticipant.UserId == id)
            {
                FirstParticipant.BlockedChat = false;
                return true;
            }

            if (SecondParticipant.UserId == id)
            {
                SecondParticipant.BlockedChat = false;
                return true;
            }
            return false;
        }
    }
}

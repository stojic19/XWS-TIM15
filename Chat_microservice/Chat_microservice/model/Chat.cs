using System;
using System.Collections;
using System.Collections.Generic;
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
    }
}

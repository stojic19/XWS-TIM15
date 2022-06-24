using System;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace Chat_microservice.model
{
    public class ChatMessage
    {
        public Guid UserId { get; set; }
        public string Text { get; set; }
        public DateTime CreatedDate { get; set; }

    }
}

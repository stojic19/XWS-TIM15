using System;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace Messages_microservice.model
{
    public class Message
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }
        public string Text { get; set; }
    }
}

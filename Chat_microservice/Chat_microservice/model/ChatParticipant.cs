using System;

namespace Chat_microservice.model
{
    public class ChatParticipant
    {
        public Guid UserId { get; set; }
        public bool BlockedChat { get; set; }
    }
}

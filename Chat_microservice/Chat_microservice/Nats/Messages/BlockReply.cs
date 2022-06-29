namespace Chat_microservice.Nats.Messages
{
    public class BlockReply
    {
        public string BlockerId { get; set; }
        public string BlockedId { get; set; }
        public BlockReplyType Type { get; set; }

    }
    public enum BlockReplyType
    {
        ChatBlocked,
        ChatNotBlocked,
        UserBlocked,
        UserNotBlocked,
        UnknownReply
    }

}

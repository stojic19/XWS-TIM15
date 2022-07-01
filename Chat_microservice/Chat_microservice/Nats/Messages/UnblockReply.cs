namespace Chat_microservice.Nats.Messages
{
    public class UnblockReply
    {
        public string BlockerId { get; set; }
        public string BlockedId { get; set; }
        public UnblockReplyType Type { get; set; }
    }
    public enum UnblockReplyType
    {
        ChatUnblocked,
        ChatNotUnblocked,
        UserUnblocked,
        UserNotUnblocked,
        UnknownUnblockReply
    }
}

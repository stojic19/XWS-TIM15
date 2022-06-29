namespace Chat_microservice.Nats.Messages
{
    public class UnblockCommand
    {
        public string BlockerId { get; set; }
        public string BlockedId { get; set; }
        public UnblockCommandType Type { get; set; }
        public bool IsRelevant() => Type == UnblockCommandType.UnblockChat;
    }

    public enum UnblockCommandType
    {
        UnblockChat,
        ConfirmUnblock,
        RevertUnblock,
        UnknownUnblockCommand
    }
}

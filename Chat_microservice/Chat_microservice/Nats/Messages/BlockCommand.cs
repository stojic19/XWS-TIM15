using Google.Protobuf.Reflection;

namespace Chat_microservice.Nats.Messages
{
    public class BlockCommand
    {
        public string BlockerId { get; set; }
        public string BlockedId { get; set; }
        public BlockCommandType Type { get; set; }
        public bool IsRelevant() => Type == BlockCommandType.BlockChat;

    }

    

    public enum BlockCommandType
    {
        BlockChat,
        ConfirmBlock,
        RevertBlock,
        UnknownCommand
    }
}

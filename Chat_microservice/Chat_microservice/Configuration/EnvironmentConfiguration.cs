using System;

namespace Chat_microservice.Configuration
{
    public class EnvironmentConfiguration
    {
        public string Port { get; set; }
        public string NatsHost { get; set; }
        public string NatsPort { get; set; }
        public string NatsUser { get; set; }
        public string NatsPass { get; set; }
        public string BlockCommandSubject { get; set; }
        public string BlockReplySubject { get; set; }
        public string UnblockCommandSubject { get; set; }
        public string UnblockReplySubject { get; set; }
        public string QueueName { get; set; }

        public EnvironmentConfiguration()
        {
            Port = GetEnvironmentVarOrDefault("CHAT_PORT", "8004");
            NatsHost = GetEnvironmentVarOrDefault("NATS_HOST", "localhost");
            NatsPort = GetEnvironmentVarOrDefault("NATS_PORT", "4222");
            NatsUser = GetEnvironmentVarOrDefault("NATS_USER", "ruser");
            NatsPass = GetEnvironmentVarOrDefault("NATS_PASS", "T0pS3cr3t");
            BlockCommandSubject = GetEnvironmentVarOrDefault("BLOCK_COMMAND_SUBJECT", "block.command");
            BlockReplySubject = GetEnvironmentVarOrDefault("BLOCK_REPLY_SUBJECT", "block.reply");
            UnblockCommandSubject = GetEnvironmentVarOrDefault("UNBLOCK_COMMAND_SUBJECT", "unblock.command");
            UnblockReplySubject = GetEnvironmentVarOrDefault("UNBLOCK_REPLY_SUBJECT", "unblock.reply");
            QueueName = "chat_service";
        }

        private string GetEnvironmentVarOrDefault(string var, string def) => Environment.GetEnvironmentVariable(var) ?? def;
    }
}

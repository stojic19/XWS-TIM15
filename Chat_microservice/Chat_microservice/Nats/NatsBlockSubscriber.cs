using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using AutoMapper;
using Chat_microservice.Configuration;
using Chat_microservice.model;
using Chat_microservice.Nats.Messages;
using Chat_microservice.Repository;
using Chat_microservice.Utilities;
using Microsoft.AspNetCore.Server.IIS.Core;
using Microsoft.Extensions.Hosting;
using NATS.Client;

namespace Chat_microservice.Nats
{
    public class NatsBlockSubscriber : BackgroundService
    {
        private IAsyncSubscription _subscription;
        private IConnection _connection;
        private IChatRepository _chatRepository;
        private IMapper _mapper;
        private readonly EnvironmentConfiguration _config;
        public NatsBlockSubscriber(IChatRepository chatRepository, IMapper mapper)
        {
            _mapper = mapper;
            _chatRepository = chatRepository;
            ConnectionFactory cf = new ConnectionFactory();
            var opts = ConnectionFactory.GetDefaultOptions();
            _config = new EnvironmentConfiguration();
            opts.Password = _config.NatsPass;
            opts.User = _config.NatsUser;
            opts.Url = "nats://" + _config.NatsHost + ":" + _config.NatsPort;
            _connection = cf.CreateConnection(opts);
            _subscription = _connection.SubscribeAsync(_config.BlockCommandSubject, _config.QueueName);
        }

        protected override Task ExecuteAsync(CancellationToken stoppingToken)
        {
            EventHandler<MsgHandlerEventArgs> h = (e, args) =>
            {
                try
                {
                    var command = ConversionUtilities.DeserializeBinary<BlockCommand>(args.Message.Data);
                    if (!command.IsRelevant()) return;
                    var chat = _chatRepository.GetByParticipants(Guid.Parse(command.BlockerId), Guid.Parse(command.BlockedId));
                    if (chat == null)
                    {
                        chat = new Chat
                        {
                            FirstParticipant = new ChatParticipant
                                { BlockedChat = true, UserId = Guid.Parse(command.BlockerId) },
                            SecondParticipant = new ChatParticipant
                                { BlockedChat = false, UserId = Guid.Parse(command.BlockedId) },
                            Messages = new List<ChatMessage>()
                        };
                        _chatRepository.Add(chat);
                        var notBlockedBlockReply = _mapper.Map<BlockReply>(command);
                        notBlockedBlockReply.Type = BlockReplyType.ChatBlocked;
                        Publish(_config.BlockReplySubject, ConversionUtilities.SerializeBinary(notBlockedBlockReply));
                        return;
                    }
                    chat.SetToBlocked(Guid.Parse(command.BlockerId));
                    _chatRepository.Update(chat);
                    var blockedReply = _mapper.Map<BlockReply>(command);
                    blockedReply.Type = BlockReplyType.ChatBlocked;
                    Publish(_config.BlockReplySubject, ConversionUtilities.SerializeBinary(blockedReply));
                }
                catch
                {
                    var command = ConversionUtilities.DeserializeBinary<BlockCommand>(args.Message.Data);
                    var notBlockedBlockReply = _mapper.Map<BlockReply>(command);
                    notBlockedBlockReply.Type = BlockReplyType.ChatNotBlocked;
                    Publish(_config.BlockReplySubject, ConversionUtilities.SerializeBinary(notBlockedBlockReply));
                }
                
            };
            _subscription.MessageHandler += h;
            _subscription.Start();
            return Task.CompletedTask;
        }

        private void Publish(string subject, byte[] data) => _connection.Publish(subject, data);

        public override void Dispose()
        {
            _subscription.Unsubscribe();
            _subscription.Dispose();
            _connection.Dispose();
            base.Dispose();
        }
    }
}

using System;
using System.Threading;
using System.Threading.Tasks;
using AutoMapper;
using Chat_microservice.Configuration;
using Chat_microservice.Repository;
using Microsoft.Extensions.Hosting;
using NATS.Client;

namespace Chat_microservice.Nats
{
    public class NatsUnblockSubscriber : BackgroundService
    {
        private IAsyncSubscription _subscription;
        private IConnection _connection;
        private IChatRepository _chatRepository;
        private IMapper _mapper;
        private readonly EnvironmentConfiguration _config;
        public NatsUnblockSubscriber(IChatRepository chatRepository, IMapper mapper)
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
            _subscription = _connection.SubscribeAsync(_config.UnblockCommandSubject, _config.QueueName);
        }
        protected override Task ExecuteAsync(CancellationToken stoppingToken)
        {
            EventHandler<MsgHandlerEventArgs> h = (e, args) =>
            {

            };
            _subscription.MessageHandler += h;
            _subscription.Start();
            return Task.CompletedTask;
        }
        public override void Dispose()
        {
            _subscription.Unsubscribe();
            _subscription.Dispose();
            _connection.Dispose();
            base.Dispose();
        }
    }
}

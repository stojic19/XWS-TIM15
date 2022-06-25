using System;
using System.Diagnostics;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Hosting;
using NATS.Client;

namespace Chat_microservice.Nats
{
    public class NatsBlockSubscriber : BackgroundService
    {
        private IAsyncSubscription _subscription;
        public NatsBlockSubscriber()
        {
            ConnectionFactory cf = new ConnectionFactory();
            var opts = ConnectionFactory.GetDefaultOptions();
            opts.Password = "T0pS3cr3t";
            opts.User = "ruser";
            opts.Url = "nats://localhost:4222";
            IConnection c = cf.CreateConnection(opts);
            _subscription = c.SubscribeAsync("chat_block");
        }

        protected override Task ExecuteAsync(CancellationToken stoppingToken)
        {
            EventHandler<MsgHandlerEventArgs> h = (e, args) =>
            {
                Console.WriteLine("EVO ME TU SAM");
            };
            _subscription.MessageHandler += h;
            _subscription.Start();
            return Task.CompletedTask;
        }

        public override void Dispose()
        {
            _subscription.Unsubscribe();
            _subscription.Dispose();
            base.Dispose();
        }
    }
}

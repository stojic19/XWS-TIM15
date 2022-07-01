using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Hosting;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Chat_microservice.Configuration;
using Microsoft.AspNetCore;

namespace Chat_microservice
{
    public class Program
    {
        public static void Main(string[] args)
        {
            CreateHostBuilder(args).Run();
        }

        public static IWebHost CreateHostBuilder(string[] args) =>
            WebHost.CreateDefaultBuilder(args)
                .UseUrls(@"http://0.0.0.0:" + new EnvironmentConfiguration().Port)
                .UseStartup<Startup>()
                .Build();
    }
}

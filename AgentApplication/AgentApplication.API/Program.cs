using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.AspNetCore;
using Microsoft.AspNetCore.Hosting;

namespace AgentApplication.API
{
    public class Program
    {
        public static void Main(string[] args)
        {
            CreateHostBuilder(args).Run();
        }

        public static IWebHost CreateHostBuilder(string[] args) =>
            WebHost.CreateDefaultBuilder(args)
                .UseUrls("http://localhost:9000")//http:// + Environment.GetEnvironmentVariable("AGENT_APPLICATION_HOST") + : + Environment.GetEnvironmentVariable("AGENT_APPLICATION_PORT");
                .UseStartup<Startup>()
                .Build();
    }
}

using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using System.Threading.Tasks;
using Autofac;
using Autofac.Extensions.DependencyInjection;
using AutoMapper.Contrib.Autofac.DependencyInjection;
using Chat_microservice.AutoMapperProfiles;
using Chat_microservice.Nats;
using Chat_microservice.Repository;
using Chat_microservice.Services;
using Microsoft.AspNetCore.Authentication;

namespace Chat_microservice
{
    public class Startup
    {
        public IServiceProvider ConfigureServices(IServiceCollection services)
        {
            services.AddHostedService<NatsBlockSubscriber>();
            services.AddHostedService<NatsUnblockSubscriber>();
            services.AddGrpc();
            List<Assembly> assemblies = new List<Assembly> { typeof(ChatProfile).Assembly };
            var containerBuilder = new ContainerBuilder();
            containerBuilder.RegisterAutoMapper(propertiesAutowired: false, assemblies.ToArray());
            containerBuilder.RegisterType<ChatRepository>().As<IChatRepository>();
            containerBuilder.Populate(services);
            var container = containerBuilder.Build();
            return new AutofacServiceProvider(container);
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            app.UseRouting();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapGrpcService<ChatService>();
            });
        }
    }
}

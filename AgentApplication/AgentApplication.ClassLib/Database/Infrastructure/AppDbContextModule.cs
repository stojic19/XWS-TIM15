using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Database.EfStructures;
using Autofac;

namespace AgentApplication.ClassLib.Database.Infrastructure
{
    public class AppDbContextModule : Module
    {
        protected override void Load(ContainerBuilder builder)
        {
            builder.RegisterType<AppDbContext>()
                .WithParameter("options", AppDbContextFactory.GetOptions())
                .InstancePerLifetimeScope();
        }
    }
}

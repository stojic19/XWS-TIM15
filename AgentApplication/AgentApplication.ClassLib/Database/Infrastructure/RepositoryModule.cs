using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Threading.Tasks;
using Autofac;

namespace AgentApplication.ClassLib.Database.Infrastructure
{
    public class RepositoryModule : Autofac.Module
    {
        public List<Assembly> RepositoryAssemblies { get; set; }
        public string Namespace { get; set; }

        protected override void Load(ContainerBuilder builder)
        {
            builder.RegisterAssemblyTypes(RepositoryAssemblies.ToArray())
                .Where(x => x.Namespace.Contains(Namespace))
                .AsImplementedInterfaces();
        }
    }
}

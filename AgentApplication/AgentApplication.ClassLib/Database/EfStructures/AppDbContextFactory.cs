using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Design;

namespace AgentApplication.ClassLib.Database.EfStructures
{
    public class AppDbContextFactory : IDesignTimeDbContextFactory<AppDbContext>
    {
        public AppDbContext CreateDbContext(string[] args)
        {
            var optionsBuilder = SetupOptions();
            return new AppDbContext(optionsBuilder.Options);
        }
        public static DbContextOptions<AppDbContext> GetOptions()
        {
            var optionsBuilder = SetupOptions();
            return optionsBuilder.Options;
        }

        private static DbContextOptionsBuilder<AppDbContext> SetupOptions()
        {
            var optionsBuilder = new DbContextOptionsBuilder<AppDbContext>();
            var connectionString = Environment.GetEnvironmentVariable("AGENT_APPLICATION_DB_CONNECTION_STRING");
            optionsBuilder.UseNpgsql(connectionString);
            return optionsBuilder;
        }
    }
}

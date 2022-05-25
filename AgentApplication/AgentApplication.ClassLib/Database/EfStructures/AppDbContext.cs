using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Model;
using Microsoft.EntityFrameworkCore;

namespace AgentApplication.ClassLib.Database.EfStructures
{
    public class AppDbContext : DbContext
    {
        public DbSet<Company> Companies { get; set; }

        public AppDbContext(DbContextOptions<AppDbContext> options) : base(options)
        {

        }

    }
}

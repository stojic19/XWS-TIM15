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
        public DbSet<User> Users { get; set; }

        public AppDbContext(DbContextOptions<AppDbContext> options) : base(options)
        {
            
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Company>().OwnsMany(t => t.Comments);
            modelBuilder.Entity<Company>().OwnsMany(t => t.Grades);
            modelBuilder.Entity<Company>().OwnsMany(t => t.JobOffers);
            
            modelBuilder.Entity<User>().OwnsOne(t => t.PersonalInfo);
            base.OnModelCreating(modelBuilder);
        }

    }
}

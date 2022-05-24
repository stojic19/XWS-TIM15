using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Database.EfStructures;
using Microsoft.EntityFrameworkCore;

namespace AgentApplication.ClassLib.Database.Repository.Base
{
    public class BaseReadRepository<TKey, TEntity> : IBaseReadRepository<TKey, TEntity> where TEntity : class, new()
    {
        private readonly AppDbContext _context;

        protected BaseReadRepository(AppDbContext context)
        {
            _context = context;
        }

        public TEntity GetById(TKey id)
        {
            return _context.Set<TEntity>().Find(id);
        }

        public DbSet<TEntity> GetAll()
        {
            return _context.Set<TEntity>();
        }
    }
}

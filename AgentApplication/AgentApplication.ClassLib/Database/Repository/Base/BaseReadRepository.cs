using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Database.EfStructures;
using AgentApplication.ClassLib.Database.Repository.Enums;
using AgentApplication.ClassLib.Model;
using Microsoft.EntityFrameworkCore;

namespace AgentApplication.ClassLib.Database.Repository.Base
{
    public class BaseReadRepository<TKey, TEntity> : IBaseReadRepository<TKey, TEntity> where TEntity : PersistentEntity, new()
    {
        private readonly AppDbContext _context;

        protected BaseReadRepository(AppDbContext context)
        {
            _context = context;
        }

        public virtual TEntity GetById(TKey id, FetchType fetchType = FetchType.Lazy)
        {
            var set = GetSet();
            IQueryable<TEntity> query = null;
            if (fetchType == FetchType.Eager)
            {
                var properties = typeof(TEntity).GetProperties();
                properties.ToList().ForEach(p =>
                {
                    if (p.PropertyType.BaseType.Equals(typeof(PersistentEntity)))
                    {
                        query = set.Include(p.Name);
                    }
                });
            }
            return set.Find(id);
        }

        public virtual DbSet<TEntity> GetAll(FetchType fetchType = FetchType.Lazy)
        {
            return GetSet();
        }

        protected DbSet<TEntity> GetSet()
        {
            return _context.Set<TEntity>();
        }
    }
}

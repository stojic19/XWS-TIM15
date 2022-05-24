using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;

namespace AgentApplication.ClassLib.Database.Repository.Base
{
    public interface IBaseReadRepository<TKey, TEntity> where TEntity : class
    {
        TEntity GetById(TKey id);
        DbSet<TEntity> GetAll();
    }
}

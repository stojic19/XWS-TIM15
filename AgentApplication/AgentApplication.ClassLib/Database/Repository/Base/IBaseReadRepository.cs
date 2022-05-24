using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Database.Repository.Enums;
using AgentApplication.ClassLib.Model;
using Microsoft.EntityFrameworkCore;

namespace AgentApplication.ClassLib.Database.Repository.Base
{
    public interface IBaseReadRepository<TKey, TEntity> where TEntity : PersistentEntity
    {
        TEntity GetById(TKey id, FetchType type = FetchType.Lazy);
        DbSet<TEntity> GetAll();
    }
}

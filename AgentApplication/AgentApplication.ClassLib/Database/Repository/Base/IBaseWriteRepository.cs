using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Model;

namespace AgentApplication.ClassLib.Database.Repository.Base
{
    public interface IBaseWriteRepository<TEntity> where TEntity : PersistentEntity
    {
        TEntity Add(TEntity entity, bool persist = true);
        TEntity Update(TEntity entity, bool persist = true);
        void Delete(TEntity entity, bool persist = true);
    }
}

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Database.Repository.Base;
using AgentApplication.ClassLib.Database.Repository.Enums;
using AgentApplication.ClassLib.Model;

namespace AgentApplication.ClassLib.Database.Repository
{
    public interface IUserReadRepository : IBaseReadRepository<Guid, User>
    {
        public User GetByUsername(string username);
    }
}

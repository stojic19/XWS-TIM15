using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Database.Repository.Base;
using AgentApplication.ClassLib.Database.Repository.Enums;
using AgentApplication.ClassLib.Model;
using Microsoft.EntityFrameworkCore;

namespace AgentApplication.ClassLib.Database.Repository
{
    public interface ICompanyReadRepository : IBaseReadRepository<Guid, Company>
    {
        IQueryable<Company> GetFromUser(Guid userId, FetchType type = FetchType.Lazy);
        IQueryable<Company> GetRegistered(FetchType type = FetchType.Lazy);
        IQueryable<Company> GetNotRegistered(FetchType type = FetchType.Lazy);
    }
}

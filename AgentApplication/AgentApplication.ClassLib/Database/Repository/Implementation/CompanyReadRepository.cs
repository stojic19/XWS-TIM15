using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Database.EfStructures;
using AgentApplication.ClassLib.Database.Repository.Base;
using AgentApplication.ClassLib.Database.Repository.Enums;
using AgentApplication.ClassLib.Model;
using Microsoft.EntityFrameworkCore;

namespace AgentApplication.ClassLib.Database.Repository.Implementation
{
    public class CompanyReadRepository : BaseReadRepository<Guid, Company>, ICompanyReadRepository
    {
        public CompanyReadRepository(AppDbContext context) : base(context)
        {
        }

        public override Company GetById(Guid id, FetchType fetchType = FetchType.Lazy)
        {
            var set = GetAll();
            if (fetchType == FetchType.Eager)
            {
                return set.Include(c => c.Grades).ThenInclude(c => c.User)
                    .Include(c => c.Comments).ThenInclude(c => c.User)
                    .Include(c => c.JobOffers)
                    .Include(c => c.Owner)
                    .FirstOrDefault(c => c.Id == id);
            }
            return set.Find(id);
        }
    }
}

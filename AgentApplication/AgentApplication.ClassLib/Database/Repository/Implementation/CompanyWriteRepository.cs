using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Database.EfStructures;
using AgentApplication.ClassLib.Database.Repository.Base;
using AgentApplication.ClassLib.Model;

namespace AgentApplication.ClassLib.Database.Repository.Implementation
{
    public class CompanyWriteRepository : BaseWriteRepository<Company>, ICompanyWriteRepository
    {
        public CompanyWriteRepository(AppDbContext context) : base(context)
        {
        }
    }
}

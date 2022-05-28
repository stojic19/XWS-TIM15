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
    public class JobOfferWriteRepository : BaseWriteRepository<JobOffer>, IJobOfferWriteRepository
    {
        public JobOfferWriteRepository(AppDbContext context) : base(context)
        {
        }
    }
}

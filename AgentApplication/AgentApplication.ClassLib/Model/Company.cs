using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Npgsql.EntityFrameworkCore.PostgreSQL.Infrastructure.Internal;

namespace AgentApplication.ClassLib.Model
{
    public class Company : PersistentEntity
    {
        public CompanyInfo CompanyInfo { get; set; }
        public bool Registered { get; set; }
        public DateTime TimeOfCreation { get; set; }
        public Guid OwnerId { get; set; }
        public User Owner { get; set; }
        public List<Comment> Comments { get; set; }
        public List<Grade> Grades { get; set; }
        public List<JobOffer> JobOffers { get; set; }
    }
}

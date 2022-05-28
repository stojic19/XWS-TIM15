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
        public string Name { get; set; }
        public string Address { get; set; }
        public string Email { get; set; }
        public string PhoneNumber { get; set; }
        public string Description { get; set; }
        public string Culture { get; set; }
        public bool Registered { get; set; }
        public DateTime TimeOfCreation { get; set; }
        public User Owner { get; set; }
    }
}

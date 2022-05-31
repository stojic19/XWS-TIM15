using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AgentApplication.ClassLib.Model
{
    public class JobOffer : PersistentEntity
    {
        public DateTime TimeOfCreation { get; set; }
        public string Position { get; set; }
        public string Description { get; set; }
        public string Requirements { get; set; }
        public bool IsActive { get; set; }
    }
}

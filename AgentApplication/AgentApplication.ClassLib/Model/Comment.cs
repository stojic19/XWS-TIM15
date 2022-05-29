using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AgentApplication.ClassLib.Model
{
    public class Comment : PersistentEntity
    {
        public Guid UserId { get; set; }
        public User User { get; set; }
        public string Content { get; set; }
        public DateTime TimeOfCreation { get; set; }
    }
}

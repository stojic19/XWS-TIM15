using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AgentApplication.ClassLib.Model
{
    public class Grade : PersistentEntity
    {
        public Grade()
        {
            TimeOfCreation = DateTime.Now;
        }
        public Guid UserId { get; set; }
        public User User { get; set; }
        public int Value { get; set; }
        public DateTime TimeOfCreation { get; set; }
    }
}

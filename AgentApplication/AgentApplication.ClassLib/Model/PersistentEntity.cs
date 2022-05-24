using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AgentApplication.ClassLib.Model
{
    public abstract class PersistentEntity
    {
        public Guid Id { get; set; }
    }
}

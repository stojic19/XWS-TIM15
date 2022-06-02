using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Model;

namespace AgentApplication.ClassLib.Service
{
    public interface IJwtGenerator
    {
        public string GenerateToken(User user);
    }
}

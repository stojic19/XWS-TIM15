using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Model;

namespace AgentApplication.ClassLib.Service
{
    public interface IAuthenticationService
    {
        // Summary:
        // Exceptions:
        //   T:AgentApplication.ClassLib.Exceptions.RegistrationException:
        //     username already exists.
        public void Register(User user);
        public string LogIn(string username, string password);
    }
}

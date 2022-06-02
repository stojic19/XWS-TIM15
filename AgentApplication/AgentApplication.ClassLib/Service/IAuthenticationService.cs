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
        /// <Summary>
        ///     Registers new users if there are no existing users with given username
        /// </Summary>
        /// <exception cref="AgentApplication.ClassLib.Exceptions.RegistrationException">User name exists</exception>
        public void Register(User user);
        public string LogIn(string username, string password);
    }
}

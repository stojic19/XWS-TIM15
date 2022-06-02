using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Threading.Tasks;
using AgentApplication.ClassLib.Database.Infrastructure;
using AgentApplication.ClassLib.Database.Repository;
using AgentApplication.ClassLib.Exceptions;
using AgentApplication.ClassLib.Model;

namespace AgentApplication.ClassLib.Service.Impl
{
    public class AuthenticationService : IAuthenticationService
    {
        private readonly IUnitOfWork _uow;
        private readonly IJwtGenerator _jwtGenerator;

        public AuthenticationService(IUnitOfWork uow, IJwtGenerator jwtGenerator)
        {
            _uow = uow;
            _jwtGenerator = jwtGenerator;
        }


        /// <Summary>
        ///     Registers new users if there are no existing users with given username
        /// </Summary>
        /// <exception cref="AgentApplication.ClassLib.Exceptions.RegistrationException">User name exists</exception>
        public void Register(User user)
        {
            if (_uow.GetRepository<IUserReadRepository>().GetByUsername(user.Username) != null)
                throw new RegistrationException("Username already exists");
            user.Salt = Encoder.CreateSalt(16);
            user.Password = Encoder.EncodePassword(user.Password, user.Salt);
            _uow.GetRepository<IUserWriteRepository>().Add(user);
        }
        /// <Summary>
        ///     Logs in user, returns jwt
        /// </Summary>
        /// <exception cref="AgentApplication.ClassLib.Exceptions.LogInException">User name exists</exception>
        public string LogIn(string username, string password)
        {
            var user = _uow.GetRepository<IUserReadRepository>().GetByUsername(username);
            if (user == null) throw new LogInException("User with given username not found!");
            if (user.Password != Encoder.EncodePassword(password, user.Salt)) throw new LogInException("Invalid password!");
            return _jwtGenerator.GenerateToken(user);
        }
    }
}

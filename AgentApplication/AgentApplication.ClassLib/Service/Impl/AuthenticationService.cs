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


        // Exceptions:
        //   T:AgentApplication.ClassLib.Exceptions.RegistrationException:
        //     username exists.
        public void Register(User user)
        {
            if (_uow.GetRepository<IUserReadRepository>().GetByUsername(user.Username) != null)
                throw new RegistrationException("Username already exists");
            user.Salt = CreateSalt(16);
            user.Password = EncodePassword(user.Password, user.Salt);
            _uow.GetRepository<IUserWriteRepository>().Add(user);
        }

        public string LogIn(string username, string password)
        {
            var user = _uow.GetRepository<IUserReadRepository>().GetByUsername(username);
            if (user == null) throw new LogInException("User with given username not found!");
            if (user.Password != EncodePassword(password, user.Salt)) throw new LogInException("Invalid password!");
            return _jwtGenerator.GenerateToken(user);
        }

        public static string CreateSalt(int size)
        {
            RNGCryptoServiceProvider rng = new RNGCryptoServiceProvider();
            byte[] buff = new byte[size];
            rng.GetBytes(buff);
            return Convert.ToBase64String(buff);
        }

        private string EncodePassword(string password, string salt)
        {
            using var sha = SHA256.Create();
            var computedHash = sha.ComputeHash(Encoding.Unicode.GetBytes(salt + password));
            return Convert.ToBase64String(computedHash);
        }
    }
}

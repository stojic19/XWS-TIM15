using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Threading.Tasks;

namespace AgentApplication.ClassLib.Service.Impl
{
    internal class Encoder
    {
        public static string CreateSalt(int size)
        {
            RNGCryptoServiceProvider rng = new RNGCryptoServiceProvider();
            byte[] buff = new byte[size];
            rng.GetBytes(buff);
            return Convert.ToBase64String(buff);
        }

        public static string EncodePassword(string password, string salt)
        {
            using var sha = SHA256.Create();
            var computedHash = sha.ComputeHash(Encoding.Unicode.GetBytes(salt + password));
            return Convert.ToBase64String(computedHash);
        }
    }
}

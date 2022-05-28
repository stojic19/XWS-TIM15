using AgentApplication.ClassLib.Model;
using System;
using AgentApplication.ClassLib.Model.Enumerations;

namespace AgentApplication.API.Dto
{
    public class PostUserDto
    {
        public string Username;
        public string Password;
        public string FirstName;
        public string MiddleName;
        public string LastName;
        public DateTime BirthDate;
        public string Email;
        public string PhoneNumber;
        public Gender Gender;
    }
}

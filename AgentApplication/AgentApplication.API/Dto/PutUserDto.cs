using AgentApplication.ClassLib.Model;
using AgentApplication.ClassLib.Model.Enumerations;
using System;

namespace AgentApplication.API.Dto
{
    public class PutUserDto
    {
        public string Id;
        public string Username;
        public string FirstName;
        public string MiddleName;
        public string LastName;
        public DateTime BirthDate;
        public string Email;
        public string PhoneNumber;
        public Gender Gender;
    }
}

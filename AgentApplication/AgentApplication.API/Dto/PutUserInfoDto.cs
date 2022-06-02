using AgentApplication.ClassLib.Model;
using AgentApplication.ClassLib.Model.Enumerations;
using System;

namespace AgentApplication.API.Dto
{
    public class PutUserInfoDto
    {
        public string FirstName { get; set; }
        public string MiddleName { get; set; }
        public string LastName { get; set; }
        public DateTime BirthDate { get; set; }
        public string Email { get; set; }
        public string PhoneNumber { get; set; }
        public Gender Gender { get; set; }
    }
}

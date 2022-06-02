using AgentApplication.ClassLib.Model;
using System;
using AgentApplication.ClassLib.Model.Enumerations;

namespace AgentApplication.API.Dto
{
    public class PostUserInfoDto
    {
        public string FirstName { get; set; }
        public string MiddleName { get; set; }
        public string LastName { get; set; }
        public DateTime BirthDate { get; set; }
        public string Email { get; set; }
        public string PhoneNumber { get; set; }
        public Gender Gender { get; set; }
    }

    public class PostUserDto
    {
        public string Username { get; set; }
        public string Password { get; set; }
        public PostUserInfoDto PersonalInfo { get; set; }
    }
}

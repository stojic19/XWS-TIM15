using AgentApplication.ClassLib.Model;
using System;
using AgentApplication.ClassLib.Model.Enumerations;

namespace AgentApplication.API.Dto
{
    public class PostUserInfoDto
    {
        public string FirstName;
        public string MiddleName;
        public string LastName;
        public DateTime BirthDate;
        public string Email;
        public string PhoneNumber;
        public Gender Gender;
    }

    public class PostUserDto
    {
        public string Username;
        public string Password;
        private readonly PostUserInfoDto _postUserInfoDto = new PostUserInfoDto();
    }
}

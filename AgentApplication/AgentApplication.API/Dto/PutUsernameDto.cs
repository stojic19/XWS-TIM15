using AgentApplication.ClassLib.Model;
using AgentApplication.ClassLib.Model.Enumerations;
using System;

namespace AgentApplication.API.Dto
{
    public class PutUsernameDto
    {
        public Guid Id { get; set; }
        public string Username { get; set; }
    }
}

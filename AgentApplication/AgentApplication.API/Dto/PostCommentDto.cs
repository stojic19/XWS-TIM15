using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PostCommentDto
    {
        public string Content;
        public Company Company;
        public User Owner;
    }
}

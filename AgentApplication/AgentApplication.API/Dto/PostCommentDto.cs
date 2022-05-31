using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PostCommentDto
    {
        public string Content;
        public Guid CompanyId;
        public Guid UserId;
    }
}

using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PostCommentDto
    {
        public string Content { get; set; }
        public Guid CompanyId { get; set; }
    }
}

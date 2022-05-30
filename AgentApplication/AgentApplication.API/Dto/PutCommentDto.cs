using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PutCommentDto
    {
        public Guid Id;
        public Guid CompanyId;
        public string Content;
    }
}

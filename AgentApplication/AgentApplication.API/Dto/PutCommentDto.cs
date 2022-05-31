using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace AgentApplication.API.Dto
{
    public class PutCommentDto
    {
        public Guid Id { get; set; }
        public Guid CompanyId { get; set; }
        public string Content { get; set; }
    }
}

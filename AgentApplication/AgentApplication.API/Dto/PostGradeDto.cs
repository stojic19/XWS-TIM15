using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PostGradeDto
    {
        public int Value { get; set; }
        public Guid CompanyId { get; set; }
        public Guid UserId { get; set; }
    }
}

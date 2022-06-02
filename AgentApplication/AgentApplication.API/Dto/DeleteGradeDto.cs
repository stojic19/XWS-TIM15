using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace AgentApplication.API.Dto
{
    public class DeleteGradeDto
    {
        public Guid Id { get; set; }
        public Guid CompanyId { get; set; }
    }
}

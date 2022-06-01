using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace AgentApplication.API.Dto
{
    public class DeleteJobOfferDto
    {
        public Guid Id { get; set; }
        public Guid CompanyId { get; set; }

    }
}

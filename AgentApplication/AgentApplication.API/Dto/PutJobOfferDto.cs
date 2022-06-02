using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace AgentApplication.API.Dto
{
    public class PutJobOfferDto
    {
        public Guid Id { get; set; }
        public Guid CompanyId { get; set; }
        public string Position { get; set; }
        public string Description { get; set; }
        public string Requirements { get; set; }
    }
}

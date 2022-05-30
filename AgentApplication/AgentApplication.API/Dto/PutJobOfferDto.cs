using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace AgentApplication.API.Dto
{
    public class PutJobOfferDto
    {
        public Guid Id;
        public Guid CompanyId;
        public string Position;
        public string Description;
        public string Requirements;
    }
}

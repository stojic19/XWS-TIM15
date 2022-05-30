using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PutJobOfferDto
    {
        public Guid Id;
        public Guid CompanyId;
        public string Position;
        public string Description;
        public string Requirements;
        public bool IsActive;
    }
}

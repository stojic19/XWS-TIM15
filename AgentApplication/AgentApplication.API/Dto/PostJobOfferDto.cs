using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PostJobOfferDto
    {
        public string Position;
        public string Description;
        public string Requirements;
        public Guid CompanyId;
    }
}

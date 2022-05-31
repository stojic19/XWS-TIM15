using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PostJobOfferDto
    {
        public string Position { get; set; }
        public string Description { get; set; }
        public string Requirements { get; set; }
        public Guid CompanyId { get; set; }
    }
}

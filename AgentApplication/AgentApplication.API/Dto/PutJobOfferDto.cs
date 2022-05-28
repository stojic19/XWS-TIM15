using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PutJobOfferDto
    {
        public string Id;
        public string Position;
        public string Description;
        public string Requirements;
        public bool IsActive;
    }
}

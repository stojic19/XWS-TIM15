using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PostCompanyInfoDto
    {
        public string Name;
        public string Address;
        public string Email;
        public string PhoneNumber;
        public string Description;
        public string Culture;
    }

    public class PostCompanyDto
    {
        public Guid OwnerId;
        public PostCompanyInfoDto CompanyInfo;
    }
}

using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PostCompanyInfoDto
    {
        public string Name { get; set; }
        public string Address { get; set; }
        public string Email { get; set; }
        public string PhoneNumber { get; set; }
        public string Description { get; set; }
        public string Culture { get; set; }
    }

    public class PostCompanyDto
    {
        public PostCompanyInfoDto CompanyInfo { get; set; }
    }
}

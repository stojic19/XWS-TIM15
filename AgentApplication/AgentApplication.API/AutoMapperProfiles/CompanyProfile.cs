using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Model;
using AutoMapper;

namespace AgentApplication.API.AutoMapperProfiles
{
    public class CompanyProfile : Profile
    {
        public CompanyProfile()
        {
            CreateMap<PostCompanyDto, Company>();
            CreateMap<PostCompanyInfoDto, CompanyInfo>();
        }
    }
}

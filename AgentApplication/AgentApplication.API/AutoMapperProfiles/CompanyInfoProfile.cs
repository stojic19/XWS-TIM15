using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Model;
using AutoMapper;

namespace AgentApplication.API.AutoMapperProfiles
{
    public class CompanyInfoProfile : Profile
    {
        public CompanyInfoProfile()
        {
            CreateMap<PutCompanyInfoDto, CompanyInfo>();
        }
    }
}

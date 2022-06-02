using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Model;
using AutoMapper;

namespace AgentApplication.API.AutoMapperProfiles
{
    public class UserInfoProfile : Profile
    {
        public UserInfoProfile()
        {
            CreateMap<PutUserInfoDto, UserPersonalInfo>();
        }
    }
}

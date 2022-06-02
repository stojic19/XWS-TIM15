using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Model;
using AutoMapper;

namespace AgentApplication.API.AutoMapperProfiles
{
    public class UserProfile : Profile
    {
        public UserProfile()
        {
            CreateMap<PostUserDto, User>();
            CreateMap<PostUserInfoDto, UserPersonalInfo>();
        }
    }
}

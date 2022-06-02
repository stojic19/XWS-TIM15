using System;
using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Model;
using AgentApplication.ClassLib.Model.Enumerations;
using AutoMapper;

namespace AgentApplication.API.AutoMapperProfiles
{
    public class UserProfile : Profile
    {
        public UserProfile()
        {
            CreateMap<PostUserDto, User>()
                .BeforeMap((s, d) =>
                {
                    d.Role = Role.Regular;
                    d.TimeOfRegistration = DateTime.Now;
                }); 
            CreateMap<PostUserInfoDto, UserPersonalInfo>();
        }
    }
}

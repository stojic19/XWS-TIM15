using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Model;
using AutoMapper;

namespace AgentApplication.API.AutoMapperProfiles
{
    public class GradeProfile : Profile
    {
        public GradeProfile()
        {
            CreateMap<PostGradeDto, Grade>();
        }
    }
}

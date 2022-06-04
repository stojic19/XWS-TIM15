using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Model;
using AutoMapper;

namespace AgentApplication.API.AutoMapperProfiles
{
    public class JobOfferProfile : Profile
    {
        public JobOfferProfile()
        {
            CreateMap<PostJobOfferDto, JobOffer>();
            CreateMap<JobOffer, NewJobOfferToDislinktDto>();
        }
    }
}

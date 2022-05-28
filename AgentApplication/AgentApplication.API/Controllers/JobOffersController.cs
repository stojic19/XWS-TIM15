using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Reflection;
using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Database.Infrastructure;
using AgentApplication.ClassLib.Database.Repository;
using AgentApplication.ClassLib.Database.Repository.Enums;
using AgentApplication.ClassLib.Model;
using AutoMapper;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace AgentApplication.API.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class JobOffersController : ControllerBase
    {
        private readonly IUnitOfWork _uow;
        private readonly IMapper _mapper;
        public JobOffersController(IUnitOfWork uow, IMapper mapper)
        {
            _uow = uow;
            _mapper = mapper;
        }

        [HttpGet]
        public IActionResult GetAll()
        {
            return Ok(_uow.GetRepository<IJobOfferReadRepository>().GetAll());
        }

        [HttpGet("{id:guid}")]
        public IActionResult GetById(Guid id)
        {
            return Ok(_uow.GetRepository<IJobOfferReadRepository>().GetById(id, FetchType.Eager));
        }

        [HttpPost]
        public IActionResult PostJobOffer(PostJobOfferDto dto)
        {
            JobOffer jobOffer = _mapper.Map<JobOffer>(dto);
            jobOffer.TimeOfCreation = DateTime.Now;
            jobOffer.IsActive = false;
            return Ok(_uow.GetRepository<IJobOfferWriteRepository>().Add(jobOffer));
        }

        [HttpPut]
        public IActionResult UpdateJobOffer(PutJobOfferDto dto)
        {
            return Ok(_uow.GetRepository<IJobOfferWriteRepository>().Update(_mapper.Map<JobOffer>(dto)));
        }
    }
}

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
    public class CompaniesController : ControllerBase
    {
        private readonly IUnitOfWork _uow;
        private readonly IMapper _mapper;
        public CompaniesController(IUnitOfWork uow, IMapper mapper)
        {
            _uow = uow;
            _mapper = mapper;
        }

        [HttpGet]
        public IActionResult GetAll()
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetAll());
        }

        [HttpGet("{id:guid}")]
        public IActionResult GetById(Guid id)
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetById(id, FetchType.Eager));
        }

        [HttpPost]
        public IActionResult PostCompany(PostCompanyDto dto)
        {
            Company company = _mapper.Map<Company>(dto);
            company.Registered = false;
            company.TimeOfCreation = DateTime.Now;
            company.Comments = new List<Comment>();
            company.Grades = new List<Grade>();
            company.JobOffers = new List<JobOffer>();
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Add(company));
        }

        [HttpPut]
        public IActionResult UpdateCompany(PutCompanyDto dto)
        {
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(_mapper.Map<Company>(dto)));
        }
    }
}

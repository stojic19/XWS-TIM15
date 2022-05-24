using System;
using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Database.Infrastructure;
using AgentApplication.ClassLib.Database.Repository;
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
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetById(id));
        }

        [HttpPost]
        public IActionResult PostCompany(PostCompanyDto dto)
        {
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Add(_mapper.Map<Company>(dto)));
        }
    }
}

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
    [Route("api/[controller]/[action]")]
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

        [HttpGet]
        public IActionResult GetAllDetails()
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetAll(FetchType.Eager));
        }

        [HttpGet("{id:guid}")]
        public IActionResult GetById(Guid id)
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetById(id, FetchType.Eager));
        }

        [HttpGet("{userId:guid}")]
        public IActionResult GetFromUser(Guid userId)
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetFromUser(userId, FetchType.Eager));
        }

        [HttpGet]
        public IActionResult GetRegistered()
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetRegistered(FetchType.Eager));
        }

        [HttpGet]
        public IActionResult GetNotRegistered()
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetNotRegistered(FetchType.Eager));
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
        public IActionResult UpdateCompany(PutCompanyInfoDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.Id);
            company.CompanyInfo = _mapper.Map<CompanyInfo>(dto);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPut]
        public IActionResult RegisterCompany(Guid id)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(id);
            company.Registered = true;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPut]
        public IActionResult AddGrade(PostGradeDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId);
            if (company.Grades.FirstOrDefault(g => g.UserId == dto.UserId) != null)
            {
                Grade grade = company.Grades.FirstOrDefault(g => g.UserId == dto.UserId);
                grade.Value = dto.Value;
                grade.TimeOfCreation = DateTime.Now;
                return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
            }
            Grade newGrade = _mapper.Map<Grade>(dto);
            newGrade.TimeOfCreation = DateTime.Now;
            company.Grades.Add(newGrade);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPost]
        public IActionResult AddComment(PostCommentDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId);
            Comment comment = _mapper.Map<Comment>(dto);
            comment.TimeOfCreation = DateTime.Now;
            company.Comments.Add(comment);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPost]
        public IActionResult AddJobOffer(PostJobOfferDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId);
            JobOffer jobOffer = _mapper.Map<JobOffer>(dto);
            jobOffer.TimeOfCreation = DateTime.Now;
            company.JobOffers.Add(jobOffer);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPut]
        public IActionResult ActivateJobOffer(ActivateJobOfferDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId);
            JobOffer jobOffer = company.JobOffers.FirstOrDefault(g => g.Id == dto.Id);
            jobOffer.IsActive = true;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }
    }
}

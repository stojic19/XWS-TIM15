using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Reflection;
using AgentApplication.API.Attributes;
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

        [Authorize(new []{"Regular"})]
        [HttpGet]
        public IActionResult GetAll()
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetAll());
        }

        [HttpGet("Details")]
        public IActionResult GetAllDetails()
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetAll(FetchType.Eager));
        }

        [HttpGet("{id:guid}")]
        public IActionResult GetById(Guid id)
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetById(id, FetchType.Eager));
        }

        [HttpGet("User/{userId:guid}")]
        public IActionResult GetFromUser(Guid userId)
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetFromUser(userId, FetchType.Eager));
        }

        [HttpGet("Registered")]
        public IActionResult GetRegistered()
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetRegistered(FetchType.Eager));
        }

        [HttpGet("NotRegistered")]
        public IActionResult GetNotRegistered()
        {
            return Ok(_uow.GetRepository<ICompanyReadRepository>().GetNotRegistered(FetchType.Eager));
        }

        [HttpPost]
        public IActionResult PostCompany(PostCompanyDto dto)
        {
            Company company = _mapper.Map<Company>(dto);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Add(company));
        }

        [HttpPut]
        public IActionResult UpdateCompany(PutCompanyInfoDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.Id);
            if (company == null) return NotFound("Company not found");
            company.CompanyInfo = _mapper.Map<CompanyInfo>(dto);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPut("{id:guid}/Register")]
        public IActionResult RegisterCompany(Guid id)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(id);
            if (company == null) return NotFound("Company not found");
            company.Registered = true;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPut("Grade")]
        public IActionResult AddGrade(PostGradeDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            if (company.Grades.FirstOrDefault(g => g.UserId == dto.UserId) != null)
            {
                Grade grade = company.Grades.FirstOrDefault(g => g.UserId == dto.UserId);
                if (grade == null) return NotFound("Grade not found");
                grade.Value = dto.Value;
                grade.TimeOfCreation = DateTime.Now;
                return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
            }
            Grade newGrade = _mapper.Map<Grade>(dto);
            company.Grades.Add(newGrade);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPost("Comment")]
        public IActionResult AddComment(PostCommentDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            Comment comment = _mapper.Map<Comment>(dto);
            company.Comments.Add(comment);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPut("Comment")]
        public IActionResult UpdateComment(PutCommentDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            Comment comment = company.Comments.Find(c => c.Id == dto.Id);
            if (comment == null) return NotFound("Comment not found");
            comment.Content = dto.Content;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPost("JobOffer")]
        public IActionResult AddJobOffer(PostJobOfferDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            JobOffer jobOffer = _mapper.Map<JobOffer>(dto);
            company.JobOffers.Add(jobOffer);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPut("JobOffer/Activate")]
        public IActionResult ActivateJobOffer(ActivateJobOfferDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            JobOffer jobOffer = company.JobOffers.FirstOrDefault(g => g.Id == dto.Id);
            if (jobOffer == null) return NotFound("Job offer not found");
            jobOffer.IsActive = true;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpPut("JobOffer")]
        public IActionResult UpdateJobOffer(PutJobOfferDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            JobOffer jobOffer = company.JobOffers.Find(c => c.Id == dto.Id);
            if (jobOffer == null) return NotFound("Job offer not found");
            jobOffer.Position = dto.Position;
            jobOffer.Description = dto.Description;
            jobOffer.Requirements = dto.Requirements;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }


        [HttpDelete("Grade")]
        public IActionResult DeleteGrade(DeleteGradeDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            Grade grade = company.Grades.FirstOrDefault(g => g.Id == dto.Id);
            if (grade == null) return NotFound("Grade not found");
            company.Grades.Remove(grade);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpDelete("Comment")]
        public IActionResult DeleteComment(DeleteCommentDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            Comment comment = company.Comments.FirstOrDefault(c => c.Id == dto.Id);
            if (comment == null) return NotFound("Comment not found");
            company.Comments.Remove(comment);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [HttpDelete("JobOffer")]
        public IActionResult DeleteJobOffer(DeleteJobOfferDto dto)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            JobOffer jobOffer = company.JobOffers.FirstOrDefault(j => j.Id == dto.Id);
            if (jobOffer == null) return NotFound("Job offer not found");
            company.JobOffers.Remove(jobOffer);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }
    }
}

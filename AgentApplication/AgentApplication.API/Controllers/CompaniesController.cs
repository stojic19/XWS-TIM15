using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Reflection;
using System.Threading.Tasks;
using AgentApplication.API.Attributes;
using AgentApplication.API.Controllers.Base;
using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Database.Infrastructure;
using AgentApplication.ClassLib.Database.Repository;
using AgentApplication.ClassLib.Database.Repository.Enums;
using AgentApplication.ClassLib.Model;
using AutoMapper;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;

namespace AgentApplication.API.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class CompaniesController : BaseApiController
    {
        public CompaniesController(IUnitOfWork uow, IMapper mapper, IConfiguration config) : base(uow, mapper, config) { }

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

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPost]
        public IActionResult PostCompany(PostCompanyDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            Company company = _mapper.Map<Company>(dto);
            company.OwnerId = id;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Add(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPut]
        public IActionResult UpdateCompany(PutCompanyInfoDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            var role = (HttpContext.Items["role"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.Id);
            if (company == null) return NotFound("Company not found");
            if (!role.Equals("Admin") && company.OwnerId != id)
            {
                return Unauthorized("You are not allowed to update this company!");
            }
            company.CompanyInfo = _mapper.Map<CompanyInfo>(dto);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Admin" })]
        [HttpPut("{id:guid}/Register")]
        public IActionResult RegisterCompany(Guid id)
        {
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(id);
            if (company == null) return NotFound("Company not found");
            company.Registered = true;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPut("Grade")]
        public IActionResult AddGrade(PostGradeDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) 
                return NotFound("Company not found");
            if (company.OwnerId == id)
                return BadRequest("Grading your own company is not allowed!");
            
            var grade = company.Grades.FirstOrDefault(g => g.UserId == id);
            if (grade != null)
            {
                grade.Value = dto.Value;
                grade.TimeOfCreation = DateTime.Now;
                return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
            }

            Grade newGrade = _mapper.Map<Grade>(dto);
            newGrade.UserId = id;
            company.Grades.Add(newGrade);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPost("Comment")]
        public IActionResult AddComment(PostCommentDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) 
                return NotFound("Company not found");
            if (company.OwnerId == id)
                return BadRequest("Commenting on your own company is not allowed!");
            Comment comment = _mapper.Map<Comment>(dto);
            comment.UserId = id;
            company.Comments.Add(comment);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPut("Comment")]
        public IActionResult UpdateComment(PutCommentDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) 
                return NotFound("Company not found");
            Comment comment = company.Comments.Find(c => c.Id == dto.Id);
            if (comment == null) 
                return NotFound("Comment not found");
            if (comment.UserId != id)
                return BadRequest("Changing other user's comments is not allowed!");
            comment.Content = dto.Content;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPost("JobOffer")]
        public IActionResult AddJobOffer(PostJobOfferDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            if (company.OwnerId != id) return BadRequest("Cannot add job offer since you do not own this company!");
            JobOffer jobOffer = _mapper.Map<JobOffer>(dto);
            company.JobOffers.Add(jobOffer);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPut("JobOffer/Activate")]
        public async Task<IActionResult> ActivateJobOffer(ActivateJobOfferDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            if (company.OwnerId != id) return BadRequest("Cannot activate job offer since you do not own this company!");
            JobOffer jobOffer = company.JobOffers.FirstOrDefault(g => g.Id == dto.Id);
            if (jobOffer == null) return NotFound("Job offer not found");
            if (jobOffer.IsActive) return BadRequest("Job offer already active");
            jobOffer.IsActive = true;
            var request = new HttpRequestMessage
            {
                Method = HttpMethod.Post,
                RequestUri = new Uri(_dislinktApiGatewayBaseUrl + "job_offers"),
                Content = GetContent(_mapper.Map<NewJobOfferToDislinktDto>(jobOffer))
            };
            request.Headers.Add("apiKey", _config["DislinktApiKey"]);
            var result = await _httpClient.SendAsync(request);
            if (result.StatusCode != HttpStatusCode.OK) return BadRequest(result.ToString());
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPut("JobOffer")]
        public IActionResult UpdateJobOffer(PutJobOfferDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            if (company.OwnerId != id) return BadRequest("Cannot update job offer since you do not own this company!");
            JobOffer jobOffer = company.JobOffers.Find(c => c.Id == dto.Id);
            if (jobOffer == null) return NotFound("Job offer not found");
            jobOffer.Position = dto.Position;
            jobOffer.Description = dto.Description;
            jobOffer.Requirements = dto.Requirements;
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpDelete("Grade")]
        public IActionResult DeleteGrade(DeleteGradeDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            var role = (HttpContext.Items["role"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            Grade grade = company.Grades.FirstOrDefault(g => g.Id == dto.Id);
            if (!role.Equals("Admin") && grade.UserId != id)
            {
                return Unauthorized("You are not allowed to delete this grade!");
            }
            if (grade == null) return NotFound("Grade not found");
            company.Grades.Remove(grade);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpDelete("Comment")]
        public IActionResult DeleteComment(DeleteCommentDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            var role = (HttpContext.Items["role"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            Comment comment = company.Comments.FirstOrDefault(c => c.Id == dto.Id);
            if (comment == null) return NotFound("Comment not found");
            if (!role.Equals("Admin") && comment.UserId != id)
            {
                return Unauthorized("You are not allowed to delete this comment!");
            }
            company.Comments.Remove(comment);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpDelete("JobOffer")]
        public IActionResult DeleteJobOffer(DeleteJobOfferDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            var role = (HttpContext.Items["role"]?.ToString() ?? string.Empty);
            Company company = _uow.GetRepository<ICompanyReadRepository>().GetById(dto.CompanyId, FetchType.Eager);
            if (company == null) return NotFound("Company not found");
            if (!role.Equals("Admin") && company.OwnerId != id)
            {
                return Unauthorized("You are not the owner of this company! You cannot delete its job offers");
            }
            JobOffer jobOffer = company.JobOffers.FirstOrDefault(j => j.Id == dto.Id);
            if (jobOffer == null) return NotFound("Job offer not found");
            company.JobOffers.Remove(jobOffer);
            return Ok(_uow.GetRepository<ICompanyWriteRepository>().Update(company));
        }
    }
}

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
using AgentApplication.ClassLib.Exceptions;
using AgentApplication.ClassLib.Model;
using AgentApplication.ClassLib.Model.Enumerations;
using AgentApplication.ClassLib.Service;
using AutoMapper;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace AgentApplication.API.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class UsersController : ControllerBase
    {
        private readonly IUnitOfWork _uow;
        private readonly IMapper _mapper;
        private readonly IAuthenticationService _authenticationService;
        private readonly IJwtGenerator _jwtGenerator;
        public UsersController(IUnitOfWork uow, IMapper mapper, IAuthenticationService authenticationService, IJwtGenerator jwtGenerator)
        {
            _uow = uow;
            _mapper = mapper;
            _authenticationService = authenticationService;
            _jwtGenerator = jwtGenerator;
        }

        [HttpGet]
        public IActionResult GetAll()
        {
            return Ok(_uow.GetRepository<IUserReadRepository>().GetAll());
        }

        [HttpGet("{id:guid}")]
        public IActionResult GetById(Guid id)
        {
            return Ok(_uow.GetRepository<IUserReadRepository>().GetById(id, FetchType.Eager));
        }

        [HttpPost]
        public IActionResult PostUser(PostUserDto dto)
        {
            User user = _mapper.Map<User>(dto);
            try
            {
                _authenticationService.Register(user);
                return Ok();
            }
            catch (Exception ex)
            {
                switch (ex)
                {
                    case RegistrationException: return BadRequest(ex.Message);
                    default: return Problem("Oops, something went wrong! Try again later.");
                }
            }
            
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPut("Username")]
        public IActionResult UpdateUsername(PutUsernameDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            User user = _uow.GetRepository<IUserReadRepository>().GetById(id);
            if (user == null) return NotFound("User not found");
            if (_uow.GetRepository<IUserReadRepository>().GetByUsername(dto.Username) != null)
                return BadRequest("Username already exists!");
            user.Username = dto.Username;
            return Ok(_uow.GetRepository<IUserWriteRepository>().Update(user));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpPut("Info")]
        public IActionResult UpdateUserInfo(PutUserInfoDto dto)
        {
            Guid id = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            User user = _uow.GetRepository<IUserReadRepository>().GetById(id);
            if (user == null) return NotFound("User not found");
            user.PersonalInfo = _mapper.Map<UserPersonalInfo>(dto);
            return Ok(_uow.GetRepository<IUserWriteRepository>().Update(user));
        }

        [Authorize(new[] { "Regular", "Admin" })]
        [HttpDelete]
        public IActionResult DeleteUser(Guid id)
        {
            Guid fromJwtId = Guid.Parse(HttpContext.Items["id"]?.ToString() ?? string.Empty);
            var role = (HttpContext.Items["role"]?.ToString() ?? string.Empty);
            if (id != fromJwtId && !role.Equals("Admin"))
                return Unauthorized("Error while deleting user!");
            User user = _uow.GetRepository<IUserReadRepository>().GetById(id);
            if (user == null) return NotFound("User not found");
            _uow.GetRepository<IUserWriteRepository>().Delete(user);
            return Ok();
        }
    }
}

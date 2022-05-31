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
using AgentApplication.ClassLib.Model.Enumerations;
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
        public UsersController(IUnitOfWork uow, IMapper mapper)
        {
            _uow = uow;
            _mapper = mapper;
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
            user.TimeOfRegistration = DateTime.Now;
            user.Role = Role.Regular;
            return Ok(_uow.GetRepository<IUserWriteRepository>().Add(user));
        }

        [HttpPut("Username")]
        public IActionResult UpdateUsername(PutUsernameDto dto)
        {
            User user = _uow.GetRepository<IUserReadRepository>().GetById(dto.Id);
            if (user == null) return NotFound("User not found");
            user.Username = dto.Username;
            return Ok(_uow.GetRepository<IUserWriteRepository>().Update(user));
        }

        [HttpPut("Info")]
        public IActionResult UpdateUserInfo(PutUserInfoDto dto)
        {
            User user = _uow.GetRepository<IUserReadRepository>().GetById(dto.Id);
            if (user == null) return NotFound("User not found");
            user.PersonalInfo = _mapper.Map<UserPersonalInfo>(dto);
            return Ok(_uow.GetRepository<IUserWriteRepository>().Update(user));
        }

        [HttpDelete]
        public IActionResult DeleteUser(Guid id)
        {
            User user = _uow.GetRepository<IUserReadRepository>().GetById(id);
            if (user == null) return NotFound("User not found");
            _uow.GetRepository<IUserWriteRepository>().Delete(user);
            return Ok();
        }
    }
}

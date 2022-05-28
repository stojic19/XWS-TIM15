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
    public class GradesController : ControllerBase
    {
        private readonly IUnitOfWork _uow;
        private readonly IMapper _mapper;
        public GradesController(IUnitOfWork uow, IMapper mapper)
        {
            _uow = uow;
            _mapper = mapper;
        }

        [HttpGet]
        public IActionResult GetAll()
        {
            return Ok(_uow.GetRepository<IGradeReadRepository>().GetAll());
        }

        [HttpGet("{id:guid}")]
        public IActionResult GetById(Guid id)
        {
            return Ok(_uow.GetRepository<IGradeReadRepository>().GetById(id, FetchType.Eager));
        }

        [HttpPost]
        public IActionResult PostGrade(PostGradeDto dto)
        {
            Grade grade = _mapper.Map<Grade>(dto);
            grade.TimeOfCreation = DateTime.Now;
            return Ok(_uow.GetRepository<IGradeWriteRepository>().Add(grade));
        }

        [HttpPut]
        public IActionResult UpdateGrade(PutGradeDto dto)
        {
            return Ok(_uow.GetRepository<IGradeWriteRepository>().Update(_mapper.Map<Grade>(dto)));
        }
    }
}

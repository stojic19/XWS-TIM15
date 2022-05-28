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
    public class CommentsController : ControllerBase
    {
        private readonly IUnitOfWork _uow;
        private readonly IMapper _mapper;
        public CommentsController(IUnitOfWork uow, IMapper mapper)
        {
            _uow = uow;
            _mapper = mapper;
        }

        [HttpGet]
        public IActionResult GetAll()
        {
            return Ok(_uow.GetRepository<ICommentReadRepository>().GetAll());
        }

        [HttpGet("{id:guid}")]
        public IActionResult GetById(Guid id)
        {
            return Ok(_uow.GetRepository<ICommentReadRepository>().GetById(id, FetchType.Eager));
        }

        [HttpPost]
        public IActionResult PostComment(PostCompanyDto dto)
        {
            Comment comment = _mapper.Map<Comment>(dto);
            comment.TimeOfCreation = DateTime.Now;
            return Ok(_uow.GetRepository<ICommentWriteRepository>().Add(comment));
        }

        [HttpPut]
        public IActionResult UpdateComment(PutCompanyDto dto)
        {
            return Ok(_uow.GetRepository<ICommentWriteRepository>().Update(_mapper.Map<Comment>(dto)));
        }
    }
}

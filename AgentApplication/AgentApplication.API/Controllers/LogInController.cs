using System;
using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Exceptions;
using AgentApplication.ClassLib.Service;
using Microsoft.AspNetCore.Cors;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace AgentApplication.API.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class LogInController : ControllerBase
    {
        private readonly IAuthenticationService _authenticationService;

        public LogInController(IAuthenticationService authenticationService)
        {
            _authenticationService = authenticationService;
        }

        [EnableCors]
        [HttpPost]
        public IActionResult LogIn(LogInDto dto)
        {
            try
            {
                var tokenAndRole = _authenticationService.LogIn(dto.Username, dto.Password);
                return Ok(new LogInResponseDto() { Token = tokenAndRole[0], Role = tokenAndRole[1] });
            }
            catch (Exception ex)
            {
                switch (ex)
                {
                    case LogInException: return NotFound(ex.Message);
                    default: return Problem("Oops, something went wrong. Try again later");
                }
            }
        }
    }
}

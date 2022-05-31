using AgentApplication.API.Dto;
using AgentApplication.ClassLib.Service;
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

        [HttpPost]
        public IActionResult LogIn(LogInDto dto)
        {
            var token = _authenticationService.LogIn(dto.Username, dto.Password);
            return Ok(token);
        }
    }
}

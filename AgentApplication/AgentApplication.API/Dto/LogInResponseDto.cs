using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace AgentApplication.API.Dto
{
    public class LogInResponseDto
    {
        public string Token { get; set; }
        public string Role { get; set; }
        public string Id { get; set; }
    
    }
}

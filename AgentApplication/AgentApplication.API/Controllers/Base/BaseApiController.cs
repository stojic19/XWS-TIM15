using System.Net.Http;
using System.Text;
using AgentApplication.ClassLib.Database.Infrastructure;
using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

namespace AgentApplication.API.Controllers.Base
{
    public class BaseApiController : ControllerBase
    {
        protected string _dislinktApiGatewayBaseUrl => "http://localhost:8000/";
        protected static HttpClient _httpClient = new ();
        protected readonly IUnitOfWork _uow;
        protected readonly IMapper _mapper;

        public BaseApiController(IUnitOfWork uow, IMapper mapper)
        {
            _uow = uow;
            _mapper = mapper;
        }

        protected StringContent GetContent(object content)
        {
            return new StringContent(JsonConvert.SerializeObject(content, settings: new JsonSerializerSettings()
            {
                ReferenceLoopHandling = Newtonsoft.Json.ReferenceLoopHandling.Ignore
            }), Encoding.UTF8, "application/json");
        }
    }
}

using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PutGradeDto
    {
        public Guid Id;
        public Guid CompanyId;
        public int Value;
    }
}

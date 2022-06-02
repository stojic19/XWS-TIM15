﻿using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc.Filters;

namespace AgentApplication.API.Attributes
{
    public class AuthorizeAttribute : Attribute, Microsoft.AspNetCore.Mvc.Filters.IAuthorizationFilter
    {
        private readonly List<string> _roles;

        public AuthorizeAttribute(string[] roles)
        {
            _roles = roles.ToList();
        }
        public void OnAuthorization(AuthorizationFilterContext context)
        {
            if (_roles.Count == 0) return;
            var role = context.HttpContext.Items["role"];
            bool found = false;
            foreach (var atrRole in _roles)
                if (atrRole.Equals(role))
                {
                    found = true;
                    break;
                }

            if (!found || context.HttpContext.Items["given_name"] == null || context.HttpContext.Items["id"] == null)
            {
                context.Result = new Microsoft.AspNetCore.Mvc.JsonResult(new { message = "Unauthorized" }) { StatusCode = StatusCodes.Status401Unauthorized };
            }
        }
    }
}

﻿using AgentApplication.ClassLib.Model;
using System;

namespace AgentApplication.API.Dto
{
    public class PostCompanyDto
    {
        public string Name;
        public string Address;
        public string Email;
        public string PhoneNumber;
        public string Description;
        public string Culture;
        public User Owner;
    }
}

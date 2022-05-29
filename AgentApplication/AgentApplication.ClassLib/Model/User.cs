﻿using AgentApplication.ClassLib.Model.Enumerations;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AgentApplication.ClassLib.Model
{
    

    public class User : PersistentEntity
    {
        public string Username { get; set; }
        public string Password { get; set; }
        public UserPersonalInfo PersonalInfo { get; set; }
        public DateTime TimeOfRegistration { get; set; }
        public Role Role { get; set; }
    }
}

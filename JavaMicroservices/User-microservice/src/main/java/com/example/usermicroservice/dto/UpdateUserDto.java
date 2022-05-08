package com.example.usermicroservice.dto;

import com.example.usermicroservice.model.Education;
import com.example.usermicroservice.model.Gender;
import com.example.usermicroservice.model.WorkExperience;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

public class UpdateUserDto extends AddUserDto{

    private String newUsername;

    public UpdateUserDto(String id, String username, String password, String name, String email, String telephoneNo,
                         String gender, Date dateOfBirth, String biography, boolean isPrivate,
                         List<Education> educationList, List<WorkExperience> workExperienceList, List<String> interests,
                         List<String> skills, String newUsername) {
        super(id, username, password, name, email, telephoneNo, gender, dateOfBirth, biography,
                isPrivate, educationList, workExperienceList, interests, skills);
        this.newUsername = newUsername;
    }

    public String getNewUsername() {
        return newUsername;
    }

    public void setNewUsername(String newUsername) {
        this.newUsername = newUsername;
    }
}

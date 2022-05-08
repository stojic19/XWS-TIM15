package com.example.usermicroservice.dto;

import com.example.usermicroservice.model.WorkExperience;
import lombok.Getter;
import lombok.Setter;

import java.util.Date;
import java.util.List;

public class UpdateWorkExperienceDto {
    private String userId;
    private List<WorkExperience> workExperienceList;

    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }

    public List<WorkExperience> getWorkExperienceList() {
        return workExperienceList;
    }

    public void setWorkExperienceList(List<WorkExperience> workExperienceList) {
        this.workExperienceList = workExperienceList;
    }
}

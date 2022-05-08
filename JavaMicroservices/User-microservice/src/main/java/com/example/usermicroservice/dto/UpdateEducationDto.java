package com.example.usermicroservice.dto;

import com.example.usermicroservice.model.Education;

import java.util.List;

public class UpdateEducationDto {
    private String userId;
    private List<Education> educationList;
    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }

    public List<Education> getEducationList() {
        return educationList;
    }

    public void setEducationList(List<Education> educationList) {
        this.educationList = educationList;
    }
}

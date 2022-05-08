package com.example.usermicroservice.dto;

import java.util.List;

public class UpdateSkillsDto {
    private String userId;
    private List<String> skills;

    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }

    public List<String> getSkills() {
        return skills;
    }

    public void setSkills(List<String> skills) {
        this.skills = skills;
    }
}

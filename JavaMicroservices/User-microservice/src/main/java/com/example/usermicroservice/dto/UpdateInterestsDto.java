package com.example.usermicroservice.dto;

import java.util.List;

public class UpdateInterestsDto {
    private String userId;
    private List<String> interests;

    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }

    public List<String> getInterests() {
        return interests;
    }

    public void setInterests(List<String> interests) {
        this.interests = interests;
    }
}

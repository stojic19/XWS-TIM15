package com.example.usermicroservice.dto;

import com.example.usermicroservice.model.Gender;

import java.util.Date;

public class UpdateUserDto extends UserDto{

    private String newUsername;

    public UpdateUserDto(String id, String username, String password, String name, String email, String telephoneNo, String gender, Date dateOfBirth, String biography, String newUsername) {
        super(id, username, password, name, email, telephoneNo, gender, dateOfBirth, biography);
        this.newUsername = newUsername;
    }

    public String getNewUsername() {
        return newUsername;
    }

    public void setNewUsername(String newUsername) {
        this.newUsername = newUsername;
    }
}

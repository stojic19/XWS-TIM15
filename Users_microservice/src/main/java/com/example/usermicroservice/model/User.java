package com.example.usermicroservice.model;

import lombok.Getter;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.Date;

@Getter
@Setter
@Document(collection = "User")
public class User {
    @Id
    private int id;
    private String username;
    private String password;
    private String name;
    private String email;
    private String telephoneNo;
    private Gender gender;
    private Date dateOfBirth;
    private String biography;
}

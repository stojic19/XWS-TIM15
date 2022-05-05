package com.example.usermicroservice.dto;

import com.example.usermicroservice.model.Education;
import com.example.usermicroservice.model.Gender;
import com.example.usermicroservice.model.WorkExperience;

import javax.validation.constraints.Email;
import javax.validation.constraints.NotEmpty;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

public class UserDto {

    private String id;
    @NotEmpty(message = "Username cannot be null or empty!")
    private String username;
    @NotEmpty(message = "Password cannot be null or empty!")
    private String password;
    @NotEmpty(message = "Name cannot be null or empty!")
    private String name;
    @Email (message = "Email should be valid!")
    private String email;
    @NotEmpty(message = "Telephone number cannot be null or empty!")
    private String telephoneNo;
    @NotEmpty(message = "Gender cannot be null or empty!")
    private Gender gender;
    @NotEmpty(message = "Date of birth cannot be null or empty!")
    private Date dateOfBirth;
    @NotEmpty(message = "Biography cannot be null or empty!")
    private String biography;
    private boolean isPrivate;
    private List<Education> educationList;
    private List<WorkExperience> workExperienceList;

    public UserDto(String id, String username, String password, String name,
                   String email, String telephoneNo, Gender gender, Date dateOfBirth,
                   String biography) {
        this.id = id;
        this.username = username;
        this.password = password;
        this.name = name;
        this.email = email;
        this.telephoneNo = telephoneNo;
        this.gender = gender;
        this.dateOfBirth = dateOfBirth;
        this.biography = biography;
        this.isPrivate = false;
        this.educationList = new ArrayList<Education>();
        this.workExperienceList = new ArrayList<WorkExperience>();

    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getTelephoneNo() {
        return telephoneNo;
    }

    public void setTelephoneNo(String telephoneNo) {
        this.telephoneNo = telephoneNo;
    }

    public Gender getGender() {
        return gender;
    }

    public void setGender(Gender gender) {
        this.gender = gender;
    }

    public Date getDateOfBirth() {
        return dateOfBirth;
    }

    public void setDateOfBirth(Date dateOfBirth) {
        this.dateOfBirth = dateOfBirth;
    }

    public String getBiography() {
        return biography;
    }

    public void setBiography(String biography) {
        this.biography = biography;

    }

    public boolean isPrivate() {
        return isPrivate;
    }

    public void setPrivate(boolean aPrivate) {
        isPrivate = aPrivate;
    }

    public List<Education> getEducationList() {
        return educationList;
    }

    public void setEducationList(List<Education> educationList) {
        this.educationList = educationList;
    }

    public List<WorkExperience> getWorkExperienceList() {
        return workExperienceList;
    }

    public void setWorkExperienceList(List<WorkExperience> workExperienceList) {
        this.workExperienceList = workExperienceList;
    }
}

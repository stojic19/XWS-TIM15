package com.example.usermicroservice.service;

import com.example.usermicroservice.dto.*;
import com.example.usermicroservice.model.User;

import java.util.List;
import java.util.Optional;

public interface UserService {
    boolean usernameExists(String username);
    User addUser(User user);
    List<User> getUsers();
    Optional<User> findById(String id);
    User findByUsername(String username);
    void deleteById(String id);
    User updateUser(User userWithNewData);
    User updateWorkExperience(UpdateWorkExperienceDto updateWorkExperienceDto);
    User updateEducation(UpdateEducationDto updateEducationDto);
    User updateInterests(UpdateInterestsDto updateInterestsDto);
    User updateSkills(UpdateSkillsDto updateSkillsDto);
    List<User> searchPublicUsers(String searchTerm);
}

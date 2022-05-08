package com.example.usermicroservice.service.Impl;

import com.example.usermicroservice.dto.UpdateEducationDto;
import com.example.usermicroservice.dto.UpdateInterestsDto;
import com.example.usermicroservice.dto.UpdateSkillsDto;
import com.example.usermicroservice.dto.UpdateWorkExperienceDto;
import com.example.usermicroservice.model.User;
import com.example.usermicroservice.model.WorkExperience;
import com.example.usermicroservice.repository.UserRepository;
import com.example.usermicroservice.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
public class UserServiceImpl implements UserService {

    @Autowired
    private UserRepository userRepository;

    @Override
    public boolean usernameExists(String username) {
        return userRepository.findUserByUsername(username) != null;
    }

    @Override
    public User addUser(User user) {
        user.setId(UUID.randomUUID().toString());
        user.setWorkExperience(new ArrayList<>());
        user.setEducation(new ArrayList<>());
        user.setInterests(new ArrayList<>());
        user.setSkills(new ArrayList<>());
        userRepository.save(user);
        return user;
    }

    @Override
    public List<User> getUsers() {
        return userRepository.findAll();
    }

    @Override
    public Optional<User> findById(String id) {
        return userRepository.findById(id);
    }

    @Override
    public User findByUsername(String username) {
        return userRepository.findUserByUsername(username);
    }

    @Override
    public void deleteById(String id) {
        userRepository.deleteById(id);
    }

    @Override
    public User updateUser(User userWithNewData) {
        Optional<User> user = findById(userWithNewData.getId());
        userWithNewData.setEducation(user.get().getEducation());
        userWithNewData.setWorkExperience(user.get().getWorkExperience());
        userWithNewData.setSkills(user.get().getSkills());
        userWithNewData.setInterests(user.get().getInterests());
        return userRepository.save(userWithNewData);
    }

    @Override
    public User updateWorkExperience(UpdateWorkExperienceDto updateWorkExperienceDto) {
        Optional<User> user = findById(updateWorkExperienceDto.getUserId());
        User userForUpdate = user.get();
        userForUpdate.setWorkExperience(updateWorkExperienceDto.getWorkExperienceList());
        return userRepository.save(userForUpdate);
    }

    @Override
    public User updateEducation(UpdateEducationDto updateEducationDto) {
        Optional<User> user = findById(updateEducationDto.getUserId());
        User userForUpdate = user.get();
        userForUpdate.setEducation(updateEducationDto.getEducationList());
        return userRepository.save(userForUpdate);
    }

    @Override
    public User updateInterests(UpdateInterestsDto updateInterestsDto) {
        Optional<User> user = findById(updateInterestsDto.getUserId());
        User userForUpdate = user.get();
        userForUpdate.setInterests(updateInterestsDto.getInterests());
        return userRepository.save(userForUpdate);
    }

    @Override
    public User updateSkills(UpdateSkillsDto updateSkillsDto) {
        Optional<User> user = findById(updateSkillsDto.getUserId());
        User userForUpdate = user.get();
        userForUpdate.setSkills(updateSkillsDto.getSkills());
        return userRepository.save(userForUpdate);
    }

    @Override
    public List<User> searchPublicUsers(String searchTerm) {
        List<User> publicUsers = new ArrayList<>();
        for(User user:userRepository.findUsersByIsPrivateFalse()){
            if(user.getUsername().toLowerCase().contains(searchTerm.toLowerCase())
                    ||user.getName().toLowerCase().contains(searchTerm.toLowerCase()))
                publicUsers.add(user);
        }
        return publicUsers;
    }
}

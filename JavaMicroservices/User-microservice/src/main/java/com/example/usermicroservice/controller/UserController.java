package com.example.usermicroservice.controller;

import com.example.usermicroservice.dto.*;
import com.example.usermicroservice.mapper.UserMapper;
import com.example.usermicroservice.model.User;
import com.example.usermicroservice.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/users")
public class UserController {
    @Autowired
    private UserService userService;

    final private UserMapper userMapper = new UserMapper();

    @PostMapping()
    public String saveUser(@RequestBody AddUserDto userDto){
        if(isNullOrEmpty(userDto.getUsername(), userDto.getPassword(), userDto.getName(), userDto.getEmail(), userDto.getTelephoneNo(), userDto.getGender().toString(), userDto.getDateOfBirth().toString(), userDto.getBiography()))
            return "None of fields cannot be empty!";
        if(userService.usernameExists(userDto.getUsername()))
            return "Username already exists!";
        return "Added user with id " + userService.addUser(userMapper.AddUserDtoToUser(userDto)).getId();
    }

    @PutMapping()
    public String updateUser(@RequestBody UpdateUserDto userDto){
        if(isNullOrEmpty(userDto.getId(), userDto.getUsername(), userDto.getPassword(), userDto.getName(), userDto.getEmail(), userDto.getTelephoneNo(), userDto.getGender().toString(), userDto.getDateOfBirth().toString(), userDto.getBiography()))
            return "None of fields cannot be empty!";
        if(!isNullOrEmpty(userDto.getNewUsername())&&!userDto.getUsername().equals(userDto.getNewUsername()))
            if(userService.usernameExists(userDto.getNewUsername()))
                return "Username already exists!";
            else
                userDto.setUsername(userDto.getNewUsername());
        return "Updated user with id " + userService.updateUser(userMapper.UpdateUserDtoToUser(userDto)).getId();
    }

    private static boolean isNullOrEmpty(String... strArr){
        for (String st : strArr) {
            if  (st==null || st.equals(""))
                return true;

        }
        return false;
    }

    @GetMapping
    public List<UserDto> getUsers(){
        List<UserDto> userDtoList = new ArrayList<>();
        for(User user : userService.getUsers())
            userDtoList.add(userMapper.UserToUserDto(user));
        return userDtoList;
    }

    @GetMapping("/{id}")
    public UserDto getOne(@PathVariable String id){
        Optional<User> user = userService.findById(id);
        return user.map(userMapper::UserToUserDto).orElse(null);
    }

    @GetMapping("/username/{username}")
    public UserDto getOneByUsername(@PathVariable String username){
        return userMapper.UserToUserDto(userService.findByUsername(username));
    }

    @GetMapping("/searchPublicUsers/{searchTerm}")
    public List<UserDto> searchPublicUsers(@PathVariable String searchTerm){
        List<UserDto> userDtoList = new ArrayList<>();
        for(User user : userService.searchPublicUsers(searchTerm))
            userDtoList.add(userMapper.UserToUserDto(user));
        return userDtoList;
    }

    @DeleteMapping("/{id}")
    public String deleteUser(@PathVariable String id){
        userService.deleteById(id);
        return "User deleted with id " + id;
    }

    @PutMapping("/interests")
    public String updateInterests(@RequestBody UpdateInterestsDto updateInterestsDto){
        return "Updated user with id " + userService.updateInterests(updateInterestsDto).getId();
    }

    @PutMapping("/skills")
    public String updateSkills(@RequestBody UpdateSkillsDto updateSkillsDto){
        return "Updated user with id " + userService.updateSkills(updateSkillsDto).getId();
    }

    @PutMapping("/education")
    public String updateEducation(@RequestBody UpdateEducationDto updateEducationDto){
        return "Updated user with id " + userService.updateEducation(updateEducationDto).getId();
    }

    @PutMapping("/workExperience")
    public String updateWorkExperience(@RequestBody UpdateWorkExperienceDto updateWorkExperienceDto){
        return "Updated work experience for user with id " + userService.updateWorkExperience(updateWorkExperienceDto).getId();
    }
}

package com.example.usermicroservice.mapper;

import com.example.usermicroservice.dto.AddUserDto;
import com.example.usermicroservice.dto.UpdateUserDto;
import com.example.usermicroservice.dto.UserDto;
import com.example.usermicroservice.model.User;

public class UserMapper {

    public User AddUserDtoToUser(AddUserDto userDto){
        return new User(userDto.getId(), userDto.getUsername(), userDto.getPassword(), userDto.getName(),
                userDto.getEmail(), userDto.getTelephoneNo(), userDto.getGender(), userDto.getDateOfBirth(),
                userDto.getBiography(), userDto.isPrivate(), userDto.getEducationList(),
                userDto.getWorkExperienceList(), userDto.getInterests(), userDto.getSkills());
    }
    public UserDto UserToUserDto(User user){
        return new UserDto(user.getId(), user.getUsername(), user.getName(), user.getEmail(),
                user.getTelephoneNo(), user.getGender(), user.getDateOfBirth(), user.getBiography(), user.isPrivate(),
                user.getEducation(), user.getWorkExperience(), user.getInterests(), user.getSkills());
    }

    public User UpdateUserDtoToUser(UpdateUserDto userDto) {
        return new User(userDto.getId(), userDto.getUsername(), userDto.getPassword(), userDto.getName(),
                userDto.getEmail(), userDto.getTelephoneNo(), userDto.getGender(), userDto.getDateOfBirth(),
                userDto.getBiography(), userDto.isPrivate(), userDto.getEducationList(),
                userDto.getWorkExperienceList(), userDto.getInterests(), userDto.getSkills());
    }
}

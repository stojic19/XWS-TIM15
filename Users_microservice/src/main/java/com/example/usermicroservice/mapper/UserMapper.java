package com.example.usermicroservice.mapper;

import com.example.usermicroservice.dto.UserDto;
import com.example.usermicroservice.model.User;

public class UserMapper {

    public User UserDtoToUser(UserDto userDto){
        return new User(userDto.getId(), userDto.getUsername(), userDto.getPassword(), userDto.getName(), userDto.getEmail(), userDto.getTelephoneNo(), userDto.getGender(), userDto.getDateOfBirth(), userDto.getBiography());
    }
    public UserDto UserToUserDto(User user){
        return new UserDto(user.getId(), user.getUsername(), user.getPassword(), user.getName(), user.getEmail(), user.getTelephoneNo(), user.getGender(), user.getDateOfBirth(), user.getBiography());
    }
}

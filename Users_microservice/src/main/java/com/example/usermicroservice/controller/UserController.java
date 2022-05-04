package com.example.usermicroservice.controller;

import com.example.usermicroservice.dto.UpdateUserDto;
import com.example.usermicroservice.dto.UserDto;
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

    @PostMapping("/addUser")
    public String saveUser(@RequestBody UserDto userDto){
        if(isNullOrEmpty(userDto.getUsername(), userDto.getPassword(), userDto.getName(), userDto.getEmail(), userDto.getTelephoneNo(), userDto.getGender().toString(), userDto.getDateOfBirth().toString(), userDto.getBiography()))
            return "None of fields cannot be empty!";
        if(userService.usernameExists(userDto.getUsername()))
            return "Username already exists!";
        return "Added user with id " + userService.addUser(userMapper.UserDtoToUser(userDto)).getId();
    }

    @PostMapping("/updateUser")
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

    @DeleteMapping("/{id}")
    public String deleteUser(@PathVariable String id){
        userService.deleteById(id);
        return "User deleted with id " + id;
    }

}

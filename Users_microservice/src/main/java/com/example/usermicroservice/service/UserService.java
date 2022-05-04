package com.example.usermicroservice.service;

import com.example.usermicroservice.model.User;

import java.util.List;
import java.util.Optional;

public interface UserService {

    boolean usernameExists(String username);
    User addUser(User user);
    List<User> getUsers();
    Optional<User> findById(String id);
    void deleteById(String id);
}

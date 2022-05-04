package com.example.usermicroservice.service.Impl;

import com.example.usermicroservice.model.User;
import com.example.usermicroservice.repository.UserRepository;
import com.example.usermicroservice.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

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
    public void deleteById(String id) {
        userRepository.deleteById(id);
    }

    @Override
    public User updateUser(User userWithNewData) {
        return userRepository.save(userWithNewData);
    }
}

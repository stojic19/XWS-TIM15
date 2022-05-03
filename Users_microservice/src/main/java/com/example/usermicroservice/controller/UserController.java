package com.example.usermicroservice.controller;

import com.example.usermicroservice.model.User;
import com.example.usermicroservice.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.swing.text.html.Option;
import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/users")
public class UserController {
    @Autowired
    private UserRepository repository;

    @PostMapping("/addUser")
    public String saveUser(@RequestBody User user){
        repository.save(user);
        return "Added book with id " + user.getId();
    }

    @GetMapping
    public List<User> getUsers(){
        return repository.findAll();
    }

    @GetMapping("/{id}")
    public Optional<User> getOne(@PathVariable int id){
        return repository.findById(id);
    }

    @DeleteMapping("/{id}")
    public String deleteUser(@PathVariable int id){
        repository.deleteById(id);
        return "User deleted with id " + id;
    }

}

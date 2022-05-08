package com.example.usermicroservice.repository;

import com.example.usermicroservice.model.User;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;

import java.util.List;

public interface UserRepository extends MongoRepository<User, String> {

    @Query("{username:'?0'}")
    User findUserByUsername(String username);
    List<User> findUsersByIsPrivateFalse();
}

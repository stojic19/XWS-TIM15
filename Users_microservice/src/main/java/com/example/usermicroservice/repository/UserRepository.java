package com.example.usermicroservice.repository;

import com.example.usermicroservice.model.User;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface UserRepository extends MongoRepository<User, Integer> {
}

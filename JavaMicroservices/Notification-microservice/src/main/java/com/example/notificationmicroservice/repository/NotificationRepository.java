package com.example.notificationmicroservice.repository;

import com.example.notificationmicroservice.model.Notification;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.List;

public interface NotificationRepository extends MongoRepository<Notification, String> {

    List<Notification> findByFollowerId(String id);
    List<Notification> findByUserId(String id);
}

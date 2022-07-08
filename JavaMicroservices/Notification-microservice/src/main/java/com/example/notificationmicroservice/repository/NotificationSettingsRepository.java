package com.example.notificationmicroservice.repository;

import com.example.notificationmicroservice.model.Notification;
import com.example.notificationmicroservice.model.NotificationSettings;
import com.example.notificationmicroservice.service.NotificationSettingsService;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface NotificationSettingsRepository extends MongoRepository<NotificationSettings, String> {
    NotificationSettings findByUserId(String id);
    void deleteByUserId(String id);
}

package com.example.notificationmicroservice.service;

import com.example.notificationmicroservice.model.Notification;
import java.util.List;
import java.util.Optional;

public interface NotificationService {
    Notification addNotification(Notification notification);
    List<Notification> getNotifications();
    Optional<Notification> findById(String id);
    List<Notification> findByFollowerId(String id);
    List<Notification> findByUserId(String id);
    void deleteById(String id);
}

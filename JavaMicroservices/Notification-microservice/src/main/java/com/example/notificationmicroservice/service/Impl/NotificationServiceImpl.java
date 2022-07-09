package com.example.notificationmicroservice.service.Impl;
import com.example.notificationmicroservice.model.Notification;
import com.example.notificationmicroservice.repository.NotificationRepository;
import com.example.notificationmicroservice.service.NotificationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class NotificationServiceImpl implements NotificationService {

    @Autowired
    private NotificationRepository repository;

    @Override
    public Notification addNotification(Notification notification) {
        notification.setId(UUID.randomUUID().toString());
        repository.save(notification);
        return notification;
    }

    @Override
    public List<Notification> getNotifications() {
        return repository.findAll();
    }

    @Override
    public Optional<Notification> findById(String id) {
        return repository.findById(id);
    }

    @Override
    public List<Notification> findByFollowerId(String id) {
        List<Notification> list = repository.findByFollowerId(id);
        list.sort(Comparator.comparing(Notification::getTime).reversed());
        return list;
    }

    @Override
    public List<Notification> findByUserId(String id) {
        List<Notification> list = repository.findByUserId(id);
        list.sort(Comparator.comparing(Notification::getTime).reversed());
        return list;
    }

    @Override
    public void deleteById(String id) {
        repository.deleteById(id);
    }
}

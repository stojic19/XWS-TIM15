package com.example.notificationmicroservice.service.Impl;
import com.example.notificationmicroservice.model.Notification;
import com.example.notificationmicroservice.model.NotificationSettings;
import com.example.notificationmicroservice.model.NotificationType;
import com.example.notificationmicroservice.repository.NotificationRepository;
import com.example.notificationmicroservice.repository.NotificationSettingsRepository;
import com.example.notificationmicroservice.service.NotificationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class NotificationServiceImpl implements NotificationService {

    @Autowired
    private NotificationRepository repository;
    @Autowired
    private NotificationSettingsRepository settingsRepository;

    @Override
    public Notification addNotification(Notification notification) {

        if(checkNotificationSettings(notification)){
            notification.setId(UUID.randomUUID().toString());
            repository.save(notification);
            return notification;
        }
        return new Notification();
    }

    @Override
    public Notification addNotificationForNewPost(Notification notification) {
        if(notification.getType() == NotificationType.profile){
            List<NotificationSettings> settingsForUsers = settingsRepository.findAll();

            for (NotificationSettings settings: settingsForUsers) {
                for (String s : settings.getFollowerIdsForPosts()) {
                    if(s.equals(notification.getFollowerId())) {
                        Notification newNotification = new Notification();
                        newNotification.setType(notification.getType());
                        newNotification.setUserId(settings.getUserId());
                        newNotification.setAction(notification.getAction());
                        newNotification.setPostId(notification.getPostId());
                        newNotification.setFollowerId(notification.getFollowerId());
                        newNotification.setTime(notification.getTime());
                        repository.save(newNotification);
                        break;
                    }
                }
            }
            return notification;
        }
        return new Notification();
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

    public Boolean checkNotificationSettings(Notification notification){

        NotificationSettings settings = settingsRepository.findByUserId(notification.getUserId());

        if(notification.getType() == NotificationType.message){
            List<String> followers = settings.getFollowerIdsForMessages(); //pratioci za koje su iskljucene notifikacije za poruke
            for (String s: followers) {
                if(s == notification.getFollowerId())
                    return false;
            }
            return true;
        }
        else if (notification.getType() == NotificationType.post){
            System.out.println(settings.getGetNotificationsForMyPosts());
            if(settings.getGetNotificationsForMyPosts()){
                return true;
            }
            return false;
        }
        else{
            return false;
        }
    }

}

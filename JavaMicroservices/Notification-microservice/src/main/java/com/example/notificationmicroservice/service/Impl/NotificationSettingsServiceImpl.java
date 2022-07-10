package com.example.notificationmicroservice.service.Impl;

import com.example.notificationmicroservice.dto.UpdateSettingsForMessagesDto;
import com.example.notificationmicroservice.dto.UpdateSettingsForPostsDto;
import com.example.notificationmicroservice.model.NotificationSettings;
import com.example.notificationmicroservice.repository.NotificationSettingsRepository;
import com.example.notificationmicroservice.service.NotificationSettingsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.UUID;

@Service
public class NotificationSettingsServiceImpl implements NotificationSettingsService {

    @Autowired
    private NotificationSettingsRepository repository;


    @Override
    public boolean settingsForUserExists(String userId) {
        return repository.findByUserId(userId) != null;
    }

    @Override
    public NotificationSettings addNotificationSettings(NotificationSettings settings) {
        if(!settingsForUserExists(settings.getUserId())){
            settings.setId(UUID.randomUUID().toString());
            settings.setFollowerIdsForMessages(settings.getFollowerIdsForMessages());
            settings.setFollowerIdsForPosts(settings.getFollowerIdsForPosts());
            settings.setGetNotificationsForMyPosts(settings.getGetNotificationsForMyPosts());

            repository.save(settings);
            return settings;
        }
        return null;
    }

    @Override
    public NotificationSettings findByUserId(String id) {
        if(settingsForUserExists(id))
            return repository.findByUserId(id);
        return repository.save(new NotificationSettings(UUID.randomUUID().toString(), id, new ArrayList<>(), new ArrayList<>(), true));
    }

    @Override
    public void deleteByUserId(String id) {
        repository.deleteByUserId(id);
    }

    @Override
    public NotificationSettings updateSettings(NotificationSettings settingsWithNewData) {
        NotificationSettings notificationSettings = findByUserId(settingsWithNewData.getUserId());
        notificationSettings.setUserId(notificationSettings.getUserId());
        notificationSettings.setFollowerIdsForMessages(notificationSettings.getFollowerIdsForMessages());
        notificationSettings.setFollowerIdsForPosts(notificationSettings.getFollowerIdsForPosts());
        notificationSettings.setGetNotificationsForMyPosts(notificationSettings.getGetNotificationsForMyPosts());
        return repository.save(settingsWithNewData);
    }

    @Override
    public NotificationSettings updateSettingsForPosts(UpdateSettingsForPostsDto dto) {
        NotificationSettings notificationSettings = findByUserId(dto.getUserId());
        notificationSettings.setFollowerIdsForPosts(dto.getFollowerIdsForPosts());
        return repository.save(notificationSettings);
    }

    @Override
    public NotificationSettings updateSettingsForMessages(UpdateSettingsForMessagesDto dto) {
        NotificationSettings notificationSettings = findByUserId(dto.getUserId());
        notificationSettings.setFollowerIdsForMessages(dto.getFollowerIdsForMessages());
        return repository.save(notificationSettings);
    }
}

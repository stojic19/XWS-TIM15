package com.example.notificationmicroservice.service;

import com.example.notificationmicroservice.dto.UpdateSettingsForMessagesDto;
import com.example.notificationmicroservice.dto.UpdateSettingsForPostsDto;
import com.example.notificationmicroservice.model.NotificationSettings;

public interface NotificationSettingsService {
    boolean settingsForUserExists(String userId);
    NotificationSettings addNotificationSettings(NotificationSettings settings);
    NotificationSettings findByUserId(String id);
    void deleteByUserId(String id);
    NotificationSettings updateSettings(NotificationSettings settingsWithNewData);
    NotificationSettings updateSettingsForPosts(UpdateSettingsForPostsDto dto);
    NotificationSettings updateSettingsForMessages(UpdateSettingsForMessagesDto dto);
}

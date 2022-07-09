package com.example.notificationmicroservice.mapper;

import com.example.notificationmicroservice.dto.NotificationDto;
import com.example.notificationmicroservice.dto.NotificationSettingsDto;
import com.example.notificationmicroservice.model.Notification;
import com.example.notificationmicroservice.model.NotificationAction;
import com.example.notificationmicroservice.model.NotificationSettings;
import com.example.notificationmicroservice.model.NotificationType;

public class NotificationMapper {

    public NotificationDto NotificationToDto(Notification notification){
        return new NotificationDto(notification.getId(), notification.getType().toString(), notification.getAction().toString(),
                notification.getUserId(), notification.getFollowerId(),notification.getPostId(), notification.getMessagesId(),notification.getTime());
    }

    public Notification DtoToNotification(NotificationDto dto) {
        return new Notification(dto.getId(), NotificationType.valueOf(dto.getType()), NotificationAction.valueOf(dto.getAction()),
                dto.getUserId(),dto.getFollowerId(), dto.getPostId(), dto.getMessagesId(), dto.getTime());
    }

    public NotificationSettings DtoToNotificationSettings(NotificationSettingsDto dto) {
        return new NotificationSettings(dto.getId(),dto.getUserId(), dto.getFollowerIdsForPosts(),
                                        dto.getFollowerIdsForMessages(), dto.getGetNotificationsForMyPosts());
    }

    public NotificationSettingsDto NotificationSettingsToDto(NotificationSettings settings) {
        return new NotificationSettingsDto(settings.getId(),settings.getUserId(), settings.getFollowerIdsForPosts(),
                settings.getFollowerIdsForMessages(), settings.getGetNotificationsForMyPosts());
    }
}

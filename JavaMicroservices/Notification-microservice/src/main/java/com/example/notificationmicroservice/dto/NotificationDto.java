package com.example.notificationmicroservice.dto;

import com.example.notificationmicroservice.model.NotificationType;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.Date;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class NotificationDto {

    private String id;
    private String type;
    private String action;
    private String userId;
    private String followerId;
    private String postId;
    private String messagesId;
    private Date time;
}

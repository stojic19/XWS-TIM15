package com.example.notificationmicroservice.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.Date;
import java.util.List;
import java.util.UUID;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Document(collection = "Notification")
public class Notification {
    @Id
    private String id;
    private NotificationType type;
    private NotificationAction action;
    private String followerId;
    private Date time;


//    private UUID apiKey;

}

package com.example.notificationmicroservice.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.List;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Document(collection = "NotificationSettings")
public class NotificationSettings {

    @Id
    private String id;
    private String userId;
    private List<String> followerIdsForPosts; //ukljucene notifikacije za svaki novi post pratioca
    private List<String> followerIdsForMessages; //iskljucene notifikacije za poruke od pratioca
    private Boolean getNotificationsForMyPosts;
}

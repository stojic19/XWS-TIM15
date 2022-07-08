package com.example.notificationmicroservice.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.List;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class NotificationSettingsDto {

    private String id;
    private String userId;
    private List<String> followerIdsForPosts; //ukljucene notifikacije za svaki novi post pratioca
    private List<String> followerIdsForMessages; //iskljucene notifikacije za poruke od pratioca
    private Boolean getNotificationsForMyPosts;
}

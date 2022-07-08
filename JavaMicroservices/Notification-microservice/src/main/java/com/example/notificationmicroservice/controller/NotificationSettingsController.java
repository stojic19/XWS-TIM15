package com.example.notificationmicroservice.controller;

import com.example.notificationmicroservice.dto.NotificationSettingsDto;
import com.example.notificationmicroservice.dto.UpdateSettingsForMessagesDto;
import com.example.notificationmicroservice.dto.UpdateSettingsForPostsDto;
import com.example.notificationmicroservice.mapper.NotificationMapper;
import com.example.notificationmicroservice.service.NotificationSettingsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;


@RestController
@RequestMapping("/notificationSettings")
public class NotificationSettingsController {
    @Autowired
    private NotificationSettingsService service;
    final private NotificationMapper mapper = new NotificationMapper();

    @PostMapping()
    public String saveNotificationSettings(@RequestBody NotificationSettingsDto dto){
        if(service.settingsForUserExists(dto.getUserId()))
            return "Settings for user already exists!";
        else
            return "Added notification settings with id " + service.addNotificationSettings(mapper.DtoToNotificationSettings(dto)).getId();
    }

    @GetMapping("/user/{id}")
    public NotificationSettingsDto getOneByUserId(@PathVariable String id){
        return mapper.NotificationSettingsToDto(service.findByUserId(id));
    }

    @DeleteMapping("/user/{id}")
    public String deleteNotificationSettings(@PathVariable String id){
        service.deleteByUserId(id);
        return "Notification settings deleted for user with id " + id;
    }

    @PutMapping()
    public String updateSettings(@RequestBody NotificationSettingsDto dto){
        if(!service.settingsForUserExists(dto.getUserId()))
            return "Settings for user doesn't exist!";
        else if (dto.getId() == null)
            return "Settings id is missing!";
        else
            return "Updated user with id " + service.updateSettings(mapper.DtoToNotificationSettings(dto)).getId();
    }

    @PutMapping("/messages")
    public String updateMessagesSettings(@RequestBody UpdateSettingsForMessagesDto dto){
        return "Updated messages settings for user with id " + service.updateSettingsForMessages(dto).getUserId();
    }
    @PutMapping("/posts")
    public String updatePostsSettings(@RequestBody UpdateSettingsForPostsDto dto){
        return "Updated posts settings for user with id " + service.updateSettingsForPosts(dto).getUserId();
    }
}

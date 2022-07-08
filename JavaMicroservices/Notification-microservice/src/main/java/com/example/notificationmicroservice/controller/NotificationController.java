package com.example.notificationmicroservice.controller;

import com.example.notificationmicroservice.dto.*;
import com.example.notificationmicroservice.mapper.NotificationMapper;
import com.example.notificationmicroservice.model.Notification;
import com.example.notificationmicroservice.service.NotificationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/notifications")
public class NotificationController {
    @Autowired
    private NotificationService service;

    final private NotificationMapper mapper = new NotificationMapper();

    @PostMapping()
    public String saveNotification(@RequestBody NotificationDto dto){
        if(isNullOrEmpty(dto.getType(), dto.getAction(), dto.getFollowerId(), dto.getTime().toString()))
            return "None of fields cannot be empty!";
        return "Added notification with id " + service.addNotification(mapper.DtoToNotification(dto)).getId();
    }

    private static boolean isNullOrEmpty(String... strArr){
        for (String st : strArr) {
            if  (st==null || st.equals(""))
                return true;

        }
        return false;
    }

    @GetMapping
    public List<NotificationDto> getNotifications(){
        List<NotificationDto> dtoList = new ArrayList<>();
        for(Notification notification : service.getNotifications())
            dtoList.add(mapper.NotificationToDto(notification));
        return dtoList;
    }

    @GetMapping("/{id}")
    public NotificationDto getOne(@PathVariable String id){
        Optional<Notification> notification = service.findById(id);
        return notification.map(mapper::NotificationToDto).orElse(null);
    }

    @GetMapping("/follower/{id}")
    public List<NotificationDto> getAllByFollowerId(@PathVariable String id){
        List<NotificationDto> dtoList = new ArrayList<>();
        for(Notification notification : service.findByFollowerId(id))
            dtoList.add(mapper.NotificationToDto(notification));
        return dtoList;
    }

    @DeleteMapping("/{id}")
    public String deleteNotification(@PathVariable String id){
        service.deleteById(id);
        return "Notification deleted with id " + id;
    }

}

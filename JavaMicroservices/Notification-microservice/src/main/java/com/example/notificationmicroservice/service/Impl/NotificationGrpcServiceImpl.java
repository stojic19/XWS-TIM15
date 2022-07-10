package com.example.notificationmicroservice.service.Impl;

import com.example.notificationmicroservice.*;
import com.example.notificationmicroservice.NotificationsServiceGrpc;
import com.example.notificationmicroservice.dto.NotificationDto;
import com.example.notificationmicroservice.dto.NotificationSettingsDto;
import com.example.notificationmicroservice.dto.UpdateSettingsForMessagesDto;
import com.example.notificationmicroservice.dto.UpdateSettingsForPostsDto;
import com.example.notificationmicroservice.mapper.NotificationMapper;
import com.example.notificationmicroservice.model.Notification;
import com.example.notificationmicroservice.model.NotificationType;
import com.example.notificationmicroservice.service.NotificationService;
import com.example.notificationmicroservice.service.NotificationSettingsService;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Optional;

@GrpcService
public class NotificationGrpcServiceImpl extends  NotificationsServiceGrpc.NotificationsServiceImplBase{

    @Autowired
    private NotificationService notificationService;
    @Autowired
    private NotificationSettingsService notificationSettingsService;

    final private NotificationMapper mapper = new NotificationMapper();

    @Override
    public void saveNotification(SaveNotificationRequest request, StreamObserver<StringResponse> responseObserver) {
        StringResponse response;
        if(isNullOrEmpty(request.getNotification().getType()))
            response = StringResponse.newBuilder().setResponse("None of fields cannot be empty!").setStatus(400).build();
        else{
            if(request.getNotification().getType().equals(NotificationType.profile.toString()))
                response = StringResponse.newBuilder().setResponse("Added notification with id " + notificationService.addNotificationForNewPost(mapper.DtoToNotification(new NotificationDto(request.getNotification().getId(), request.getNotification().getType(), request.getNotification().getAction(), request.getNotification().getUserId(), request.getNotification().getFollowerId(), request.getNotification().getPostId(), request.getNotification().getMessagesId(), new Date()))).getId()).setStatus(200).build();
            else
                response = StringResponse.newBuilder().setResponse("Added notification with id " + notificationService.addNotification(mapper.DtoToNotification(new NotificationDto(request.getNotification().getId(), request.getNotification().getType(), request.getNotification().getAction(), request.getNotification().getUserId(), request.getNotification().getFollowerId(), request.getNotification().getPostId(), request.getNotification().getMessagesId(), new Date()))).getId()).setStatus(200).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    private static boolean isNullOrEmpty(String... strArr){
        for (String st : strArr) {
            if  (st==null || st.equals(""))
                return true;

        }
        return false;
    }

    public static String getDateString(Date date) {
        String strDateFormat = "dd.MM.yyyy. HH:mm:ss";
        DateFormat dateFormat = new SimpleDateFormat(strDateFormat);
        return dateFormat.format(date);
    }

    @Override
    public void getNotifications(GetNotificationsRequest request, StreamObserver<GetNotificationsResponse> responseObserver) {
        responseObserver.onNext(GetNotificationsResponse.newBuilder().addAllNotifications(getNotificationsForResponse(notificationService.getNotifications())).build());
        responseObserver.onCompleted();
    }

    @Override
    public void getOneNotification(GetByIdRequest request, StreamObserver<GetOneNotificationResponse> responseObserver) {
        GetOneNotificationResponse response;
        Optional<Notification> notification = notificationService.findById(request.getId());
        if(notification.isPresent()){
            response = GetOneNotificationResponse.newBuilder().setNotification(getProtoNotification(notification.get())).build();
        }else{
            response = GetOneNotificationResponse.newBuilder().setNotification(com.example.notificationmicroservice.Notification.newBuilder()).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getAllNotificationsByFollowerId(GetByIdRequest request, StreamObserver<GetNotificationsResponse> responseObserver) {
        responseObserver.onNext(GetNotificationsResponse.newBuilder().addAllNotifications(getNotificationsForResponse(notificationService.findByFollowerId(request.getId()))).build());
        responseObserver.onCompleted();
    }

    @Override
    public void getAllNotificationsByUserId(GetByIdRequest request, StreamObserver<GetNotificationsResponse> responseObserver) {
        responseObserver.onNext(GetNotificationsResponse.newBuilder().addAllNotifications(getNotificationsForResponse(notificationService.findByUserId(request.getId()))).build());
        responseObserver.onCompleted();
    }

    private List<com.example.notificationmicroservice.Notification> getNotificationsForResponse(List<Notification> notificationsFromService){
        List<com.example.notificationmicroservice.Notification> notifications = new ArrayList<com.example.notificationmicroservice.Notification>();
        for(Notification notification: notificationsFromService){
            notifications.add(getProtoNotification(notification));
        }
        return notifications;
    }

    private com.example.notificationmicroservice.Notification getProtoNotification(Notification notification){
        return com.example.notificationmicroservice.Notification.newBuilder()
                .setId(notification.getId())
                .setAction(notification.getAction().toString())
                .setFollowerId(notification.getFollowerId())
                .setTime(getDateString(notification.getTime()))
                .setType(notification.getType().toString())
                .setMessagesId(notification.getMessagesId())
                .setUserId(notification.getUserId())
                .setPostId(notification.getPostId())
                .build();
    }

    @Override
    public void deleteNotification(DeleteByIdRequest request, StreamObserver<StringResponse> responseObserver) {
        notificationService.deleteById(request.getId());
        responseObserver.onNext(com.example.notificationmicroservice.StringResponse.newBuilder().setResponse("User deleted with id " + request.getId()).setStatus(200).build());
        responseObserver.onCompleted();
    }

    @Override
    public void saveNotificationSettings(SaveNotificationSettingsRequest request, StreamObserver<StringResponse> responseObserver) {
        StringResponse response;
        if(notificationSettingsService.settingsForUserExists(request.getNotificationSettings().getUserId()))
            response = StringResponse.newBuilder().setResponse("Settings for user already exists!").setStatus(400).build();
        else
            response = StringResponse.newBuilder().setResponse("Added notification settings with id " + notificationSettingsService.addNotificationSettings(mapper.DtoToNotificationSettings(new NotificationSettingsDto(request.getNotificationSettings().getId(), request.getNotificationSettings().getUserId(), request.getNotificationSettings().getFollowerIdsForPostsList(), request.getNotificationSettings().getFollowerIdsForMessagesList(),request.getNotificationSettings().getGetNotificationsForMyPosts()))).getId()).setStatus(200).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getSettingsByUserId(GetByIdRequest request, StreamObserver<GetSettingsByUserIdResponse> responseObserver) {
        responseObserver.onNext(GetSettingsByUserIdResponse.newBuilder().setNotificationSettings(getNotificationSettingsProto(notificationSettingsService.findByUserId(request.getId()))).build());
        responseObserver.onCompleted();
    }

    private NotificationSettings getNotificationSettingsProto(com.example.notificationmicroservice.model.NotificationSettings notificationSettings) {
        return com.example.notificationmicroservice.NotificationSettings.newBuilder()
                .setId(notificationSettings.getId())
                .setGetNotificationsForMyPosts(notificationSettings.getGetNotificationsForMyPosts())
                .setUserId(notificationSettings.getUserId())
                .addAllFollowerIdsForMessages(notificationSettings.getFollowerIdsForMessages())
                .addAllFollowerIdsForPosts(notificationSettings.getFollowerIdsForPosts())
                .build();
    }

    @Override
    public void deleteNotificationSettings(DeleteByIdRequest request, StreamObserver<StringResponse> responseObserver) {
        notificationSettingsService.deleteByUserId(request.getId());
        responseObserver.onNext(StringResponse.newBuilder().setResponse("Notification settings deleted for user with id " + request.getId()).setStatus(200).build());
        responseObserver.onCompleted();
    }

    @Override
    public void updateSettings(UpdateSettingsRequest request, StreamObserver<StringResponse> responseObserver) {
        if(!notificationSettingsService.settingsForUserExists(request.getNotificationSettings().getUserId())){
            responseObserver.onNext(StringResponse.newBuilder().setResponse("Settings for user doesn't exist!").setStatus(400).build());
            responseObserver.onCompleted();
        }
        else if (request.getNotificationSettings().getId() == null){
            responseObserver.onNext(StringResponse.newBuilder().setResponse("Settings id is missing!").setStatus(400).build());
            responseObserver.onCompleted();
        }
        else{
            responseObserver.onNext(StringResponse.newBuilder().setResponse("Updated user with id " + notificationSettingsService.updateSettings(mapper.DtoToNotificationSettings(new NotificationSettingsDto(request.getNotificationSettings().getId(), request.getNotificationSettings().getUserId(), request.getNotificationSettings().getFollowerIdsForPostsList(), request.getNotificationSettings().getFollowerIdsForMessagesList(),request.getNotificationSettings().getGetNotificationsForMyPosts()))).getId()).setStatus(400).build());
            responseObserver.onCompleted();
        }
    }

    @Override
    public void updateMessagesSettings(UpdateMessagesSettingsRequest request, StreamObserver<StringResponse> responseObserver) {
        responseObserver.onNext(StringResponse.newBuilder().setResponse("Updated messages settings for user with id " + notificationSettingsService.updateSettingsForMessages(new UpdateSettingsForMessagesDto(request.getUserId(), request.getFollowerIdsForMessagesList())).getUserId()).setStatus(200).build());
        responseObserver.onCompleted();
    }

    @Override
    public void updatePostsSettings(UpdatePostsSettingsRequest request, StreamObserver<StringResponse> responseObserver) {
        responseObserver.onNext(StringResponse.newBuilder().setResponse("Updated posts settings for user with id " + notificationSettingsService.updateSettingsForPosts(new UpdateSettingsForPostsDto(request.getUserId(), request.getFollowerIdsForPostsList())).getUserId()).setStatus(200).build());
        responseObserver.onCompleted();
    }
}

package com.example.notificationmicroservice.service.Impl;

import com.example.notificationmicroservice.*;
import com.example.notificationmicroservice.NotificationsServiceGrpc;
import io.grpc.stub.StreamObserver;

public class NotificationGrpcServiceImpl extends  NotificationsServiceGrpc.NotificationsServiceImplBase{

    @Override
    public void saveNotification(SaveNotificationRequest request, StreamObserver<StringResponse> responseObserver) {
        super.saveNotification(request, responseObserver);
    }

    @Override
    public void getNotifications(GetNotificationsRequest request, StreamObserver<GetNotificationsResponse> responseObserver) {
        super.getNotifications(request, responseObserver);
    }

    @Override
    public void getOneNotification(GetByIdRequest request, StreamObserver<GetOneNotificationResponse> responseObserver) {
        super.getOneNotification(request, responseObserver);
    }

    @Override
    public void getAllNotificationsByFollowerId(GetByIdRequest request, StreamObserver<GetNotificationsResponse> responseObserver) {
        super.getAllNotificationsByFollowerId(request, responseObserver);
    }

    @Override
    public void deleteNotification(DeleteByIdRequest request, StreamObserver<StringResponse> responseObserver) {
        super.deleteNotification(request, responseObserver);
    }

    @Override
    public void saveNotificationSettings(SaveNotificationSettingsRequest request, StreamObserver<StringResponse> responseObserver) {
        super.saveNotificationSettings(request, responseObserver);
    }

    @Override
    public void getSettingsByUserId(GetByIdRequest request, StreamObserver<GetSettingsByUserIdResponse> responseObserver) {
        super.getSettingsByUserId(request, responseObserver);
    }

    @Override
    public void deleteNotificationSettings(DeleteByIdRequest request, StreamObserver<StringResponse> responseObserver) {
        super.deleteNotificationSettings(request, responseObserver);
    }

    @Override
    public void updateSettings(UpdateSettingsRequest request, StreamObserver<StringResponse> responseObserver) {
        super.updateSettings(request, responseObserver);
    }

    @Override
    public void updateMessagesSettings(UpdateMessagesSettingsRequest request, StreamObserver<StringResponse> responseObserver) {
        super.updateMessagesSettings(request, responseObserver);
    }

    @Override
    public void updatePostsSettings(UpdatePostsSettingsRequest request, StreamObserver<StringResponse> responseObserver) {
        super.updatePostsSettings(request, responseObserver);
    }
}

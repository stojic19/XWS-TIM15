package com.example.usermicroservice.service.Impl;

import com.example.usermicroservice.*;
import com.example.usermicroservice.model.User;
import com.example.usermicroservice.repository.UserRepository;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.*;

@GrpcService
public class UserServiceImpl extends UsersServiceGrpc.UsersServiceImplBase {

    @Autowired
    private UserRepository userRepository;

    @Override
    public void addUser(AddUserRequest request, StreamObserver<AddUserResponse> responseObserver) {
        AddUserResponse response;
        if(userRepository.findUserByUsername(request.getUsername()) != null)
            response = AddUserResponse.newBuilder().setResponse("false").build();
        else if(isNullOrEmpty(request.getUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), request.getDateOfBirth().toString(), request.getBiography()))
            response = AddUserResponse.newBuilder().setResponse("None of fields cannot be empty!").build();
        else{
            Date dateOfBirth = new Date();
            try {
                dateOfBirth = new SimpleDateFormat("dd/MM/yyyy").parse(request.getDateOfBirth());
            } catch (ParseException e) {
                e.printStackTrace();
            }
            response = AddUserResponse.newBuilder().setResponse("Added user with id " + userRepository.save(new User(UUID.randomUUID().toString(), request.getUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), dateOfBirth, request.getBiography())).getId()).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void updateUser(UpdateUserRequest request, StreamObserver<UpdateUserResponse> responseObserver) {
        UpdateUserResponse response;
        if(isNullOrEmpty(request.getUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), request.getDateOfBirth().toString(), request.getBiography()))
            response = UpdateUserResponse.newBuilder().setResponse("None of fields cannot be empty!").build();
        else{
            Date dateOfBirth = new Date();
            try {
                dateOfBirth = new SimpleDateFormat("dd/MM/yyyy").parse(request.getDateOfBirth());
            } catch (ParseException e) {
                e.printStackTrace();
            }
            if(!isNullOrEmpty(request.getNewUsername())&&!request.getUsername().equals(request.getNewUsername())) {
                if(userRepository.findUserByUsername(request.getNewUsername()) != null)
                    response = UpdateUserResponse.newBuilder().setResponse("Username already exists!").build();
                else
                {
                    User user = userRepository.findUserByUsername(request.getUsername());
                    response = UpdateUserResponse.newBuilder().setResponse("Updated user with id " + userRepository.save(new User(user.getId(), request.getNewUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), dateOfBirth, request.getBiography())).getId()).build();
                }
            }
            else{
                User user = userRepository.findUserByUsername(request.getUsername());
                response = UpdateUserResponse.newBuilder().setResponse("Updated user with id " + userRepository.save(new User(user.getId(), request.getUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), dateOfBirth, request.getBiography())).getId()).build();
            }
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getUser(GetUserRequest request, StreamObserver<GetUserResponse> responseObserver) {
        GetUserResponse response;
        Optional<User> user = userRepository.findById(request.getId());
        if(user.isPresent()){
            User presentUser = user.get();
            response = GetUserResponse.newBuilder().setUser(com.example.usermicroservice.User.newBuilder().
                    setUsername(presentUser.getUsername())
                    .setEmail(presentUser.getEmail())
                    .setBiography(presentUser.getBiography())
                    .setDateOfBirth(presentUser.getDateOfBirth().toString())
                    .setGender(presentUser.getGender())
                    .setId(presentUser.getId())
                    .setTelephoneNo(presentUser.getTelephoneNo())
                    .setName(presentUser.getName())
            ).build();
        }else{
            response = GetUserResponse.newBuilder().setUser(com.example.usermicroservice.User.newBuilder()).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getUsers(GetUsersRequest request, StreamObserver<GetUsersResponse> responseObserver) {
        GetUsersResponse response;
        List<User> savedUsers = userRepository.findAll();
        List<com.example.usermicroservice.User> users = new ArrayList<com.example.usermicroservice.User>();
        for(User user: savedUsers){
            com.example.usermicroservice.User protoUser = com.example.usermicroservice.User.newBuilder().setUsername(user.getUsername())
                    .setEmail(user.getEmail())
                    .setBiography(user.getBiography())
                    .setDateOfBirth(user.getDateOfBirth().toString())
                    .setGender(user.getGender())
                    .setId(user.getId())
                    .setTelephoneNo(user.getTelephoneNo())
                    .setName(user.getName()).build();
            users.add(protoUser);
        }
        response = GetUsersResponse.newBuilder().addAllUsers(users).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getUserByUsername(GetUserByUsernameRequest request, StreamObserver<GetUserResponse> responseObserver) {
        GetUserResponse response;
        User user = userRepository.findUserByUsername(request.getUsername());
        if(user != null){
            response = GetUserResponse.newBuilder().setUser(com.example.usermicroservice.User.newBuilder().
                    setUsername(user.getUsername())
                    .setEmail(user.getEmail())
                    .setBiography(user.getBiography())
                    .setDateOfBirth(user.getDateOfBirth().toString())
                    .setGender(user.getGender())
                    .setId(user.getId())
                    .setTelephoneNo(user.getTelephoneNo())
                    .setName(user.getName())
            ).build();
        }else{
            response = GetUserResponse.newBuilder().setUser(com.example.usermicroservice.User.newBuilder()).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void deleteUserById(GetUserRequest request, StreamObserver<StringResponse> responseObserver) {
        StringResponse stringResponse;
        if(userRepository.findById(request.getId()).isPresent()){
            userRepository.deleteById(request.getId());
            stringResponse = StringResponse.newBuilder().setResponse("User deleted with id " + request.getId()).build();
        }else{
            stringResponse = StringResponse.newBuilder().setResponse("No existing user with id " + request.getId() + "!").build();
        }
        responseObserver.onNext(stringResponse);
        responseObserver.onCompleted();
    }

    private static boolean isNullOrEmpty(String... strArr){
        for (String st : strArr) {
            if  (st==null || st.equals(""))
                return true;

        }
        return false;
    }
    /*
    public boolean usernameExists(String username) {
        return userRepository.findUserByUsername(username) != null;
    }

    public User addUser(User user) {
        user.setId(UUID.randomUUID().toString());
        userRepository.save(user);
        return user;
    }

    public List<User> getUsers() {
        return userRepository.findAll();
    }

    public Optional<User> findById(String id) {
        return userRepository.findById(id);
    }

    public void deleteById(String id) {
        userRepository.deleteById(id);
    }

    public User updateUser(User userWithNewData) {
        return userRepository.save(userWithNewData);
    }*/
}

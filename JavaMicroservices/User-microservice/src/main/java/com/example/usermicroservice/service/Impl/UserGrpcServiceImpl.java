package com.example.usermicroservice.service.Impl;

import com.example.usermicroservice.*;
import com.example.usermicroservice.model.User;
import com.example.usermicroservice.model.WorkExperience;
import com.example.usermicroservice.repository.UserRepository;
import com.example.usermicroservice.service.UserService;
import com.example.usermicroservice.token.JwtTokenUtil;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.*;

@GrpcService
public class UserGrpcServiceImpl extends UsersServiceGrpc.UsersServiceImplBase {

    @Autowired
    private UserRepository userRepository;
    @Autowired
    private UserService userService;
    @Autowired
    private JwtTokenUtil jwtTokenUtil;

    @Override
    public void addUser(AddUserRequest request, StreamObserver<AddUserResponse> responseObserver) {
        AddUserResponse response;
        if(userRepository.findUserByUsername(request.getUsername()) != null)
            response = AddUserResponse.newBuilder().setResponse("Username already exists!").build();
        else if(isNullOrEmpty(request.getUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), request.getDateOfBirth().toString(), request.getBiography()))
            response = AddUserResponse.newBuilder().setResponse("None of fields cannot be empty!").build();
        else{
            Date dateOfBirth = new Date();
            try {
                dateOfBirth = new SimpleDateFormat("dd/MM/yyyy").parse(request.getDateOfBirth());
            } catch (ParseException e) {
                e.printStackTrace();
            }
            response = AddUserResponse.newBuilder().setResponse("Added user with id " + userRepository.save(new User(UUID.randomUUID().toString(), request.getUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), dateOfBirth, request.getBiography(),request.getIsPrivate(), new ArrayList<>(), new ArrayList<>(), new ArrayList<>(), new ArrayList<>())).getId()).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void updateUser(UpdateUserRequest request, StreamObserver<UpdateUserResponse> responseObserver) {
        UpdateUserResponse response;
        if(isNullOrEmpty(request.getUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), request.getDateOfBirth().toString(), request.getBiography()))
            response = UpdateUserResponse.newBuilder().setError("None of fields cannot be empty!").build();
        else{
            Date dateOfBirth = new Date();
            try {
                dateOfBirth = new SimpleDateFormat("dd/MM/yyyy").parse(request.getDateOfBirth());
            } catch (ParseException e) {
                e.printStackTrace();
            }
            if(!isNullOrEmpty(request.getNewUsername())&&!request.getUsername().equals(request.getNewUsername())) {
                if(userRepository.findUserByUsername(request.getNewUsername()) != null)
                    response = UpdateUserResponse.newBuilder().setError("Username already exists!").build();
                else
                {
                    User user = userRepository.findUserByUsername(request.getUsername());
                    response = UpdateUserResponse.newBuilder().setResponse("Updated user with id " + userRepository.save(new User(user.getId(), request.getNewUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), dateOfBirth, request.getBiography(), request.getIsPrivate(), user.getEducation(), user.getWorkExperience(), user.getInterests(), user.getSkills())).getId()).build();
                }
            }
            else{
                User user = userRepository.findUserByUsername(request.getUsername());
                response = UpdateUserResponse.newBuilder().setResponse("Updated user with id " + userRepository.save(new User(user.getId(), request.getUsername(), request.getPassword(), request.getName(), request.getEmail(), request.getTelephoneNo(), request.getGender(), dateOfBirth, request.getBiography(), request.getIsPrivate(), user.getEducation(), user.getWorkExperience(), user.getInterests(), user.getSkills())).getId()).build();
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
            Date dateOfBirth = presentUser.getDateOfBirth();
            String stringDate = dateOfBirth.toString();
            try{
            SimpleDateFormat DateFor = new SimpleDateFormat("MM/dd/yyyy");
            stringDate= DateFor.format(dateOfBirth);
            }catch (Exception e){
                e.printStackTrace();
            }
            response = GetUserResponse.newBuilder().setUser(com.example.usermicroservice.User.newBuilder().
                    setUsername(presentUser.getUsername())
                    .setEmail(presentUser.getEmail())
                    .setBiography(presentUser.getBiography())
                    .setDateOfBirth(stringDate)
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

    @Override
    public void searchPublicUsers(SearchRequest request, StreamObserver<GetUsersResponse> responseObserver) {
        GetUsersResponse response;
        List<User> publicUsers = userService.searchPublicUsers(request.getSearchTerm());
        List<com.example.usermicroservice.User> users = new ArrayList<com.example.usermicroservice.User>();
        for(User user: publicUsers){
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
    public void updateInterests(UpdateInterestsRequest request, StreamObserver<StringResponse> responseObserver) {
        StringResponse stringResponse;
        if(userRepository.findById(request.getUserId()).isPresent()){
            Optional<User> optionalUser = userRepository.findById(request.getUserId());
            User user = optionalUser.get();
            user.setInterests(request.getInterestsList());
            userRepository.save(user);
            stringResponse = StringResponse.newBuilder().setResponse("Updated user with id " + request.getUserId()).build();
        }else{
            stringResponse = StringResponse.newBuilder().setResponse("No existing user with id " + request.getUserId() + "!").build();
        }
        responseObserver.onNext(stringResponse);
        responseObserver.onCompleted();
    }

    @Override
    public void updateSkills(UpdateSkillsRequest request, StreamObserver<StringResponse> responseObserver) {
        StringResponse stringResponse;
        if(userRepository.findById(request.getUserId()).isPresent()){
            Optional<User> optionalUser = userRepository.findById(request.getUserId());
            User user = optionalUser.get();
            user.setSkills(request.getSkillsList());
            userRepository.save(user);
            stringResponse = StringResponse.newBuilder().setResponse("Updated user with id " + request.getUserId()).build();
        }else{
            stringResponse = StringResponse.newBuilder().setResponse("No existing user with id " + request.getUserId() + "!").build();
        }
        responseObserver.onNext(stringResponse);
        responseObserver.onCompleted();
    }

    @Override
    public void updateWorkExperience(UpdateWorkExperienceRequest request, StreamObserver<StringResponse> responseObserver) {
        StringResponse stringResponse;
        if(userRepository.findById(request.getUserId()).isPresent()){
            Optional<User> optionalUser = userRepository.findById(request.getUserId());
            User user = optionalUser.get();
            List<WorkExperience> workExperienceList = new ArrayList<>();
            for(com.example.usermicroservice.WorkExperience workExperience:request.getWorkExperiencesList())
                workExperienceList.add(new WorkExperience(workExperience.getCompanyName(), workExperience.getJobTitle(),
                        getFormattedDate(workExperience.getStartDate()), getFormattedDate(workExperience.getEndDate())));
            user.setWorkExperience(workExperienceList);
            userRepository.save(user);
            stringResponse = StringResponse.newBuilder().setResponse("Updated user with id " + request.getUserId()).build();
        }else{
            stringResponse = StringResponse.newBuilder().setResponse("No existing user with id " + request.getUserId() + "!").build();
        }
        responseObserver.onNext(stringResponse);
        responseObserver.onCompleted();
    }

    @Override
    public void updateEducation(UpdateEducationRequest request, StreamObserver<StringResponse> responseObserver) {
        StringResponse stringResponse;
        if(userRepository.findById(request.getUserId()).isPresent()){
            Optional<User> optionalUser = userRepository.findById(request.getUserId());
            User user = optionalUser.get();
            List<com.example.usermicroservice.model.Education> educationList = new ArrayList<>();
            for(com.example.usermicroservice.Education education:request.getEducationList())
                educationList.add(new com.example.usermicroservice.model.Education(education.getInstitutionType(),
                        education.getInstitutionName(), education.getTitle(), education.getGpa(),
                        getFormattedDate(education.getStartDate()), getFormattedDate(education.getEndDate())));
            user.setEducation(educationList);
            userRepository.save(user);
            stringResponse = StringResponse.newBuilder().setResponse("Updated user with id " + request.getUserId()).build();
        }else{
            stringResponse = StringResponse.newBuilder().setResponse("No existing user with id " + request.getUserId() + "!").build();
        }
        responseObserver.onNext(stringResponse);
        responseObserver.onCompleted();
    }

    @Override
    public void getInterests(GetUserRequest request, StreamObserver<GetInterestsResponse> responseObserver) {
        GetInterestsResponse response;
        if(userRepository.findById(request.getId()).isEmpty()){
            response = GetInterestsResponse.newBuilder().addAllInterests(new ArrayList<>()).build();
        }else {
            User user = userRepository.findById(request.getId()).get();
            response = GetInterestsResponse.newBuilder().addAllInterests(user.getInterests()).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getSkills(GetUserRequest request, StreamObserver<GetSkillsResponse> responseObserver) {
        GetSkillsResponse response;
        if(userRepository.findById(request.getId()).isEmpty()){
            response = GetSkillsResponse.newBuilder().addAllSkills(new ArrayList<>()).build();
        }else {
            User user = userRepository.findById(request.getId()).get();
            response = GetSkillsResponse.newBuilder().addAllSkills(user.getSkills()).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getWorkExperience(GetUserRequest request, StreamObserver<GetWorkExperienceResponse> responseObserver) {
        GetWorkExperienceResponse response;
        if(userRepository.findById(request.getId()).isEmpty()){
            response = GetWorkExperienceResponse.newBuilder().addAllWorkExperience(new ArrayList<>()).build();
        }else {
            List<com.example.usermicroservice.WorkExperience> workExperienceList = new ArrayList<>();
            User user = userRepository.findById(request.getId()).get();
            for(WorkExperience workExperience:user.getWorkExperience())
                workExperienceList.add(com.example.usermicroservice.WorkExperience.newBuilder().
                        setCompanyName(workExperience.getCompanyName())
                        .setJobTitle(workExperience.getJobTitle())
                        .setStartDate(workExperience.getStartDate().toString())
                        .setEndDate(workExperience.getEndDate().toString())
                        .build());
            response = GetWorkExperienceResponse.newBuilder().addAllWorkExperience(workExperienceList).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getEducation(GetUserRequest request, StreamObserver<GetEducationResponse> responseObserver) {
        GetEducationResponse response;
        if(userRepository.findById(request.getId()).isEmpty()){
            response = GetEducationResponse.newBuilder().addAllEducation(new ArrayList<>()).build();
        }else {
            List<com.example.usermicroservice.Education> educationList = new ArrayList<>();
            User user = userRepository.findById(request.getId()).get();
            for(com.example.usermicroservice.model.Education education:user.getEducation())
                educationList.add(com.example.usermicroservice.Education.newBuilder().
                        setTitle(education.getTitle())
                        .setInstitutionName(education.getInstitutionName())
                        .setInstitutionType(education.getInstitutionType())
                        .setGpa(education.getGpa())
                        .setStartDate(education.getStartDate().toString())
                        .setEndDate(education.getEndDate().toString())
                        .build());
            response = GetEducationResponse.newBuilder().addAllEducation(educationList).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void login(LoginRequest request, StreamObserver<LoginResponse> responseObserver) {
        LoginResponse response;
        if(userRepository.findUserByUsername(request.getUsername()) == null)
            response = LoginResponse.newBuilder().setError("Invalid username/password!").build();
        else{
            User user = userRepository.findUserByUsername(request.getUsername());
            if(user.getPassword().equals(request.getPassword())){
                response = LoginResponse.newBuilder().setStatus(200).setToken(jwtTokenUtil.generateToken(user.getUsername())).setId(user.getId()).build();
            }else{
                response = LoginResponse.newBuilder().setStatus(400).setError("Invalid username/password!").build();
            }
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void validate(ValidateRequest request, StreamObserver<ValidateResponse> responseObserver) {
        ValidateResponse response;
        if(!jwtTokenUtil.checkIfClaimsAreExtractableFromToken(request.getToken()))
            response = ValidateResponse.newBuilder().setStatus(401).setError("Invalid token!").build();
        else if(request.getUsername().isEmpty()||request.getToken().isEmpty())
            response = ValidateResponse.newBuilder().setStatus(401).setError("Username/token cannot be empty!").build();
        else if(jwtTokenUtil.validateToken(request.getToken(), request.getUsername()))
            response = ValidateResponse.newBuilder().setStatus(200).build();
        else
            response = ValidateResponse.newBuilder().setStatus(401).setError("Token authentication failed or token expired!").build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getUserForEdit(GetUserRequest request, StreamObserver<GetUserResponse> responseObserver) {
        GetUserResponse response;
        Optional<User> user = userRepository.findById(request.getId());
        if(user.isPresent()){
            User presentUser = user.get();
            Date dateOfBirth = presentUser.getDateOfBirth();
            String stringDate = dateOfBirth.toString();
            try{
                SimpleDateFormat DateFor = new SimpleDateFormat("yyyy/MM/dd");
                stringDate= DateFor.format(dateOfBirth);
            }catch (Exception e){
                e.printStackTrace();
            }
            response = GetUserResponse.newBuilder().setUser(com.example.usermicroservice.User.newBuilder().
                    setUsername(presentUser.getUsername())
                    .setEmail(presentUser.getEmail())
                    .setBiography(presentUser.getBiography())
                    .setDateOfBirth(stringDate)
                    .setGender(presentUser.getGender())
                    .setId(presentUser.getId())
                    .setTelephoneNo(presentUser.getTelephoneNo())
                    .setName(presentUser.getName())
                    .setPassword(presentUser.getPassword())
            ).build();
        }else{
            response = GetUserResponse.newBuilder().setUser(com.example.usermicroservice.User.newBuilder()).build();
        }
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    private Date getFormattedDate(String date){
        Date formattedDate = new Date();
        try {
            formattedDate = new SimpleDateFormat("dd/MM/yyyy").parse(date);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return formattedDate;
    }

    private static boolean isNullOrEmpty(String... strArr){
        for (String st : strArr) {
            if  (st==null || st.equals(""))
                return true;

        }
        return false;
    }
}

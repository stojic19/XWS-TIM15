package com.example.usermicroservice;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class UserMicroserviceApplication {

    public static void main(String[] args) {
        System.out.println(System.getProperty("java.version"));
        SpringApplication.run(UserMicroserviceApplication.class, args);
    }

}

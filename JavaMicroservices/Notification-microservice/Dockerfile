FROM openjdk:18.0.1-jdk
MAINTAINER baeldung.com
COPY ./JavaMicroservices/Notification-microservice/target/Notification-microservice-0.0.1-SNAPSHOT.jar notifications.jar
ENTRYPOINT ["java","-jar","/notifications.jar"]
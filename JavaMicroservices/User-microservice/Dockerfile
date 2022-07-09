FROM openjdk:18.0.1-jdk
MAINTAINER baeldung.com
COPY ./JavaMicroservices/User-microservice/target/User-microservice-0.0.1-SNAPSHOT.jar users.jar
ENTRYPOINT ["java","-jar","/users.jar"]
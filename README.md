# ASG_2
<div align="center">
  <h3 align="center">ASG 2 Package 3.10</h3>
</div>

![ETI ASG 2](https://user-images.githubusercontent.com/83932770/152684287-88338b88-a23d-41f4-951c-d9201d111e4e.png)



There are a total of 2 microservices required for this package:
* Comment Microservice
* Rating Microservice

There are a total of 2 databases directly involved in this package:
* ETIAssignment2Comment
* ETIAssignment2Rating

Design Considerations:
I designed the architecture in a way that respects the Microservice value of being independent and loosely coupled. For the requirements for my package, two main domains stand out Comments and Ratings with Supporting Domains being the Students and Tutors that create and update them. 

As a result, the two main MS would be the Comment and Rating MS with and they both have their respective databases.To Support this feature, the an external User MS to facilitate the Users (Student and Tutor) and link Users to their Comments and Ratings. 

In my UI, it would be in the perspective of the Tutor's end, where he/she are entitled to the following features:
   * Viewing Received Comments / Ratings From Students
   * Viewing Anonymously Received Comments / Ratings From Students
   * Viewing Given Comments / Ratings
   * Create Comments and Ratings for Students
   * Update Created Comments and Ratings for Students

DockerHub Images Link:



Set Up Guide:

1. Clone the Repo
2. Install necessary Libraries:
   * go get -u github.com/go-sql-driver/mysql
   * go get -u github.com/gorilla/mux
   * go get -u github.com/gorilla/handlers
3. In MYSQL Workbench create these Connection Strings:
   * Connection Name ratingdb,  HostName 127.0.0.1, Port: 9121, User: Root
   * Connection Name commentdb,  HostName 127.0.0.1, Port: 9122, User: Root
   * Connection Name testuserdb,  HostName 127.0.0.1, Port: 9123, User: Root
4. On VSCode, Right Click the docker-compose.yml file and click "Compose Up"
5. In your Docker Hub Containers, click "Open In Browser" for the container "frontend"
6. Enter http://localhost:9128/ViewComments.html to enter the UI




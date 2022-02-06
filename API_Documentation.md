# ETIAssignment2
## API Endpoints:
| Microservice  | Port | Endpoint URL |
| ------------- | ---- | ------------ |
| Frontend  | 9128 | http://10.31.11.12:9128 |
| Comment API  | 9125 | http://10.31.11.12:9125 |
| Rating API  | 9124 | http://10.31.11.12:9124 |

## Front End Webpage Routes
```sh
# 1. View Comments Received
http://10.31.11.12:9128/ViewComments.html
# 2. View Anon Comments Received
http://10.31.11.12:9128/ViewAnonComments.html
# 3. View Ratings Received
http://10.31.11.12:9128/ViewRatings.html
# 4. View Anon Ratings Received
http://10.31.11.12:9128/ViewAnonRatings.html
# 5. View Given Comments 
http://10.31.11.12:9128/ViewGivenComments.html
# 6. View Given Ratings 
http://10.31.11.12:9128/ViewGivenRatings.html
# 7. Create Comments 
http://10.31.11.12:9128/CreateComment.html
# 8. Create Ratings 
http://10.31.11.12:9128/CreateRating.html
# 9. Update Comments 
http://10.31.11.12:9128/UpdateComment.html
# 10. Update Ratings 
http://10.31.11.12:9128/UpdateRating.html



```


### Comment Data Structure
| Attribute | Data Type |
| --------- | ---- |
| CommentID | uint |
| CreatorType | varchar(255) |
| CreatorID | varchar(9) |
| TargetType | varchar(255) |
| TargetID | varchar(9) |
| CommentData | varchar(500) |
| Anonymous | bool |
| DateTimePublished | datetime |

### Rating Data Structure
| Attribute | Data Type |
| --------- | ---- |
| RatingID | uint |
| CreatorType | varchar(255) |
| CreatorID | varchar(9) |
| TargetType | varchar(255) |
| TargetID | varchar(9) |
| RatingScore | varchar(500) |
| Anonymous | bool |
| DateTimePublished | datetime |

# Comment Microservice API Documentation
### [GET] /api/comment/tutor
Get All Comments created by all Tutor
```
Endpoint
http://10.31.11.12:9125/api/comment/tutor
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/comment/student
Get All Comments created by all Student
```
Endpoint
http://10.31.11.12:9125/api/comment/student
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/comment/tutor/unanonymized/{tutorid}
Get All received Unanonymized Comments for Tutor
```
Endpoint
http://10.31.11.12:9125/api/comment/tutor/unanonymized/{tutorid}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/comment/tutor/anonymized/{tutorid}
Get All received anonymized Comments for Tutor
```
Endpoint
http://10.31.11.12:9125/api/comment/tutor/anonymized/{tutorid}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/comment/tutor/given/{tutorid}
Get All given Comments for Tutor
```
Endpoint
http://10.31.11.12:9125/api/comment/tutor/given/{tutorid}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [POST] /api/comment/
Creates a Comment
```
Endpoint
http://10.31.11.12:9125/api/comment
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 

Example of Curl Request: 
curl -H "Content-Type:application/json" -X POST http://localhost:5001/api/comment -d "{\"CreatorID\":2,\"CreatorType\":\"Tutor\",\"TargetID\":1,\"TargetType\":\"Student\",\"CommentData\":\"Very Hardworking\",\"Anonymous\":0}"

```

### [PUT] /api/comment/
Updates a Comment
```
Endpoint
http://10.31.11.12:9125/api/comment
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 

Example of Curl Request: 
curl -H "Content-Type:application/json" -X PUT http://localhost:5001/api/comment -d "{\"CommentID\":16,\"TargetType\":\"Student\",\"CommentData\":\"Mediocre at best.\"}"

```

# Rating Microservice API Documentation
### [GET] /api/rating/student
Get All Ratings created by all Student
```
Endpoint
http://10.31.11.12:9124/api/rating/student
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/rating/tutor/unanonymized/{tutorid}
Get All Unanonymized Ratings received by Tutor
```
Endpoint
http://10.31.11.12:9124/api/rating/tutor/unanonymized/{tutorid}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/rating/tutor/anonymized/{tutorid}
Get All anonymized Ratings received by Tutor
```
Endpoint
http://10.31.11.12:9124/api/rating/tutor/anonymized/{tutorid}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/rating/tutor/given/{tutorid}
Get All anonymized Ratings given by Tutor
```
Endpoint
http://10.31.11.12:9124/api/rating/tutor/given/{tutorid}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [POST] /api/rating
Create a Rating
```
Endpoint
http://10.31.11.12:9124/api/rating
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 

Example of Curl Request:
curl -H "Content-Type:application/json" -X POST http://localhost:5002/api/rating -d "{\"CreatorID\":2,\"CreatorType\":\"Tutor\",\"TargetID\":1,\"TargetType\":\"Student\",\"RatingScore\":5,\"Anonymous\":0}"

```

### [PUT] /api/rating
Create a Rating
```
Endpoint
http://10.31.11.12:9124/api/rating
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 

Example of Curl Request:
curl -H "Content-Type:application/json" -X PUT http://localhost:5002/api/rating -d "{\"RatingID\":1,\"TargetType\":\"Student\",\"RatingScore\":100}"

```





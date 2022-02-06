package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Rating struct {
	RatingID          int
	CreatorID         int
	CreatorType       string
	TargetID          int
	TargetType        string
	RatingScore       int
	Anonymous         int
	DateTimePublished string
	CreatorName       string
	TargetName        string
}

type Object struct {
	ID   int
	Name string
}

func getAllStudents(db *sql.DB) []Object {
	url := "http://localhost:9125/api/student"
	response, err := http.Get(url)
	var studentList []Object
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode == http.StatusNotFound {
			fmt.Println("409 - No Students Found!")
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			json.Unmarshal(data, &studentList)
			fmt.Println("202 - Successfully received Students!")
		}
	}
	return studentList
}

func getAllTutors(db *sql.DB) []Object {
	url := "http://localhost:9126/api/tutor"
	response, err := http.Get(url)
	var tutorList []Object
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode == http.StatusNotFound {
			fmt.Println("409 - No Tutors Found!")
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			json.Unmarshal(data, &tutorList)
			fmt.Println("202 - Successfully received Tutors!")
		}
	}
	return tutorList
}

func linkStudentToID(db *sql.DB, id int, studentList []Object) Object {
	var student Object
	for _, student := range studentList {
		if student.ID == id {
			return student
		}
	}
	return student
}

func linkTutorToID(db *sql.DB, id int, tutorList []Object) Object {
	var tutor Object
	for _, tutor := range tutorList {
		if tutor.ID == id {
			return tutor
		}
	}
	return tutor
}

// 3.10

func rating(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(ratingdb:3306)/ratingdb")
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "POST" {
		var newRating Rating
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newRating)
			fmt.Println(newRating.RatingScore)
			postRating(db, newRating)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("201 - Rating Posted!"))
		}
	} else if r.Method == "PUT" {
		var newRating Rating
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newRating)
			fmt.Println(newRating.RatingScore)
			updateRating(db, newRating)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("201 - Rating Updated!"))
		}

	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("422 - Rating Info not in JSON format!"))
	}
}

func postRating(db *sql.DB, rating Rating) {

	CreatorID := rating.CreatorID
	CreatorType := rating.CreatorType
	TargetID := rating.TargetID
	TargetType := rating.TargetType
	RatingScore := rating.RatingScore
	Anonymous := rating.Anonymous

	query := fmt.Sprintf("INSERT INTO Rating (CreatorID, CreatorType, TargetID, TargetType, RatingScore, Anonymous, DateTimePublished) VALUES ('%d', '%s', '%d', '%s', '%d', '%b', NOW())",
		CreatorID, CreatorType, TargetID, TargetType, RatingScore, Anonymous)
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func updateRating(db *sql.DB, rating Rating) {

	RatingID := rating.RatingID
	RatingScore := rating.RatingScore
	RatingTargetType := rating.TargetType
	query := ""
	query = fmt.Sprintf("UPDATE Rating SET RatingScore = '%d' WHERE RatingID = '%d' AND TargetType = '%s'", RatingScore, RatingID, RatingTargetType)
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func allStudentRatings(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(ratingdb:3306)/ratingdb")

	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			var allStudentRatings []Rating = getAllStudentRatings(db)
			if len(allStudentRatings) > 0 {
				fmt.Println(allStudentRatings)
				json.NewEncoder(w).Encode(allStudentRatings)
				fmt.Println(allStudentRatings)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}

func getAllStudentRatings(db *sql.DB) []Rating {

	studentList := getAllStudents(db)

	query := fmt.Sprintf("SELECT * FROM Rating WHERE CreatorType = 'Student'")

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var ratingList []Rating
	println(results)
	for results.Next() {
		var rating Rating
		err = results.Scan(&rating.RatingID, &rating.CreatorID, &rating.CreatorType, &rating.TargetID, &rating.TargetType, &rating.RatingScore, &rating.Anonymous, &rating.DateTimePublished)
		if err != nil {
			panic(err.Error())
		}

		student := linkStudentToID(db, rating.CreatorID, studentList)
		rating.CreatorName = student.Name

		println(student.Name)
		ratingList = append(ratingList, rating)
	}
	return ratingList
}

func tutorReceivedRatings(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(ratingdb:3306)/ratingdb")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	tutorID := params["tutorid"]
	tutorIDint, err := strconv.Atoi(tutorID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		tutorRatingList := getTutorRatings(db, tutorIDint)
		if len(tutorRatingList) > 0 {
			fmt.Println(tutorRatingList)
			json.NewEncoder(w).Encode(tutorRatingList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getTutorRatings(db *sql.DB, targetID int) []Rating {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	fmt.Println(studentList)
	fmt.Println(tutorList)
	tutorquery := fmt.Sprintf("SELECT * FROM Rating WHERE TargetType = 'Tutor' AND TargetID = '%d' AND Anonymous = 0;", targetID)

	tutorResults, err := db.Query(tutorquery)
	if err != nil {
		panic(err.Error())
	}
	var tutorRatingList []Rating
	for tutorResults.Next() {
		var rating Rating
		err = tutorResults.Scan(&rating.RatingID, &rating.CreatorID, &rating.CreatorType, &rating.TargetID, &rating.TargetType, &rating.RatingScore, &rating.Anonymous, &rating.DateTimePublished)

		student := linkStudentToID(db, rating.CreatorID, studentList)
		rating.CreatorName = student.Name
		tutor := linkTutorToID(db, rating.TargetID, tutorList)
		rating.TargetName = tutor.Name
		fmt.Println(rating)
		tutorRatingList = append(tutorRatingList, rating)
	}

	return tutorRatingList
}

func AnonymoustutorReceivedRatings(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(ratingdb:3306)/ratingdb")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	tutorID := params["tutorid"]
	tutorIDint, err := strconv.Atoi(tutorID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		tutorRatingList := getAnonTutorRatings(db, tutorIDint)
		if len(tutorRatingList) > 0 {
			fmt.Println(tutorRatingList)
			json.NewEncoder(w).Encode(tutorRatingList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getAnonTutorRatings(db *sql.DB, targetID int) []Rating {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	fmt.Println(studentList)
	fmt.Println(tutorList)
	tutorquery := fmt.Sprintf("SELECT * FROM Rating WHERE TargetType = 'Tutor' AND TargetID = '%d' AND Anonymous = 1;", targetID)

	tutorResults, err := db.Query(tutorquery)
	if err != nil {
		panic(err.Error())
	}
	var tutorRatingList []Rating
	for tutorResults.Next() {
		var rating Rating
		err = tutorResults.Scan(&rating.RatingID, &rating.CreatorID, &rating.CreatorType, &rating.TargetID, &rating.TargetType, &rating.RatingScore, &rating.Anonymous, &rating.DateTimePublished)

		tutor := linkTutorToID(db, rating.TargetID, tutorList)
		rating.TargetName = tutor.Name

		fmt.Println(rating)
		tutorRatingList = append(tutorRatingList, rating)
	}

	return tutorRatingList
}

func tutorGivenRatings(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(ratingdb:3306)/ratingdb")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	tutorID := params["tutorid"]
	tutorIDint, err := strconv.Atoi(tutorID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		tutorGivenRatingList := getTutorGivenRatings(db, tutorIDint)
		if len(tutorGivenRatingList) > 0 {
			fmt.Println(tutorGivenRatingList)
			json.NewEncoder(w).Encode(tutorGivenRatingList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getTutorGivenRatings(db *sql.DB, targetID int) []Rating {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	fmt.Println(studentList)
	fmt.Println(tutorList)
	tutorquery := fmt.Sprintf("SELECT * FROM Rating WHERE CreatorType = 'Tutor' AND CreatorID = '%d';", targetID)

	tutorResults, err := db.Query(tutorquery)
	if err != nil {
		panic(err.Error())
	}
	var tutorRatingList []Rating
	for tutorResults.Next() {
		var rating Rating
		err = tutorResults.Scan(&rating.RatingID, &rating.CreatorID, &rating.CreatorType, &rating.TargetID, &rating.TargetType, &rating.RatingScore, &rating.Anonymous, &rating.DateTimePublished)

		tutor := linkTutorToID(db, rating.TargetID, tutorList)
		rating.TargetName = tutor.Name

		fmt.Println(rating)
		tutorRatingList = append(tutorRatingList, rating)
	}

	return tutorRatingList
}

func main() {
	// This is to allow the headers, origins and methods all to access CORS resource sharing
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router := mux.NewRouter()

	// Create &  Update Ratings

	router.HandleFunc("/api/rating", rating).Methods("POST", "PUT")

	// View all Student Ratings

	router.HandleFunc("/api/rating/student", allStudentRatings).Methods("GET")

	// View all Unanonymized Recevied Ratings

	router.HandleFunc("/api/rating/tutor/unanonymized/{tutorid}", tutorReceivedRatings).Methods("GET")

	// View all Anonymized received ratings

	router.HandleFunc("/api/rating/tutor/anonymized/{tutorid}", AnonymoustutorReceivedRatings).Methods("GET")

	// View Given Ratings

	router.HandleFunc("/api/rating/tutor/given/{tutorid}", tutorGivenRatings).Methods("GET")

	fmt.Println("Listening at port 9124")
	log.Fatal(http.ListenAndServe(":9124", handlers.CORS(headers, origins, methods)(router)))
}

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
	// "github.com/julienschmidt/httprouter"
)

type Comment struct {
	CommentID         int
	CreatorID         int
	CreatorType       string
	TargetID          int
	TargetType        string
	CommentData       string
	Anonymous         int
	DateTimePublished string
	CreatorName       string
	TargetName        string
}

type Object struct {
	ID   int
	Name string
}

//Gets all Student's ID which is tied to StudentID
func getAllStudents(db *sql.DB) []Object {
	url := "http://localhost:5003/api/student"
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

//Gets all Tutor's Names which is tied to TutorID
func getAllTutors(db *sql.DB) []Object {
	url := "http://localhost:5004/api/tutor"
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

//Tutor View Comments Received
func tutorReceivedComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")
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
		tutorCommentList := getTutorComments(db, tutorIDint)
		if len(tutorCommentList) > 0 {
			fmt.Println(tutorCommentList)
			json.NewEncoder(w).Encode(tutorCommentList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

//Tutor View Comments Received UN Anoynymous
func getTutorComments(db *sql.DB, targetID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	fmt.Println(studentList)
	fmt.Println(tutorList)
	tutorquery := fmt.Sprintf("SELECT * FROM Comment WHERE TargetType = 'Tutor' AND TargetID = '%d' AND Anonymous = 0;", targetID)

	tutorResults, err := db.Query(tutorquery)
	if err != nil {
		panic(err.Error())
	}
	var tutorCommentList []Comment
	for tutorResults.Next() {
		var comment Comment
		tutorResults.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)

		student := linkStudentToID(db, comment.CreatorID, studentList)
		comment.CreatorName = student.Name
		tutor := linkTutorToID(db, comment.TargetID, tutorList)
		comment.TargetName = tutor.Name
		fmt.Println(comment)
		tutorCommentList = append(tutorCommentList, comment)
	}

	return tutorCommentList
}

//Tutor View Anonymous Comments Received
func AnonymoustutorReceivedComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")
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
		tutorCommentList := getAnonymousTutorComments(db, tutorIDint)
		if len(tutorCommentList) > 0 {
			fmt.Println(tutorCommentList)
			json.NewEncoder(w).Encode(tutorCommentList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getAnonymousTutorComments(db *sql.DB, targetID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	fmt.Println(studentList)
	fmt.Println(tutorList)
	tutorquery := fmt.Sprintf("SELECT * FROM Comment WHERE TargetType = 'Tutor' AND TargetID = '%d' AND Anonymous = 1;", targetID)

	tutorResults, err := db.Query(tutorquery)
	if err != nil {
		panic(err.Error())
	}
	var tutorCommentList []Comment
	for tutorResults.Next() {
		var comment Comment
		tutorResults.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)

		student := linkStudentToID(db, comment.CreatorID, studentList)
		comment.CreatorName = student.Name
		tutor := linkTutorToID(db, comment.TargetID, tutorList)
		comment.TargetName = tutor.Name
		fmt.Println(comment)
		tutorCommentList = append(tutorCommentList, comment)
	}

	return tutorCommentList
}

// Given Comments

func tutorGivenComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")
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
		tutorGivenCommentList := getTutorGivenComments(db, tutorIDint)
		if len(tutorGivenCommentList) > 0 {
			fmt.Println(tutorGivenCommentList)
			json.NewEncoder(w).Encode(tutorGivenCommentList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getTutorGivenComments(db *sql.DB, targetID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	fmt.Println(studentList)
	fmt.Println(tutorList)
	tutorquery := fmt.Sprintf("SELECT * FROM Comment WHERE CreatorType = 'Tutor' AND CreatorID = '%d';", targetID)

	tutorResults, err := db.Query(tutorquery)
	if err != nil {
		panic(err.Error())
	}
	var tutorCommentList []Comment
	for tutorResults.Next() {
		var comment Comment
		tutorResults.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)

		student := linkStudentToID(db, comment.CreatorID, studentList)
		comment.TargetName = student.Name
		tutor := linkTutorToID(db, comment.TargetID, tutorList)
		comment.CreatorName = tutor.Name
		fmt.Println(comment)
		tutorCommentList = append(tutorCommentList, comment)
	}

	return tutorCommentList
}

func allTutorComments(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")

	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			var allTutorComments []Comment = getAllTutorComments(db)
			if len(allTutorComments) > 0 {
				fmt.Println(allTutorComments)
				json.NewEncoder(w).Encode(allTutorComments)
				fmt.Println(allTutorComments)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}

func getAllTutorComments(db *sql.DB) []Comment {

	tutorList := getAllTutors(db)

	query := fmt.Sprintf("SELECT * FROM ETIAssignment2Comment.Comment WHERE CreatorType = 'Tutor'")

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var commentList []Comment
	println(results)
	for results.Next() {
		var comment Comment
		err = results.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if err != nil {
			panic(err.Error())
		}

		tutor := linkTutorToID(db, comment.CreatorID, tutorList)
		comment.CreatorName = tutor.Name

		println(tutor.Name)
		commentList = append(commentList, comment)
	}
	return commentList
}

func allStudentComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")

	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			var allStudentComments []Comment = getAllStudentComments(db)
			if len(allStudentComments) > 0 {
				fmt.Println(allStudentComments)
				json.NewEncoder(w).Encode(allStudentComments)
				fmt.Println(allStudentComments)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}

func getAllStudentComments(db *sql.DB) []Comment {

	studentList := getAllStudents(db)

	query := fmt.Sprintf("SELECT * FROM ETIAssignment2Comment.Comment WHERE CreatorType = 'Student'")

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var commentList []Comment
	println(results)
	for results.Next() {
		var comment Comment
		err = results.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if err != nil {
			panic(err.Error())
		}

		student := linkStudentToID(db, comment.CreatorID, studentList)
		comment.CreatorName = student.Name

		println(student.Name)
		commentList = append(commentList, comment)
	}
	return commentList
}

func postComment(db *sql.DB, comment Comment) {
	CreatorType := comment.CreatorType
	println(comment.CreatorType)
	CreatorID := comment.CreatorID
	println(comment.CreatorID)
	TargetID := comment.TargetID
	println(comment.TargetID)
	CommentData := comment.CommentData
	println(comment.CommentData)
	Anonymous := comment.Anonymous
	println(comment.Anonymous)
	TargetType := comment.TargetType
	println(comment.TargetType)
	query := fmt.Sprintf("INSERT INTO Comment (CreatorType, CreatorID, TargetID, TargetType, CommentData, Anonymous, DateTimePublished) VALUES ('%s', '%d', '%d', '%s', '%s', '%b', NOW())",
		CreatorType, CreatorID, TargetID, TargetType, CommentData, Anonymous)
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func updateComment(db *sql.DB, comment Comment) {
	CommentID := comment.CommentID
	CommentData := comment.CommentData
	CommentTargetType := comment.TargetType
	query := ""
	query = fmt.Sprintf("UPDATE Comment SET CommentData = '%s' WHERE CommentID = '%d' AND TargetType = '%s'", CommentData, CommentID, CommentTargetType)

	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func comment(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user2:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "POST" {
		var newComment Comment
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newComment)
			fmt.Println(newComment.CommentData)
			postComment(db, newComment)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("201 - Comment Posted!"))
		}
	} else if r.Method == "PUT" {
		var newComment Comment
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newComment)
			fmt.Println(newComment.CommentData)
			updateComment(db, newComment)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("201 - Comment Updated!"))
		}

	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("422 - Comment Info not in JSON format!"))
	}
}

func main() {
	// This is to allow the headers, origins and methods all to access CORS resource sharing
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router := mux.NewRouter()

	//3.10

	// Create and Update Comments

	router.HandleFunc("/api/comment", comment).Methods("POST", "PUT")

	// Get all Tutor Comments

	router.HandleFunc("/api/comment/tutor", allTutorComments).Methods("GET")

	// View all Student Comments

	router.HandleFunc("/api/comment/student", allStudentComments).Methods("GET")

	// View all unanonymized comments to Tutor

	router.HandleFunc("/api/comment/tutor/unanonymized/{tutorid}", tutorReceivedComments).Methods("GET")

	// View all anonymized comments to Tutor

	router.HandleFunc("/api/comment/tutor/anonymized/{tutorid}", AnonymoustutorReceivedComments).Methods("GET")

	// View all given comments

	router.HandleFunc("/api/comment/tutor/given/{tutorid}", tutorGivenComments).Methods("GET")

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", handlers.CORS(headers, origins, methods)(router)))
}

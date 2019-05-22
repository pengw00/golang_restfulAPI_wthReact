package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
)

//1. data struct model
type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
var Articles []Article

// sort of controller
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func getAllActicles(w http.ResponseWriter, r *http.Request){
	fmt.Println("EndPoint Hit: getAllArticles")
	// encode all the article into jSON file
	json.NewEncoder(w).Encode(Articles)
}

func getOneArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	// fmt.Fprintf(w, "Key: " + key) for debugging
	//this will got the key from the request, so we can scan the repo

	//for search, here I can also use binary search, which will be pretty quick!
	//when the dataset is pretty quick!

	//this code block is better to wrappd into a func
	var idExist = false;
	for _, article := range Articles {
			if article.Id == key {
				idExist = true
				json.NewEncoder(w).Encode(article)
			}
	}
	if !idExist {
		fmt.Fprintf(w, "Key: " + key + "not existed!")
	}
	}

	//create operate
	func createNewArticle(w http.ResponseWriter, r *http.Request){
		reqBody, _ := ioutil.ReadAll(r.Body)
		//for testing
		// fmt.Fprintf(w, "%+v", string(reqBody))
		var article Article
		json.Unmarshal(reqBody, &article)
		//update the repo array
		Articles = append(Articles, article)
		json.NewEncoder(w).Encode(article)
		fmt.Println("EndPoint Hit: postNewArticles")
	}

	func deleteArticle(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		id := vars["id"]
		// fmt.Fprintf(w, "Key: " + id)
		for index, article := range Articles {
			if article.Id == id {
				fmt.Println("EndPoint Hit: deleteNewArticles")
				Articles = append(Articles[:index], Articles[index+1:]...)
			}
		}
	}

//router and mapping endpoint
// func handleRequests() {
	
// create new instance of mux

func main(){
	//setup port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
		}
	fmt.Println(port)
	fmt.Println("Rest API v2.0 - Mux Routers")
	fmt.Println("Restful API running in host:8000")
	Articles = []Article{
		Article{Id: "1", Title: "Hello 1", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "3", Title: "Hello 3", Desc: "Article Description", Content: "Article Content"},
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	//Try sending a new HTTP DELETE request to http://localhost:10000/article/2. 
	//This will delete the second article within your Articles array and 
	//when you subsequently hit http://localhost:10000/articles with a HTTP GET request, 
	//you should see it now only contains a single Article.
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", getAllActicles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	//delete must put before getone, otherwise it will not trigger delete!!!!!
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", getOneArticle)
	
	log.Fatal(http.ListenAndServe(":" + port, myRouter))
}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
   ID string  `json:"id"`
   ISBN string 	`json:"isbn"`
   Title string `json:"title"`
   Director *Director `json:"director"`
}

type Director struct{
     Firstname string `json:"firstname"`
	 Secondname string `json:"secondname"`

}

var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
  
	   w.Header().Set("Content-Type","application/json")
	   json.NewEncoder(w).Encode(movies)
	   
}

func deleteMovie(w http.ResponseWriter,r *http.Request){
   
	  w.Header().Set("Content-Type","application/json")
	  params := mux.Vars(r)

	  for index,item := range movies{

		   if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			fmt.Fprintf(w,"Deleted the Movie")
			
		   } 
	  }
	  fmt.Fprintf(w,"Current Movies in the directory are\n")

	  json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)


	for _,item:= range movies{
		if item.ID == params["id"]{
			 json.NewEncoder(w).Encode(item)
			 return 
		} 
	}
}

func createMovie(w http.ResponseWriter,r *http.Request){
    
	w.Header().Set("Content-Type","application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa((rand.Intn(10000)))
	movies = append(movies,movie)

	fmt.Fprintf(w,"Current Movie directory looks like\n")
	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter,r *http.Request){
	 
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
		
			updatedMovie.ID = params["id"]
			movies = append(movies,updatedMovie)
		
			fmt.Fprintf(w,"Updated the Movie")
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func main(){
	fmt.Println("Entered the main function")

   // initialising inital set of movies

   movies = append(movies,Movie{ID:"1",ISBN:"CE180",Title:"LA LA LAND",Director: &Director{Firstname: "Rachel",Secondname:"Doe"}})
   movies = append(movies,Movie{ID:"2",ISBN:"SD191",Title:"BOSS BABY",Director: &Director{Firstname: "Martin",Secondname:"Garrix"}})
   movies = append(movies,Movie{ID:"3",ISBN:"MV008",Title:"DRIVE TO SURVIVE-2021",Director: &Director{Firstname: "Max",Secondname:"Verstappen"}})
   
   

	r:= mux.NewRouter()

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")

	err := http.ListenAndServe(":8080",r); if err!=nil{
		log.Fatal(err)
	}

}







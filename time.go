package main

import("fmt"
		"net/http"
		"log"
		"time"
	)

func handle_utc (w http.ResponseWriter , r *http.Request) {

	t:=time.Now()
	//const layout=  "2006-01-02T15:04:05Z07:00"
	
	fmt.Fprintf(w,"<h1 > CURRENT TIME</h1> ")
	fmt.Fprintf(w,"<p >Local Time: %s </p> ",t )
	
	loc,_ := time.LoadLocation("UTC")
	
   				 	now := time.Now().In(loc)
    				fmt.Fprintf(w ," <p>Time : %s </p>",now) // UTC 
				
    
}




func main(){

	http.HandleFunc("/time",handle_utc) 
	log.Fatal(http.ListenAndServe(":8000",nil))
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

//TODO: add the style
func main (){
    envErr := godotenv.Load(".env")
    if envErr != nil {
        log.Fatalf("Error loading env file %v",envErr)
    }

    tmpl, parsErr := template.ParseFS(os.DirFS("./templates"),"task.html")
    if(parsErr != nil){
        log.Fatalf("we couldn't read the content of the file %v",envErr)
    }
    http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request){
        tmpl.Execute(w,nil)
    })

    var port string = os.Getenv("PORT")
    fmt.Printf("Server is running on port %s",port)
    http.ListenAndServe(":" + port,nil)
    
}

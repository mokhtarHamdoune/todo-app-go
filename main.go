package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

//TODO: handle post data
//TODO: store the data on the array list

type Task struct  {
    Content string;
    IsDone bool;
}

func main (){
    envErr := godotenv.Load(".env")
    if envErr != nil {
        log.Fatalf("Error loading env file %v",envErr)
    }


    tmpl, parsErr := template.ParseFS(os.DirFS("./templates"),"task.html")
    if(parsErr != nil){
        log.Fatalf("we couldn't read the content of the file %v",envErr)
    }
    /* 
        * Serve static file 
        * StripPrefix as the name suggeste it will remove the /static/ prefix from the url
        * FilServer will serve the files of assets dir as  css/ js/ ...
        * If I get url /static/css/style.css => css/style.css and that will mache the Filserver content e.g. css/
    */ 
    http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./assets"))))
    data  :=  []Task{
        {
            Content: "The first task",
            IsDone: true,
        },
        {
            Content: "The second task",
            IsDone: false,
        },{
            Content: "The second task",
            IsDone: true,
        },
    }
    fmt.Println(data) 
    http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request){
        tmpl.Execute(w,data)
    })

    var port string = os.Getenv("PORT")
    fmt.Printf("Server is running on port %s",port)
    http.ListenAndServe(":" + port,nil)
    
}

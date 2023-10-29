package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/joho/godotenv"
)

type Task struct  {
    Content string;
    IsDone bool;
}

func deleteTask(tasks []Task,indexOfItem int) []Task {
    newTasks := make([]Task,0)
    for i:= 0; i < len(tasks); i++ {
        if i!= int(indexOfItem) {
            newTasks = append(newTasks, tasks[i])
        }
    }
    return newTasks
}
func main (){
    envErr := godotenv.Load(".env")
    if envErr != nil {
        log.Fatalf("Error loading env file %v",envErr)
    }
    tasks  := make([]Task,0);
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
        
    http.HandleFunc("/",func (w http.ResponseWriter, r *http.Request){
        // check if we have a POST request 
        if r.Method == http.MethodPost {
            // read the request body and parses it as a form
            // and puts the results into both r.PostForm and r.Form
            r.ParseForm()
            // we could use r.FormValues but I choose to use this cause it let me test 
            // the presence of the value
            if r.Form.Has("content") {
                newTaks := Task {
                    Content:r.Form.Get("content"),
                    IsDone:false,
                }
                tasks = append(tasks,newTaks)
            }
            // TODO: refactore the code and find a simple method to delete if exist
            if(r.Form.Has("done")){
                taksIndexStr:= r.FormValue("done")
                taksIndexInt,err := strconv.ParseInt(taksIndexStr,10,0);
                if  err != nil {
                    fmt.Printf("Error: %v\n", err)
                } 
                tasks = deleteTask(tasks,int(taksIndexInt))
            }
        }
        tmpl.Execute(w,tasks)
    })
    var port string = os.Getenv("PORT")
    fmt.Printf("Server is running on port %s\n",port)
    http.ListenAndServe(":" + port,nil)
    
}

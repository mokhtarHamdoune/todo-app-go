package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func  deleteTask(tasks []Task,taskId int) []Task {
    newTasks := make([]Task,0)
    for i:= 0; i < len(tasks); i++ {
        if tasks[i].Id!= taskId {
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
    // connect to the database 
    db,dbOpenerror :=  connectDB()
    if(dbOpenerror != nil){
        log.Fatal(dbOpenerror)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    taskManager := TaskManager{database: db}
    // get all task for the firt time 
    tasks, getAllErr := taskManager.getAll()

    if getAllErr != nil {
        log.Fatal(getAllErr)
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
                task, err := taskManager.save(newTaks)
                if err != nil {
                    log.Fatal(err)
                }
                tasks = append(tasks, task)
               
            }
            if(r.Form.Has("done")){
                taskIdStr:= r.FormValue("done")
                taskId,err := strconv.ParseInt(taskIdStr,10,0);
                if  err != nil {
                    log.Fatal(getAllErr)
                } 
                // we delete the task from the database
                error := taskManager.delete(int(taskId))
                if(error != nil){
                    log.Fatal(getAllErr)
                }
                // then list
                tasks = deleteTask(tasks,int(taskId))

            }
        }
        tmpl.Execute(w,tasks)
    })
    var port string = os.Getenv("PORT")
    fmt.Printf("Server is running on port %s\n",port)
    http.ListenAndServe(":" + port,nil)
    
}

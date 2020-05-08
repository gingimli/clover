package main

import (
    "encoding/json"
    "fmt"
    "github.com/gingimli/clover/db"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
    addHeaders(w)
    err := r.ParseForm()
    if err != nil {
        fmt.Println(err)
    }
    task := r.Form.Get("task")
    _, err = db.CreateTask(task)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Added '%s' to your task list.\n", task)
}

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
    addHeaders(w)
    tasks, err := db.AllTasks()
    if err != nil {
        fmt.Println(err)
    }
    err = json.NewEncoder(w).Encode(tasks)
}

func addHeaders(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
    port := "3000"
    dbPath := filepath.Join("./", "tasks.db")
    err := db.Init(dbPath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    http.HandleFunc("/", listTasksHandler)
    http.HandleFunc("/add", addTaskHandler)
    fmt.Printf("API running on: http://localhost:%s\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

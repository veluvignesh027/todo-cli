package main

import (
        "fmt"
        "os"
        "strconv"
        "time"

        "github.com/aquasecurity/table"
        "github.com/liamg/tml"
)

type task struct{
    Title string
    Status bool
    CreatedAt time.Time
    ModifiedAt time.Time
    CompletedAt time.Time
    Notify bool
    TimeNotify time.Time
}

var Tasks []task

func TaskAdd(msg string){
    obj := task{
        Title: msg,
        Status: false,
        CreatedAt: time.Now(),
        ModifiedAt: time.Now(),
        Notify: false,
    }

    Tasks = append(Tasks, obj)
}

func TaskDel(taskid int){
    fmt.Println("task deleted")
}

func TaskToggle(taskid int){
    if Tasks[taskid].Status {
        Tasks[taskid].Status = false
    }else{
        Tasks[taskid].Status = true
    }

    Tasks[taskid].Status = true
    Tasks[taskid].CompletedAt = time.Now()
}

func TaskEdit(taskid int){
    fmt.Println("task edited")
}

func TaskList(){
    t := table.New(os.Stdout)
    t.SetHeaders("TaskID", "Title", "Status", "CreateAt", "ModifiedAt", "CompletedAt", "Notify")

    for idx, it := range Tasks{
        title := tml.Sprintf("<yellow>%s</yellow>", it.Title)

        var isdone string
        var completedat string
        if !it.Status {
            isdone = "❌"
            completedat = ""
        }else{
            isdone = "💚"
            completedat = it.CompletedAt.Format(time.RFC1123)
        }

        t.AddRow(strconv.Itoa(idx) , title, isdone, it.CreatedAt.Format(time.RFC1123), it.ModifiedAt.Format(time.RFC1123), completedat, "❌")
    }
    t.Render()
}

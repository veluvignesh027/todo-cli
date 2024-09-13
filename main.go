package main

import (
        "encoding/json"
        "fmt"
        "os"
        "strconv"
)


const (
    OP_ADD = "add"
    OP_DEL = "del"
    OP_EDI = "edit"
    OP_TOG = "toggle"
    OP_LST = "list"
)

var Operation int
const Store string = "/tmp/tasks.json"

func init(){
    nbyte, err := os.ReadFile(Store)
    if os.IsNotExist(err){
        os.Create(Store)
        fmt.Println("Create a store")
        return
    } else if err != nil{
        panic(err)
    }

    err = json.Unmarshal(nbyte, &Tasks)
    if err != nil{
        fmt.Println(err)
    }
}


func main(){

    defer SaveBeforeExit()

    arg := os.Args

    if len(arg) <= 1 || len(arg) > 3 {
        usage()
    }


    switch arg[1] {
    case OP_ADD:
        if checkArg(arg, 2){
            TaskAdd(arg[2])
            break
        }
    case OP_DEL:
        if checkArg(arg, 2){
            idx, err := strconv.Atoi(arg[2])
            if err != nil{
                fmt.Println("Not a valid index, Please enter the correct index of the task to delete")
                TaskList()
                return
            }
            TaskDel(idx)
            break
        }

    case OP_EDI:
        if checkArg(arg, 2){
            idx, err := strconv.Atoi(arg[2])
            if err != nil{
                fmt.Println("Not a valid index, Please enter the correct index of the task to edit")
                TaskList()
                return
            }
            TaskEdit(idx)
            break
        }
    case OP_TOG:
        if checkArg(arg, 2){
            idx, err := strconv.Atoi(arg[2])
            if err != nil{
                fmt.Println("Not a valid index, Please enter the correct index of the task to edit")
                TaskList()
                return
            }
            TaskToggle(idx)
            break
        }
    case OP_LST:
        TaskList()
        break
    default:
        usage()
    }
}

func checkArg(arg []string, n int)bool{
    if len(arg) != (n+1) {
        return false
    }
    return true
}

func usage(){
    fmt.Println("USAGE: todo < add | delete | edit | toggle | list >")
    os.Exit(1)
}

func SaveBeforeExit(){
    nbyte, err := json.Marshal(Tasks)
    if err != nil{
        panic(err)
    }

    err = os.WriteFile(Store, nbyte, 0666)
    if err != nil{
        panic(err)
    }
}

package main

import (
  "os"
  "fmt"
)


func main(){
  dbfname := os.Args[1]
  //q := os.Args[2]
  db,err := loadDb(dbfname)
  if err != nil {
    panic(err)
  }
  db.printGraph()
  data, err := db.getDataFromLine(0,"nodes")
  if err != nil {
    panic(err)
  }
  fmt.Println(data.(map[string]interface{}))
  //db.Query(q)
}



package main

import (
  "os"
)


func main(){
  dbfname := os.Args[1]
  q := os.Args[2]
  db,err := LoadDb(dbfname)
  if err != nil {
    panic(err)
  }
  //db.PrintGraph()
  //data, err := db.GetDataFromLine(0,"nodes")
  //if err != nil {
  //  panic(err)
  //}
  //fmt.Println(data.(map[string]interface{}))
  db.Query(q)
}



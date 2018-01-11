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
  db.Query(q)
  db.SaveDb()
}



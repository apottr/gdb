package main

import (
  "os"
)


func main(){
  dbfname := os.Args[1]
  q := os.Args[2]
  db,err := loadDb(dbfname)
  if err != nil {
    panic(err)
  }
  db.Query(q)
}



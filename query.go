package main

import "fmt"



func QueryParser(qtext string) {
  //
}

func (db *Db) Query (q string) error {
  var g map[string]interface{}
  for i,_ := range db.graphs {
    d,err := db.GetDataFromLine(i,"graphs")
    if err != nil {
      return err
    }
    if d.(map[string]interface{})["name"].(string) == q {
      g = d.(map[string]interface{})
      break
    }
  }
  fmt.Println(g)
  for _,v := range g["links"].([]interface{}) {
    l,_ := db.GetDataFromLine(int(v.(float64)),"links")
    e,_ := db.GetDataFromLine(int(v.(float64)),"edges")
    fmt.Println(l,e)
  }
  return nil
}

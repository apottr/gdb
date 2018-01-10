package main

import (
  "fmt"
  "regexp"
)

func Query_InsertHandler(q map[string]string) error {
  
}

func Query_FindHandler(q map[string]string) error {
  
}

func QueryParser(qtext string) error {
  var err error
  re := regexp.MustCompile("'.+'|\".+\"|[^,]*")
  s := re.FindAllString(qtext,-1)
  q := map[string]string{
    "command": s[0],
    "graph": s[1],
    "edge": s[2],
    "node": s[3]
  }
  switch q["command"] {
    case "insert":
      err = Query_InsertHandler(q)
    case "query":
      err = Query_FindHandler(q)
  }

  if err != nil {
    return err
  }
  return nil
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

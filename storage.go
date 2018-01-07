package main

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "fmt"
  "bytes"
)

type Db struct {
  fname string
  edges [][]byte
  nodes [][]byte
  links [][]byte
}

func (db *Db) loadFile(ftype string) error {
  var b []byte
  var err error
  if _,err = os.Stat(fmt.Sprintf("%s/%s.jf",db.fname,ftype)); os.IsNotExist(err) {
    b = []byte{}
  }else{
    b, err = ioutil.ReadFile(fmt.Sprintf("%s/%s.jf",db.fname,ftype))
    if err != nil {
      return err
    }
  }
  switch ftype {
    case "edges":
      db.edges = bytes.Split(b,[]byte("\n"))
      break
    case "nodes":
      db.nodes = bytes.Split(b,[]byte("\n"))
      break
    case "links":
      db.links = bytes.Split(b,[]byte("\n"))
      break
  }
  return nil
}

func (db *Db) printGraph() {
  fmt.Printf("nodes: %s\nedges: %s\nlinks: %s\n",db.nodes,db.edges,db.links)
}

func (db *Db) getDataFromLine(line int, dtype string) (interface{},error) {
  var out interface{}
  var err error
  switch dtype {
    case "edges":
      err = json.Unmarshal(db.edges[line],&out)
    case "nodes":
      err = json.Unmarshal(db.nodes[line],&out)
    case "links":
      err = json.Unmarshal(db.links[line],&out)
  }
  return out,err
}

func loadDb(fname string) (*Db,error) {
  out := &Db{fname: fname}
  out.loadFile("edges")
  out.loadFile("nodes")
  out.loadFile("links")
  return out,nil
}

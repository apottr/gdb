package main

import (
  "io/ioutil"
  "os"
  "fmt"
)

type Db struct {
  fname string
  edges []byte
  nodes []byte
  links []byte
}

func (db *Db) loadFile(ftype string) error {
  var b []byte
  var err error
  if _,err = os.Stat(fmt.Sprintf("%s/%s.csv",db.fname,ftype)); os.IsNotExist(err) {
    b = []byte{}
  }else{
    b, err = ioutil.ReadFile(fmt.Sprintf("%s/%s.csv",db.fname,ftype))
    if err != nil {
      return err
    }
  }
  switch ftype {
    case "edges":
      db.edges = b
      break
    case "nodes":
      db.nodes = b
      break
    case "links":
      db.links = b
      break
  }
  return nil
}

func (db *Db) printGraph() {
  fmt.Printf("nodes: %s\nedges: %s\nlinks: %s\n",string(db.nodes),string(db.edges),string(db.links))
}

func loadDb(fname string) (*Db,error) {
  out := &Db{fname: fname}
  out.loadFile("edges")
  out.loadFile("nodes")
  out.loadFile("links")
  return out,nil
}

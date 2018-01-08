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
  graphs [][]byte
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
  data := bytes.Split(b,[]byte("\n"))
  switch ftype {
    case "edges":
      db.edges = data
      break
    case "nodes":
      db.nodes = data
      break
    case "links":
      db.links = data
      break
    case "graphs":
      db.graphs = data
  }
  return nil
}

func (db *Db) PrintGraph() {
  fmt.Printf("nodes: %s\nedges: %s\nlinks: %s\ngraphs: %s\n",db.nodes,db.edges,db.links,db.graphs)
}

func (db *Db) GetDataFromLine(line int, dtype string) (interface{},error) {
  var out interface{}
  var err error
  switch dtype {
    case "edges":
      err = json.Unmarshal(db.edges[line],&out)
      break
    case "nodes":
      err = json.Unmarshal(db.nodes[line],&out)
      break
    case "links":
      err = json.Unmarshal(db.links[line],&out)
      break
    case "graphs":
      err = json.Unmarshal(db.graphs[line],&out)
  }
  return out,err
}

func LoadDb(fname string) (*Db,error) {
  out := &Db{fname: fname}
  out.loadFile("edges")
  out.loadFile("nodes")
  out.loadFile("links")
  out.loadFile("graphs")
  return out,nil
}

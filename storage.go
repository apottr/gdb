package main

import (
  "io/ioutil"
  "gonum.org/v1/hdf5"
)

type Db struct {
  fname string
  path  string
  edges byte[]
  nodes byte[]
  links byte[]
}

func (db *Db) load_file(ftype string) err {
  b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.csv",db.path,ftype))
  if err != nil {
    err
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

func loadDb(fname string) (*Db,error) {
  
}

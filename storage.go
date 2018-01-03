package main

import (
  "os"
  "gonum.org/v1/hdf5"
)

type Db struct {
  fname string
  data hdf5.CommonFG
}

func loadDb(fname string) (*Db,error) {
  var f *hdf5.File
  var err error
  if _, err = os.Stat(fname); os.IsNotExist(err) {
    f, err = hdf5.CreateFile(fname,1)
    if err != nil {
      return &Db{},err
    }
  }else{
    f, err = hdf5.OpenFile(fname,1)
    if err != nil {
      return &Db{},err
    }
  }
  
  return &Db{},nil
}

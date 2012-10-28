package main

import (
  "runtime"
)
//TODO loaded, draw, destroy should be handled in a manager

type Object struct {
  model *Model
  mchan chan *Model
  matrix Matrix4
  loaded bool
}

func (o *Object) init(path string) (err error) {
  runtime.LockOSThread()
  o.loaded = false
  o.matrix.identity()

  o.mchan = make(chan *Model)
  go mm.getModel(path, o.mchan)

	return
}

func (o *Object) destroy() {
  if o.loaded {
    o.model.destroy()
  }
}

func (o* Object) draw() {
  if o.loaded {
    o.model.setMatrix(o.matrix)
    o.model.draw()
  }
}

/*
func (o *Object) waitModel(c chan *Model) {
  runtime.LockOSThread()
  select {
  case o.model = <-c:
    fmt.Println("I RECEIVVVVVVVVVVVVVVVVVVVVVE")
    if !o.loaded {
      o.model.init()
      o.loaded = true
    }
  }
}
*/

func (o *Object) update() {
  select {
  case o.model = <-o.mchan:
    if !o.loaded {
      o.model.init()
      o.loaded = true
    }
  default:
    //fmt.Println("nothing received")
  }
}




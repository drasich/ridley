package ridley

import (
  "runtime"
)
//TODO loaded, draw, destroy should be handled in a manager

type Updatable interface {
  update()
}

type Drawable interface {
  draw()
}

type Collidable interface {
  collide(c* Collidable)
}

//TODO rename updatecomponent
type Component interface {
  Update()
}


type Object struct {
  model *Model
  mchan chan *Model
  Matrix Matrix4
  loaded bool
  Position Vec3
  Orientation Quat
  components []Component
}

func (o *Object) Init(path string) (err error) {
  runtime.LockOSThread()
  o.loaded = false
  o.Matrix.identity()

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
    o.model.setMatrix(o.Matrix)
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

  for _,c := range o.components {
    //(*c).Update()
    c.Update()
  }

  o.Matrix.Translation(o.Position.X,o.Position.Y,o.Position.Z-7)
 // o.matrix.rotate(-90, 1,0,0)
}

func (o *Object) AddComponent(c Component) {
  o.components = append(o.components,c)
}



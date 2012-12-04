package ridley

import (
  "runtime"
	"github.com/jteeuwen/glfw"
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

type Object struct {
  model *Model
  mchan chan *Model
  Matrix Matrix4
  loaded bool
  position Vec3
  orientation Quat
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

  o.control()
  o.Matrix.Translation(o.position.X,o.position.Y,o.position.Z-7)
 // o.matrix.rotate(-90, 1,0,0)
}


func (o *Object) control(){
  if glfw.Key('E') == glfw.KeyPress {
    o.position.Z -= 0.1;
  } else if glfw.Key('D') == glfw.KeyPress {
    o.position.Z += 0.1;
  } else if glfw.Key('S') == glfw.KeyPress {
    o.position.X -= 0.1;
  } else if glfw.Key('F') == glfw.KeyPress {
    o.position.X += 0.1;
  }

}



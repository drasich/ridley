package ridley

import (
  //"runtime"
)

//TODO draw, destroy should be handled in a manager

type Object struct {
  Matrix Matrix4
  Position Vec3
  Orientation Quat
  components []Component
  Mesh *MeshComponent
}

func (o *Object) Init() (err error) {
  //runtime.LockOSThread()
  o.Matrix.identity()

	return
}


func (o *Object) destroy() {
  //TODO destroy mesh and other components
}

func (o* Object) draw() {
  if o.Mesh != nil {
    o.Mesh.draw()
  }
}

func (o *Object) update() {
  for _,c := range o.components {
    //(*c).Update()
    c.Update()
  }

  if o.Mesh != nil {
    o.Mesh.update()
  }

  o.Matrix.Translation(o.Position.X,o.Position.Y,o.Position.Z-7)
 // o.matrix.rotate(-90, 1,0,0)
}

func (o *Object) AddComponent(c Component) {
  o.components = append(o.components,c)
}



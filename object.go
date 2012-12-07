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
  Box *BoxComponent
}

func (o *Object) Init() (err error) {
  //runtime.LockOSThread()
  o.Matrix.identity()
  o.Orientation.identity()
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

  mt := Matrix4Translation(o.Position)
  mr := Matrix4Quat(o.Orientation)

  //TODO change the multiply function
  newmat := mt.multiply(&mr)
  o.Matrix = newmat
}

func (o *Object) AddComponent(c Component) {
  o.components = append(o.components,c)
}



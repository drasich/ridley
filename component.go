package ridley

import (
  //"runtime"
)
//TODO rename updatecomponent?
type Component interface {
  Update()
}

type BoxComponent struct {
  size Vec3
  offset Vec3
}

func NewBoxComponent(size Vec3, offset Vec3) *BoxComponent {
  return &BoxComponent{size, offset}
}

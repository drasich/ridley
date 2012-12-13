package ridley

import (
  //"runtime"
)
//TODO rename updatecomponent?
type Component interface {
  Update()
}

type BoxComponent struct {
  box AABox
  //TODO not used for now
  //offset Vec3
  //orientation Quat
}

func NewBoxComponent(size Vec3, offset Vec3) *BoxComponent {
  return &BoxComponent{AABox{size, offset}}
}

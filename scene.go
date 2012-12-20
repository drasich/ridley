package ridley


import (
	gl "github.com/chsc/gogl/gl21"
  //"fmt"
  "runtime"
)


type Scene struct {
  Name string
  Id int
  Objects []*Object
}

var (
	rotx, roty  float64
	//ambient     []gl.Float  = []gl.Float{0.5, 0.5, 0.5, 1}
	//diffuse     []gl.Float  = []gl.Float{1, 1, 1, 1}
	//lightpos    []gl.Float  = []gl.Float{-5, 5, 10, 0}

  //model = new(Model)
  camera = new(Camera)
  mm ModelManager
  sphere *Object
  test *Object

  //TODO remove this from here
  projection = MakeFrustum(-1, 1, -1, 1, 1, 100.0)
)

func (s *Scene) Init() (err error) {
  runtime.LockOSThread()
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.DEPTH_TEST)

	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.ClearDepth(1)
	//gl.DepthFunc(gl.LEQUAL)

	gl.Viewport(0, 0, Width, Height)

  mm.init()

	return
}

func (s* Scene) Destroy() {
  for _,o := range s.Objects {
    o.destroy()
  } 
}


func (s *Scene) Draw() {
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
  for _,o := range s.Objects {
    o.draw()
  } 
}

func (s *Scene) Update() {
  runtime.LockOSThread()

  for _,o := range s.Objects {
    o.update()
  } 
}



func (s *Scene) AddObject(o *Object) (err error) {
  s.Objects = append(s.Objects,o)
  return
}


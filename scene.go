package main


import (
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
  //"fmt"
  "runtime"
)


type scene struct {
  name string
  id int
  objects []*Object
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

  projection = MakeFrustum(-1, 1, -1, 1, 1, 100.0)
)

func initScene() (err error) {
  runtime.LockOSThread()
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.DEPTH_TEST)

	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.ClearDepth(1)
	//gl.DepthFunc(gl.LEQUAL)

	gl.Viewport(0, 0, Width, Height)

  mm.init()

  sphere = new(Object)
  sphere.init("model/tex.bin")

  test = new(Object)
  test.init("model/tex.bin")

  var mat Matrix4

  mat.translation(0,0,-7)
  mat.rotate(-rotx, 0,1,0)
  mat.rotate(-90, 1,0,0)
  test.matrix = mat

	return
}

func destroyScene() {
  test.destroy()
  sphere.destroy()
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
  test.draw()
  //sphere.draw()
}


func updateScene() {
  runtime.LockOSThread()

  rotx +=5
  if glfw.Key(glfw.KeyEsc) == glfw.KeyPress {
    exit = true
  }
  /* else if glfw.Key('G') == glfw.KeyPress && !sent {
    go mm.getModel("model/dsphere.bin", smchan)
    sent = true
  }
  */


  var mat Matrix4

  mat.translation(0,-4,-7)
  //fmt.Println("yep : ", rotx)
  mat.rotate(rotx, 0,1,0)
  mat.rotate(-90, 1,0,0)

  sphere.matrix = mat

  mat.translation(0,0,-7)
  mat.rotate(-rotx, 0,1,0)
  mat.rotate(-90, 1,0,0)
  //test.matrix = mat

  sphere.update()
  test.update()

}

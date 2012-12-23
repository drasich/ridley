package ridley

import (
  "os"
	"fmt"
	gl "github.com/chsc/gogl/gl21"
  //"os"
  "log"
  "bytes"
  "encoding/binary"
  //"reflect"
  "unsafe"
  "sync"
)

type ModelReader struct {
  file *os.File
}

//type ModelReader os.File

func (mr* ModelReader) init(file *os.File) {
  mr.file = file
}

func readModel(path string, c chan<- *Model) {
  fmt.Println("readModel : ", path)

  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }

  var mr ModelReader
  mr.init(file)

  m := new(Model)

  mr.readModel(path, m)
  c <- m
}

func (mr *ModelReader) readuint16(value *uint16) {
  size := uint32(unsafe.Sizeof(*value))
  //fmt.Println("type : ", reflect.TypeOf(*value))

  data := make([]byte,size)
  _, erro := mr.file.Read(data)
  if erro != nil {
    log.Fatal(erro)
  }

  buf := bytes.NewBuffer(data[0:size])
  erro = binary.Read(buf, binary.LittleEndian, value)
  if erro != nil {
    fmt.Println("binary.Read failed:", erro)
  }

  //switch t := value.(type) {
	//case float32:

}

func (mr *ModelReader) readfloat32(value *float32) {
  size := uint32(unsafe.Sizeof(*value))

  data := make([]byte,size)
  _, erro := mr.file.Read(data)
  if erro != nil {
    log.Fatal(erro)
  }

  buf := bytes.NewBuffer(data[0:size])
  erro = binary.Read(buf, binary.LittleEndian, value)
  if erro != nil {
    fmt.Println("binary.Read failed:", erro)
  }
}


func (mr *ModelReader) readModel(path string, m* Model) {

  var num uint16
  mr.readuint16(&num)

  for i := 0 ; i <int(num); i++ {
    var x,y,z float32
    mr.readfloat32(&x)
    mr.readfloat32(&y)
    mr.readfloat32(&z)

    m.vertices = append(m.vertices,gl.Float(x),gl.Float(y),gl.Float(z))
    //fmt.Printf("vertices %f, %f, %f \n", x, y, z)
  }
  //fmt.Printf("m.vertices size %d \n", len(m.vertices))

  var faces uint16
  mr.readuint16(&faces)
  //fmt.Println("nb faces  : ", faces)

  for i := 0 ; i <int(faces); i++ {
    var x, y, z uint16
    mr.readuint16(&x)
    mr.readuint16(&y)
    mr.readuint16(&z)

    m.indices = append(m.indices,gl.Uint(x),gl.Uint(y),gl.Uint(z))
    //fmt.Printf("face indices :%d, %d, %d \n", x, y, z)
  }

  var normals_num uint16
  mr.readuint16(&normals_num)

  for i := 0 ; i <int(normals_num); i++ {
    var x,y,z float32
    mr.readfloat32(&x)
    mr.readfloat32(&y)
    mr.readfloat32(&z)

    m.normals = append(m.normals,gl.Float(x),gl.Float(y),gl.Float(z))
    //fmt.Printf("normal %f, %f, %f \n", x, y, z)
  }

  var uvs_num uint16
  mr.readuint16(&uvs_num)
  //fmt.Printf("uvs num  %d \n", uvs_num)

  m.has_uv = uvs_num > 0

  if m.has_uv {
    for i := 0 ; i <int(uvs_num); i++ {
      var x,y float32
      mr.readfloat32(&x)
      mr.readfloat32(&y)

      m.uvs = append(m.uvs,gl.Float(x),gl.Float(y))
      //fmt.Printf("uvs %f, %f \n", x, y)
    }
  }

}


type ModelManager struct {
  //loaded map[string]*Model
  request_acc RequestAccess
  loading_acc LoadingAccess
  loaded_acc LoadedAccess
}

func (mm* ModelManager) init() {
  //mm.acc.res = new(Resource)
  mm.loaded_acc.loaded = make(map[string]*Model)
  mm.request_acc.request = make(map[string][]chan *Model)
}


func (mm* ModelManager) getModel(path string, c chan *Model) {
  mm.loaded_acc.RLock()
  mdl, ok := mm.loaded_acc.loaded[path]
  mm.loaded_acc.RUnlock()
  if ok {
    c <- mdl
    return
  }

  mm.addRequest(path, c)

  // check if it's already loading
  mm.loading_acc.RLock()
  for _, p := range mm.loading_acc.loading {
    if p == path {
      mm.loading_acc.RUnlock()
      return
    }
  }
  mm.loading_acc.RUnlock()

  mm.loading_acc.Lock()
  index := len(mm.loading_acc.loading)
  mm.loading_acc.loading = append(mm.loading_acc.loading, path)
  mm.loading_acc.Unlock()

  anotherchan := make(chan *Model)
  go readModel(path, anotherchan)
  go mm.FulfillRequests(path, anotherchan,index)
}

func (mm* ModelManager) FulfillRequests(path string, c <-chan *Model, loading_index int) {
  select {
  case mdl := <-c:
    mm.loaded_acc.Lock()
    mm.loaded_acc.loaded[path] = mdl
    mm.loaded_acc.Unlock()

    //remove from loading
    mm.loading_acc.Lock()
    l := len(mm.loading_acc.loading)
    mm.loading_acc.loading[loading_index] = mm.loading_acc.loading[l-1]
    mm.loading_acc.loading = mm.loading_acc.loading[0:l-1]
    mm.loading_acc.Unlock()


    mm.request_acc.Lock()
    for _,v := range mm.request_acc.request[path]{
      v <- mdl
    }
    //mm.request_acc.request[path] = mm.request_acc.request[path][0:0]
    delete(mm.request_acc.request, path)
    mm.request_acc.Unlock()
  }

}

func (mm* ModelManager) addRequest(path string, c chan *Model) {
  mm.request_acc.Lock()
  mm.request_acc.request[path] = append(mm.request_acc.request[path], c)
  mm.request_acc.Unlock()
}


// struct wraps pointer to resource with embedded RWMutex, thus
// acquiring RWMutex methods
type LoadedAccess struct {
  sync.RWMutex
  loaded map[string]*Model
}

type RequestAccess struct {
  sync.RWMutex
  request map[string][]chan *Model
}

type LoadingAccess struct {
  sync.RWMutex
  loading []string
}




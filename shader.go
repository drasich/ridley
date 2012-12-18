package ridley

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
  "os"
  "log"
  "io/ioutil"
)

type Shader struct {
  vert_shader gl.Uint
  frag_shader gl.Uint
  program gl.Uint
  attribute_vertex gl.Uint
  attribute_normal gl.Uint
  uniform_test gl.Int
  uniform_matrix gl.Int
  uniform_normal_matrix gl.Int

  uniform_texture gl.Int
  attribute_texcoord gl.Uint
}

func (s* Shader) initDefault() {
  vert_shader := `varying vec4 color;
  void main(void)
  {
    color = gl_Vertex;
    gl_Position = gl_ModelViewProjectionMatrix * gl_Vertex;
  }
  `

  frag_shader := `
  void main (void)
  {
    gl_FragColor = vec4(0,0,1,1);
  }
  `
  s.init(vert_shader, frag_shader)
  //str = stringFromFile("shader/simple.vert")
  str, _ := stringFromAllFile("shader/simple.vert")
  fmt.Println("from all file : ", str)
}

func stringFromFile(path string) string {
  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  data := make([]byte, 1024)
  count, err := file.Read(data)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("tep count : ", count)
  fmt.Println("tep data : ", data[0:count])

  return string(data)
}

func stringFromAllFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s* Shader) init(vert_path string, frag_path string) {
  vert, _ := stringFromAllFile(vert_path)
  frag, _ := stringFromAllFile(frag_path)
  s.initWithString(vert, frag)
}

func (s *Shader) initShader(t gl.Enum, str string, shader *gl.Uint) {
  *shader = gl.CreateShader(t)
  if *shader == 0 {
    fmt.Println("error creating shader of type ", t)
  }
  
  src := gl.GLStringArray(str)
	defer gl.GLStringArrayFree(src)

  gl.ShaderSource(*shader, 1, &src[0], nil);
  gl.CompileShader(*shader);

  var (
		status gl.Int
		info_length gl.Int
		message *gl.Char
	)

  gl.GetShaderiv(*shader, gl.COMPILE_STATUS, &status)
  if status == gl.FALSE {
    fmt.Println("Error compiling shader")
    gl.GetShaderiv(*shader, gl.INFO_LOG_LENGTH, &info_length)
		message = gl.GLStringAlloc(gl.Sizei(info_length))
		gl.GetShaderInfoLog(*shader, gl.Sizei(info_length), nil, message)
		fmt.Println(gl.GoString(message))
		gl.GLStringFree(message)
	}
  
}

func (s* Shader) initWithString(vert_shader string, frag_shader string) {

  s.initShader(gl.VERTEX_SHADER, vert_shader, &s.vert_shader)
  s.initShader(gl.FRAGMENT_SHADER, frag_shader, &s.frag_shader)

  s.program = gl.CreateProgram()
  fmt.Println("program : ", s.program)

  gl.AttachShader(s.program, s.vert_shader)
	gl.AttachShader(s.program, s.frag_shader)
	gl.LinkProgram(s.program)

  var (
		status gl.Int
		info_length gl.Int
		message *gl.Char
	)

	gl.GetProgramiv(s.program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		fmt.Println("Error linking program")
		gl.GetProgramiv(s.program, gl.INFO_LOG_LENGTH, &info_length)
		message = gl.GLStringAlloc(gl.Sizei(info_length))
		gl.GetProgramInfoLog(s.program, gl.Sizei(info_length), nil, message)
		fmt.Println(gl.GoString(message))
		gl.GLStringFree(message)
	}

  s.initAttributes()
  s.initUniforms()

}

func (s* Shader) initAttribute(name string, att *gl.Uint ){
  att_name := gl.GLString(name)
  defer gl.GLStringFree(att_name)
  att_tmp := gl.GetAttribLocation(s.program, att_name)
  if att_tmp == -1 {
		fmt.Println("Error in getting attribute ", gl.GoString(att_name))
	} else {
    *att = gl.Uint(att_tmp)
  }
}

func (s* Shader) initAttributes() {
  s.initAttribute("vertex", &s.attribute_vertex)
  s.initAttribute("normal", &s.attribute_normal)
  s.initAttribute("texcoord", &s.attribute_texcoord)
}

func (s *Shader) initUniform(name string, uni *gl.Int) {
  uni_name := gl.GLString(name)
  defer gl.GLStringFree(uni_name)
  *uni  = gl.GetUniformLocation(s.program, uni_name)
  if *uni == -1 {
		fmt.Println("Shit Error in getting uniform", gl.GoString(uni_name))
  }
}

func (s* Shader) initUniforms() {
  s.initUniform("matrix", &s.uniform_matrix)
  s.initUniform("normal_matrix", &s.uniform_normal_matrix)
}

func (s* Shader) destroy() {
  gl.DeleteShader(s.vert_shader)
  gl.DeleteShader(s.frag_shader)
  gl.DeleteProgram(s.program)
}

func (s* Shader) use() {
  gl.UseProgram(s.program)
}


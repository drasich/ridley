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

func (s* Shader) initWithString(vert_shader string, frag_shader string) {

  s.vert_shader = gl.CreateShader(gl.VERTEX_SHADER)
  if s.vert_shader == 0 {
    fmt.Println("error creating vertex shader")
  }
  fmt.Println("shader : ", s.vert_shader)

  s.frag_shader = gl.CreateShader(gl.FRAGMENT_SHADER)

  if s.frag_shader == 0 {
    fmt.Println("error creating fragment shader")
  }

  vert_src := gl.GLStringArray(vert_shader)
	defer gl.GLStringArrayFree(vert_src)

  gl.ShaderSource(s.vert_shader, 1, &vert_src[0], nil);
  gl.CompileShader(s.vert_shader);

  var (
		status     gl.Int
		info_length gl.Int
		message    *gl.Char
	)

  gl.GetShaderiv(s.vert_shader, gl.COMPILE_STATUS, &status)
  if status == gl.FALSE {
    fmt.Println("Error compiling vertex shader")
    gl.GetShaderiv(s.vert_shader, gl.INFO_LOG_LENGTH, &info_length)
		message = gl.GLStringAlloc(gl.Sizei(info_length))
		gl.GetShaderInfoLog(s.vert_shader, gl.Sizei(info_length), nil, message)
		fmt.Println(gl.GoString(message))
		gl.GLStringFree(message)
	}

  frag_src := gl.GLStringArray(frag_shader)
	defer gl.GLStringArrayFree(frag_src)

  gl.ShaderSource(s.frag_shader, 1, &frag_src[0], nil);
  gl.CompileShader(s.frag_shader);

  gl.GetShaderiv(s.frag_shader, gl.COMPILE_STATUS, &status)
  if status == gl.FALSE {
    fmt.Println("Error compiling frag shader")
    gl.GetShaderiv(s.frag_shader, gl.INFO_LOG_LENGTH, &info_length)
		message = gl.GLStringAlloc(gl.Sizei(info_length))
		gl.GetShaderInfoLog(s.frag_shader, gl.Sizei(info_length), nil, message)
		fmt.Println(gl.GoString(message))
		gl.GLStringFree(message)
	}

  s.program = gl.CreateProgram()
  fmt.Println("program : ", s.program)

  gl.AttachShader(s.program, s.vert_shader)
	gl.AttachShader(s.program, s.frag_shader)
	gl.LinkProgram(s.program)

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
  s.initUniform()

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

func (s* Shader) initUniform() {
  uniform_name := gl.GLString("test")
  defer gl.GLStringFree(uniform_name)
  s.uniform_test  = gl.GetUniformLocation(s.program, uniform_name)
  if s.uniform_test == -1 {
		fmt.Println("Shit Error in getting uniform", gl.GoString(uniform_name))
  }

  //setting the uniform
  s.use()
  gl.Uniform1f(s.uniform_test, 1.0)

  uniform_name = gl.GLString("matrix")
  s.uniform_matrix  = gl.GetUniformLocation(s.program, uniform_name)
  if s.uniform_matrix == -1 {
		fmt.Println("Error in getting uniform", gl.GoString(uniform_name))
  }
  fmt.Println("no error ", s.uniform_matrix)


  uniform_name = gl.GLString("normal_matrix")
  s.uniform_normal_matrix  = gl.GetUniformLocation(s.program, uniform_name)
  if s.uniform_normal_matrix == -1 {
		fmt.Println("Error in getting uniform", gl.GoString(uniform_name))
  } else {
    fmt.Println("no error ", s.uniform_normal_matrix)
  }

}

func (s* Shader) destroy() {
  gl.DeleteShader(s.vert_shader)
  gl.DeleteShader(s.frag_shader)
  gl.DeleteProgram(s.program)
}

func (s* Shader) use() {
  gl.UseProgram(s.program)
}


package main

import (
	gl "github.com/chsc/gogl/gl21"
  "unsafe"
  "fmt"

  "image"
  "os"
  "log"
)

type Model struct {
  buffer      gl.Uint
  index      gl.Uint
  normal_buf gl.Uint
  texture_id gl.Uint
  texcoord gl.Uint

  vertices        []gl.Float
  indexes       []gl.Uint
  normals       []gl.Float
  uvs       []gl.Float
  shader Shader

  ngl Matrix3GLFloat
  glm Matrix4GLFloat

}

func (m *Model) init() (err error) {
  m.shader.init("shader/simple.vert", "shader/simple.frag")

  gl.GenBuffers(1, &m.buffer)
  fmt.Println("init buffer : ", m.buffer)
  gl.BindBuffer(gl.ARRAY_BUFFER, m.buffer);
  gl.BufferData(
    gl.ARRAY_BUFFER,
    gl.Sizeiptr(len(m.vertices)* int(unsafe.Sizeof(m.vertices[0]))),
    gl.Pointer(&m.vertices[0]),
    gl.STATIC_DRAW);

  gl.GenBuffers(1, &m.index);
  fmt.Println("init index : ", m.index)
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.index);
  gl.BufferData(
    gl.ELEMENT_ARRAY_BUFFER,
    gl.Sizeiptr(len(m.indexes)*int(unsafe.Sizeof(m.indexes[0]))),
    gl.Pointer(&m.indexes[0]),
    gl.STATIC_DRAW);

  gl.GenBuffers(1, &m.normal_buf);
  gl.BindBuffer(gl.ARRAY_BUFFER, m.normal_buf);
  gl.BufferData(
    gl.ARRAY_BUFFER,
    gl.Sizeiptr(len(m.normals)* int(unsafe.Sizeof(m.normals[0]))),
    gl.Pointer(&m.normals[0]),
    gl.STATIC_DRAW);

  gl.BindBuffer(gl.ARRAY_BUFFER, m.buffer);
  gl.EnableVertexAttribArray(m.shader.attribute_vertex);
  gl.VertexAttribPointer(
    m.shader.attribute_vertex,
    3,
    gl.FLOAT,
    gl.FALSE,
    0,
    gl.Pointer(nil));

  gl.BindBuffer(gl.ARRAY_BUFFER, m.normal_buf);
  gl.EnableVertexAttribArray(m.shader.attribute_normal);
  gl.VertexAttribPointer(
    m.shader.attribute_normal,
    3,
    gl.FLOAT,
    gl.FALSE,
    0,
    gl.Pointer(nil));


  //gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.index);
  m.initTexture()

  gl.GenBuffers(1, &m.texcoord)
  gl.BindBuffer(gl.ARRAY_BUFFER, m.texcoord)
  gl.BufferData(
    gl.ARRAY_BUFFER,
    gl.Sizeiptr(len(m.uvs)* int(unsafe.Sizeof(m.uvs[0]))),
    gl.Pointer(&m.uvs[0]),
    gl.STATIC_DRAW);

  gl.BindBuffer(gl.ARRAY_BUFFER, 0);
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0);
  gl.DisableVertexAttribArray(m.shader.attribute_vertex);
  gl.DisableVertexAttribArray(m.shader.attribute_normal);

	return
}

func (m* Model) initTexture() {
  file, err := os.Open("model/test.png")
  //file, err := os.Open("model/ceil.png")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  img, _, err := image.Decode(file)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("color model : ", img.ColorModel())

  rgbaImg, ok := img.(*image.NRGBA)
	if !ok {
    //return 0, errors.New("texture must be an NRGBA image")
    log.Fatal("image is not rgba")
  }

	gl.GenTextures(1, &m.texture_id)
	gl.BindTexture(gl.TEXTURE_2D, m.texture_id)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	// flip image: first pixel is lower left corner
	img_width, img_height := img.Bounds().Dx(), img.Bounds().Dy()
	data := make([]byte, img_width * img_height * 4)
	line_len := img_width * 4
	dest := len(data)-line_len
	for src := 0; src < len(rgbaImg.Pix); src+=rgbaImg.Stride {
		copy(data[dest:dest+line_len], rgbaImg.Pix[src:src+rgbaImg.Stride])
		dest-=line_len
	}
	gl.TexImage2D(
    gl.TEXTURE_2D,
    0,
    gl.RGBA, //4,
    gl.Sizei(img_width),
    gl.Sizei(img_height),
    0,
    gl.RGBA,
    gl.UNSIGNED_BYTE,
    gl.Pointer(&data[0]))

}

func (m* Model) destroy() {
  gl.DeleteBuffers(1, &m.buffer);
  gl.DeleteBuffers(1, &m.index);
  gl.DeleteTextures(1, &m.texture_id)
  m.shader.destroy()
}

func (m* Model) setMatrix(mat Matrix4) {
  m.shader.use()

  var normal_mat Matrix3

  normal_mat = mat.toMatrix3()
  normal_mat = normal_mat.getInverse()
  //already transposed
  //normal_mat = normal_mat.getTranspose()
  m.ngl = normal_mat.toGLFloat()

  var tm Matrix4
  tm  = projection.multiply(&mat)
  tm = tm.getTranspose()
  m.glm = tm.toGLFloat()
  gl.UniformMatrix4fv(m.shader.uniform_matrix, 1, gl.FALSE, &m.glm[0])
  gl.UniformMatrix3fv(m.shader.uniform_normal_matrix, 1, gl.FALSE, &m.ngl[0])
}

func (m* Model) draw() {
  m.shader.use()

  gl.ActiveTexture(gl.TEXTURE0)
  gl.BindTexture(gl.TEXTURE_2D, m.texture_id)
  gl.Uniform1i(m.shader.uniform_texture, 0)

  //texcoord
  gl.EnableVertexAttribArray(m.shader.attribute_texcoord);
  gl.BindBuffer(gl.ARRAY_BUFFER, m.texcoord);
  gl.VertexAttribPointer(
    m.shader.attribute_texcoord,
    2,
    gl.FLOAT,
    gl.FALSE,
    0,
    gl.Pointer(nil));


  //gl.BindBuffer(gl.ARRAY_BUFFER, m.buffer);

  gl.BindBuffer(gl.ARRAY_BUFFER, m.buffer);
  gl.EnableVertexAttribArray(m.shader.attribute_vertex);

  gl.VertexAttribPointer(
    m.shader.attribute_vertex,
    3,
    gl.FLOAT,
    gl.FALSE,
    0,
    gl.Pointer(nil));

  gl.BindBuffer(gl.ARRAY_BUFFER, m.normal_buf);
  gl.EnableVertexAttribArray(m.shader.attribute_normal);
  gl.VertexAttribPointer(
    m.shader.attribute_normal,
    3,
    gl.FLOAT,
    gl.FALSE,
    0,
    gl.Pointer(nil));

  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.index);
  gl.DrawElements(gl.TRIANGLES, gl.Sizei(len(m.indexes)), gl.UNSIGNED_INT, gl.Pointer(nil));
  //gl.DrawElements(gl.TRIANGLES, gl.Sizei(12), gl.UNSIGNED_INT, gl.Pointer(nil));
  gl.DisableVertexAttribArray(m.shader.attribute_vertex);
  gl.DisableVertexAttribArray(m.shader.attribute_normal);
  gl.DisableVertexAttribArray(m.shader.attribute_texcoord);

  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0);

}


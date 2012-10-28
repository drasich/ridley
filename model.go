package main

import (
	gl "github.com/chsc/gogl/gl21"
  "unsafe"
  "fmt"
)

type Model struct {
  buffer      gl.Uint
  index      gl.Uint
  normal_buf gl.Uint
  pos        []gl.Float
  indexes       []gl.Uint
  normals       []gl.Float
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
    gl.Sizeiptr(len(m.pos)* int(unsafe.Sizeof(m.pos[0]))),
    gl.Pointer(&m.pos[0]),
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

  gl.BindBuffer(gl.ARRAY_BUFFER, 0);
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0);
  gl.DisableVertexAttribArray(m.shader.attribute_vertex);
  gl.DisableVertexAttribArray(m.shader.attribute_normal);

  //gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.index);

	return
}

func (m* Model) destroy() {
  gl.DeleteBuffers(1, &m.buffer);
  gl.DeleteBuffers(1, &m.index);
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
  gl.DisableVertexAttribArray(m.shader.attribute_vertex);
  gl.DisableVertexAttribArray(m.shader.attribute_normal);

  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0);

}


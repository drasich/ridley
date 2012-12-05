package ridley

import (
  "math"
	gl "github.com/chsc/gogl/gl21"
//  "fmt"
)

const DegToRad = math.Pi / 180.0

/*
type Matrix4 struct {
  m [16]float64
}
*/
type Matrix4 [16]float64
type Matrix4GL [16]gl.Double
type Matrix4GLFloat [16]gl.Float

type Matrix3 [9]float64
type Matrix3GLFloat [9]gl.Float
/*
func createMatrix2() (m matrix2)  {
  return m
}
*/

func (m *Matrix4) copy(n* Matrix4) {
  for i := range m {
    m[i] = n[i];
  }
}

func (m *Matrix4) translate(x, y, z float64) {

    m[0] += m[12]*x
    m[1] += m[13]*x
    m[2] += m[14]*x
    m[3] += m[15]*x

    m[4] += m[12]*y
    m[5] += m[13]*y
    m[6] += m[14]*y
    m[7] += m[15]*y

    m[8] += m[12]*z
    m[9] += m[13]*z
    m[10]+= m[14]*z
    m[11]+= m[15]*z
}

func  Matrix4Identity() Matrix4 {
  var mat Matrix4
  mat.identity()
  return mat
}

func (m *Matrix4) identity() {
  m[0], m[5], m[10], m[15] = 1,1,1,1
  
  m[1], m[2], m[3], m[4],
  m[6], m[7], m[8], m[9],
  m[11], m[12], m[13], m[14] = 0,0,0,0,0,0,0,0,0,0,0,0
}

func (m *Matrix4) Translation(x,y,z float64) {
  m[0], m[5], m[10], m[15] = 1,1,1,1
  m[3], m[7], m[11] = x, y, z

  m[1], m[2], m[4],
  m[6], m[8], m[9],
  m[12], m[13], m[14] = 0,0,0,0,0,0,0,0,0
}

func Matrix4Rotation(angle,x,y,z float64) (m Matrix4)  {

  c := math.Cos(angle * DegToRad)
  s := math.Sin(angle * DegToRad)
  xx := x * x
  xy := x * y
  xz := x * z
  yy := y * y
  yz := y * z
  zz := z * z

  // build rotation matrix
  //var mat Matrix4
  m[0] = xx * (1 - c) + c
  m[1] = xy * (1 - c) - z * s
  m[2] = xz * (1 - c) + y * s
  m[3] = 0
  m[4] = xy * (1 - c) + z * s
  m[5] = yy * (1 - c) + c
  m[6] = yz * (1 - c) - x * s
  m[7] = 0
  m[8] = xz * (1 - c) - y * s
  m[9] = yz * (1 - c) + x * s
  m[10]= zz * (1 - c) + c
  m[11]= 0
  m[12]= 0
  m[13]= 0
  m[14]= 0
  m[15]= 1

  return m
}

func (m *Matrix4) multiply(n *Matrix4) Matrix4  {

  var newmat Matrix4

  newmat[0] = m[0]*n[0]  + m[1]*n[4]  + m[2]*n[8]  + m[3]*n[12]
  newmat[1] = m[0]*n[1]  + m[1]*n[5]  + m[2]*n[9]  + m[3]*n[13]
  newmat[2] = m[0]*n[2]  + m[1]*n[6]  + m[2]*n[10]  + m[3]*n[14]
  newmat[3] = m[0]*n[3]  + m[1]*n[7]  + m[2]*n[11]  + m[3]*n[15]
  newmat[4] = m[4]*n[0]  + m[5]*n[4]  + m[6]*n[8]  + m[7]*n[12]
  newmat[5] = m[4]*n[1]  + m[5]*n[5]  + m[6]*n[9]  + m[7]*n[13]
  newmat[6] = m[4]*n[2]  + m[5]*n[6]  + m[6]*n[10]  + m[7]*n[14]
  newmat[7] = m[4]*n[3]  + m[5]*n[7]  + m[6]*n[11]  + m[7]*n[15]
  newmat[8] = m[8]*n[0]  + m[9]*n[4]  + m[10]*n[8] + m[11]*n[12]
  newmat[9] = m[8]*n[1]  + m[9]*n[5]  + m[10]*n[9] + m[11]*n[13]
  newmat[10] = m[8]*n[2]  + m[9]*n[6]  + m[10]*n[10] + m[11]*n[14]
  newmat[11] = m[8]*n[3]  + m[9]*n[7]  + m[10]*n[11] + m[11]*n[15]
  newmat[12] = m[12]*n[0] + m[13]*n[4] + m[14]*n[8] + m[15]*n[12]
  newmat[13] = m[12]*n[1] + m[13]*n[5] + m[14]*n[9] + m[15]*n[13]
  newmat[14] = m[12]*n[2] + m[13]*n[6] + m[14]*n[10] + m[15]*n[14]
  newmat[15] = m[12]*n[3] + m[13]*n[7] + m[14]*n[11] + m[15]*n[15]

  return newmat
}

func (mat *Matrix4) Rotate(angle,x,y,z float64)  {
  var rot Matrix4 = Matrix4Rotation(angle,x,y,z)
  //fmt.Println("rot : ",  rot)
  rot = mat.multiply(&rot)
  mat.copy(&rot)
}

func (m *Matrix4) getTranspose() Matrix4 {
  var tm Matrix4

  tm[0] = m[0]
  tm[1] = m[4]
  tm[2] = m[8]
  tm[3] = m[12]

  tm[4] = m[1]
  tm[5] = m[5]
  tm[6] = m[9]
  tm[7] = m[13]

  tm[8] = m[2]
  tm[9] = m[6]
  tm[10]= m[10]
  tm[11]= m[14]

  tm[12]= m[3]
  tm[13]= m[7]
  tm[14]= m[11]
  tm[15]= m[15]

  return tm

}

func (m *Matrix4) toGL() Matrix4GL {
  var glm Matrix4GL

  for i := range m {
    glm[i] = gl.Double(m[i])
  }

  return glm
}

func (m *Matrix4) toGLFloat() Matrix4GLFloat {
  var glm Matrix4GLFloat

  for i := range m {
    glm[i] = gl.Float(m[i])
  }

  return glm
}


func MakeFrustum(left, right, bottom, top, near, far float64) (m Matrix4) {
  m[0]  =  2 * near / (right - left)
  m[2]  =  (right + left) / (right - left)
  m[5]  =  2 * near / (top - bottom)
  m[6]  =  (top + bottom) / (top - bottom)
  m[10] = -(far + near) / (far - near)
  m[11] = -(2 * far * near) / (far - near)
  m[14] = -1
  m[15] =  0
  return m
}

func MakeFrustumFov(fov_y, aspect_ratio, front, back float64) (m Matrix4) {

  tangent := math.Tan(fov_y/2 * DegToRad)   // tangent of half fovY
  height := front * tangent           // half height of near plane
  width := height * aspect_ratio       // half width of near plane

  // params: left, right, bottom, top, near, far
  return MakeFrustum(-width, width, -height, height, front, back)
}

func getCofactor(m0, m1, m2, m3, m4, m5, m6, m7, m8 float64) float64 {
    return m0 * (m4 * m8 - m5 * m7) -
           m1 * (m3 * m8 - m5 * m6) +
           m2 * (m3 * m7 - m4 * m6)
}

/*
func (m *Matrix4) getInverse() Matrix4 {

  cofactor0 = getCofactor(m[5],m[6],m[7], m[9],m[10],m[11], m[13],m[14],m[15])
  cofactor1 = getCofactor(m[4],m[6],m[7], m[8],m[10],m[11], m[12],m[14],m[15])
  cofactor2 = getCofactor(m[4],m[5],m[7], m[8],m[9], m[11], m[12],m[13],m[15])
  cofactor3 = getCofactor(m[4],m[5],m[6], m[8],m[9], m[10], m[12],m[13],m[14])

  // get determinant
  determinant = m[0] * cofactor0 - m[1] * cofactor1 + m[2] * cofactor2 - m[3] * cofactor3
  if math.Abs(determinant) <= 0.00001 {
    return makeMatrix4Identity()
  }

  // get rest of cofactors for adj(M)
  cofactor4 = getCofactor(m[1],m[2],m[3], m[9],m[10],m[11], m[13],m[14],m[15])
  cofactor5 = getCofactor(m[0],m[2],m[3], m[8],m[10],m[11], m[12],m[14],m[15])
  cofactor6 = getCofactor(m[0],m[1],m[3], m[8],m[9], m[11], m[12],m[13],m[15])
  cofactor7 = getCofactor(m[0],m[1],m[2], m[8],m[9], m[10], m[12],m[13],m[14])

  cofactor8 = getCofactor(m[1],m[2],m[3], m[5],m[6], m[7],  m[13],m[14],m[15])
  cofactor9 = getCofactor(m[0],m[2],m[3], m[4],m[6], m[7],  m[12],m[14],m[15])
  cofactor10= getCofactor(m[0],m[1],m[3], m[4],m[5], m[7],  m[12],m[13],m[15])
  cofactor11= getCofactor(m[0],m[1],m[2], m[4],m[5], m[6],  m[12],m[13],m[14])

  cofactor12= getCofactor(m[1],m[2],m[3], m[5],m[6], m[7],  m[9], m[10],m[11])
  cofactor13= getCofactor(m[0],m[2],m[3], m[4],m[6], m[7],  m[8], m[10],m[11])
  cofactor14= getCofactor(m[0],m[1],m[3], m[4],m[5], m[7],  m[8], m[9], m[11])
  cofactor15= getCofactor(m[0],m[1],m[2], m[4],m[5], m[6],  m[8], m[9], m[10])

  // build inverse matrix = adj(M) / det(M)
  // adjugate of M is the transpose of the cofactor matrix of M
  invDeterminant = 1.0 / determinant

  var im Matrix4

  im[0] =  invDeterminant * cofactor0
  im[1] = -invDeterminant * cofactor4
  im[2] =  invDeterminant * cofactor8
  im[3] = -invDeterminant * cofactor12

  im[4] = -invDeterminant * cofactor1
  im[5] =  invDeterminant * cofactor5
  im[6] = -invDeterminant * cofactor9
  im[7] =  invDeterminant * cofactor13

  im[8] =  invDeterminant * cofactor2
  im[9] = -invDeterminant * cofactor6
  im[10]=  invDeterminant * cofactor10
  im[11]= -invDeterminant * cofactor14

  im[12]= -invDeterminant * cofactor3
  im[13]=  invDeterminant * cofactor7
  im[14]= -invDeterminant * cofactor11
  im[15]=  invDeterminant * cofactor15

  return im

}
*/

func (m *Matrix4) toMatrix3() Matrix3 {
  var m3 Matrix3
  m3[0] = m[0]
  m3[1] = m[1]
  m3[2] = m[2]

  m3[3] = m[4]
  m3[4] = m[5]
  m3[5] = m[6]

  m3[6] = m[8]
  m3[7] = m[9]
  m3[8] = m[10]

  return m3
}

func (m *Matrix3) toGLFloat() Matrix3GLFloat {
  var glm Matrix3GLFloat

  for i := range m {
    glm[i] = gl.Float(m[i])
  }

  return glm
}


func (m *Matrix3) getTranspose() Matrix3 {
  var tm Matrix3

  tm[0] = m[0]
  tm[1] = m[3]
  tm[2] = m[6]

  tm[3] = m[1]
  tm[4] = m[4]
  tm[5] = m[7]

  tm[6] = m[2]
  tm[7] = m[5]
  tm[8] = m[8]

  return tm
}

func (m *Matrix3) getInverse() Matrix3 {
  var tm Matrix3
  var determinant, invDeterminant float64
  var tmp [9]float64

  tmp[0] = m[4] * m[8] - m[5] * m[7]
  tmp[1] = m[2] * m[7] - m[1] * m[8]
  tmp[2] = m[1] * m[5] - m[2] * m[4]
  tmp[3] = m[5] * m[6] - m[3] * m[8]
  tmp[4] = m[0] * m[8] - m[2] * m[6]
  tmp[5] = m[2] * m[3] - m[0] * m[5]
  tmp[6] = m[3] * m[7] - m[4] * m[6]
  tmp[7] = m[1] * m[6] - m[0] * m[7]
  tmp[8] = m[0] * m[4] - m[1] * m[3]

  // check determinant if it is 0
  determinant = m[0] * tmp[0] + m[1] * tmp[3] + m[2] * tmp[6]
  //if math.Abs(determinant) <= 0.00001 {
  if math.Abs(determinant) <= math.SmallestNonzeroFloat32 {
    // cannot inverse, make it identity matrix
    return makeMatrix3Identity();
  }

  // divide by the determinant
  invDeterminant = 1.0 / determinant
  tm[0] = invDeterminant * tmp[0]
  tm[1] = invDeterminant * tmp[1]
  tm[2] = invDeterminant * tmp[2]
  tm[3] = invDeterminant * tmp[3]
  tm[4] = invDeterminant * tmp[4]
  tm[5] = invDeterminant * tmp[5]
  tm[6] = invDeterminant * tmp[6]
  tm[7] = invDeterminant * tmp[7]
  tm[8] = invDeterminant * tmp[8]

  return tm
}

func  makeMatrix3Identity() Matrix3 {
  var mat Matrix3
  mat.identity()
  return mat
}

func (m *Matrix3) identity() {
  m[0], m[4], m[8] = 1,1,1
  m[1], m[2], m[3],
  m[5], m[6], m[7] = 0,0,0,0,0,0
}


//from osg
type hVec Quat
type matrixAffineParts struct {
  t Vec4 // translation component
  q Quat // essential rotation
  u Quat // Stretch rotation
  k hVec // Sign of determinant
  f float64
}
type hMatrix Matrix4


func  Matrix4Quat(q Quat) Matrix4 {
  length2 := q.length2()
  if math.Abs(length2) <= math.SmallestNonzeroFloat64 {
    //TODO osg makes the matrix33 to 0...
    return Matrix4Identity()
  }

  var rlength2 float64
  if length2 != 1 {
    rlength2 = 2/length2
  } else {
    rlength2 = 2
  }

  x2 := rlength2*q.X
  y2 := rlength2*q.Y
  z2 := rlength2*q.Z

  xx := q.X * x2
  xy := q.X * y2
  xz := q.X * z2

  yy := q.Y * y2
  yz := q.Y * z2
  zz := q.Z * z2

  wx := q.W * x2
  wy := q.W * y2
  wz := q.W * z2


  var m Matrix4
  //m.identity()
  //already 0
  //m[3], m[7], m[11], m[12], m[13], m[14] = 0,0,0,0,0,0
  m[15] = 1

  m[0] = 1 - (yy + zz)
  m[4] = xy - wz
  m[8] = xz + wy

  m[1] = xy + wz
  m[5] = 1 - (xx + zz)
  m[9] = yz - wx

  m[2] = xz - wy
  m[6] = yz + wx
  m[10] = 1 - (xx + yy)

  return m
}

func Matrix4Translation(v Vec3) (m Matrix4) {
  m[0], m[5], m[10], m[15] = 1,1,1,1
  m[3], m[7], m[11] = v.X, v.Y, v.Z
  return m

}

package ridley

import (
  "math"
  "fmt"
)

type Vec3 struct {
  X float64
  Y float64
  Z float64
}

func (v Vec3) length() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

type Vec4 struct {
  X float64
  Y float64
  Z float64
  W float64
}

func Vec3Add(p1 Vec3, p2 Vec3) (p3 Vec3) {
  p3.X = p1.X + p2.X
  p3.Y = p1.Y + p2.Y
  p3.Z = p1.Z + p2.Z
  return p3
}

func Vec3Sub(p1 Vec3, p2 Vec3) (p3 Vec3) {
  p3.X = p1.X - p2.X
  p3.Y = p1.Y - p2.Y
  p3.Z = p1.Z - p2.Z
  return p3
}

func Vec3Mul(p1 Vec3, s float64) (p2 Vec3) {
  p2.X = p1.X*s
  p2.Y = p1.Y*s
  p2.Z = p1.Z*s
  return p2
}


func Vec3Dot(p1 Vec3, p2 Vec3) float64 {
  return p1.X*p2.X + p1.Y*p2.Y + p1.Z*p2.Z
}

func (v Vec4) length() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

func (v Vec4) length2() float64 {
  return v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W
}

type Quat Vec4

func (v Quat) length2() float64 {
  return v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W
}

const epsilon = 0.0000001

func QuatIdentity() Quat {
  return Quat{0,0,0,1}
}

func (q *Quat) identity() {
  q.X = 0
  q.Y = 0
  q.Z = 0
  q.W = 1
}


func QuatAngleAxis(angle float64, axis Vec3) Quat {
  length := axis.length()
  if length < epsilon {
    return QuatIdentity()
  }

  inverse_norm := 1/length
  cos_half_angle := math.Cos(0.5*angle)
  sin_half_angle := math.Sin(0.5*angle)

  return Quat{ axis.X * sin_half_angle * inverse_norm,
               axis.Y * sin_half_angle * inverse_norm,
               axis.Z * sin_half_angle * inverse_norm,
               cos_half_angle}
}


func LaunchRay(start Vec3, direction Vec3, length float32, objects []*Object) (hitted []*Object, positions []Vec3) {

  for _, o := range objects {
    fmt.Println("object : ", o.Position)
    if o.Box == nil {
      continue
    }

  }

  return hitted, positions
}

type Line struct {
  point Vec3
  direction Vec3
}

type Ray struct {
  start Vec3
  direction Vec3
  length float64
}

type Plane struct {
  point Vec3
  normal Vec3
}

/*
func (plane *Plane) Equation() (a, b, c, d float64) {
  a = plane.normal.X
  b = plane.normal.Y
  c = plane.normal.Z
  d = -Vec3Dot(plane.normal, plane.point)

  return a, b, c, d
}
*/

func IntersectionRayPlane(ray Ray, plane Plane) (b bool, v Vec3) {
  
  dn := Vec3Dot(ray.direction, plane.normal)
    
  if dn == 0 {
    b = false
  } else {
    d := -Vec3Dot(plane.normal, plane.point)
    p0n := -Vec3Dot(ray.start, plane.normal)
    t :=  ( d - p0n) / dn
    b = true
    v = Vec3Add(ray.start, Vec3Mul(ray.direction, t))
  }

  return b, v
}


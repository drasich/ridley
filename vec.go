package ridley

import (
  "math"
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





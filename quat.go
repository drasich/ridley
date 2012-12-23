package ridley

import (
  "math"
  //"fmt"
)

type Quat Vec4

func (q *Quat) length2() float64 {
  return q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W
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

func (q *Quat) conjugate() (r Quat) {
  r.X = -q.X
  r.Y = -q.Y
  r.Z = -q.Z
  r.W = q.W
  return r
}

func (q *Quat) RotateVec3(v Vec3) (r Vec3) {
  qvec := Vec3{q.X, q.Y, q.Z}
  uv := Vec3Cross(qvec, v)
  uuv := Vec3Cross(qvec, uv)
  uv = Vec3Mul(uv, 2* q.W)
  uuv = Vec3Mul(uuv, 2)
  r = Vec3Add(v, uv)
  r = Vec3Add(r, uuv)
  return r
}


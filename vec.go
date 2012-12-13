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
  Start Vec3
  Direction Vec3 //and length
}

type Plane struct {
  Point Vec3
  Normal Vec3
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
  
  dn := Vec3Dot(ray.Direction, plane.Normal)
    
  if dn == 0 {
    b = false
  } else {
    d := -Vec3Dot(plane.Normal, plane.Point)
    p0n := Vec3Dot(ray.Start, plane.Normal)
    t :=  ( d - p0n) / dn
    if ( t > 0 && t <= 1) {
      b = true
      v = Vec3Add(ray.Start, Vec3Mul(ray.Direction, t))
    } else {
      b = false
    }
  }

  return b, v
}

type AABox struct {
  Min Vec3
  Max Vec3
}


func IntersectionRayAABox(ray Ray, box AABox) (isHit bool, isInside bool, position Vec3, normal Vec3) {

  inside := true

  var xt, xn float64
  if ray.Start.X < box.Min.X {
    xt = box.Min.X - ray.Start.X
    if xt > ray.Direction.X {
    fmt.Println("return 00")
      return false, false, position, normal
    }
    xt /= ray.Direction.X
    inside = false
    xn = -1
  } else if (ray.Start.X > box.Max.X) {
    xt = box.Max.X - ray.Start.X
    if xt < ray.Direction.X {
    fmt.Println("return 01")
      return false, false, position, normal
    }
    xt /= ray.Direction.X
    inside = false
    xn = 1
  } else {
    xt = -1
  }

  var yt, yn float64
  if ray.Start.Y < box.Min.Y {
    yt = box.Min.Y - ray.Start.Y
    if yt > ray.Direction.Y {
    fmt.Println("return 02")
      return false, false, position, normal
    }
    yt /= ray.Direction.Y
    inside = false
    yn = -1
  } else if (ray.Start.Y > box.Max.Y) {
    yt = box.Max.Y - ray.Start.Y
    if yt < ray.Direction.Y {
    fmt.Println("return 03")
      return false, false, position, normal
    }
    yt /= ray.Direction.Y
    inside = false
    yn = 1
  } else {
    yt = -1
  }

  var zt, zn float64
  if ray.Start.Z < box.Min.Z {
    zt = box.Min.Z - ray.Start.Z
    if zt > ray.Direction.Z {
    fmt.Println("return 04")
      return false, false, position, normal
    }
    zt /= ray.Direction.Z
    inside = false
    zn = -1
  } else if (ray.Start.Z > box.Max.Z) {
    zt = box.Max.Z - ray.Start.Z
    if zt < ray.Direction.Z {
    fmt.Println("return 05")
      return false, false, position, normal
    }
    zt /= ray.Direction.Z
    inside = false
    zn = 1
  } else {
    zt = -1
  }

  if inside {
    fmt.Println("return 06")
    return true, true, position, normal
  }

  which := 0
  t := xt
  if yt > t {
    which = 1
    t = yt
  }
  if zt > t {
    which = 2
    t = zt
  }

  switch which {
  case 0: // yz plane
    y := ray.Start.Y + ray.Direction.Y*t
    if y < box.Min.Y || y > box.Max.Y { return false, false, position, normal }
    z := ray.Start.Z + ray.Direction.Z*t
    if z < box.Min.Z || z > box.Max.Z { return false, false, position, normal }

    normal.X = xn
  case 1: //xz plane
    x := ray.Start.X + ray.Direction.X*t
    if x < box.Min.X || x > box.Max.X { return false, false, position, normal }
    z := ray.Start.Z + ray.Direction.Z*t
    if z < box.Min.Z || z > box.Max.Z { return false, false, position, normal }

    normal.Y = yn
  case 2:
    x := ray.Start.X + ray.Direction.X*t
    if x < box.Min.X || x > box.Max.X { return false, false, position, normal }
    y := ray.Start.Y + ray.Direction.Y*t
    if y < box.Min.Y || y > box.Max.Y { return false, false, position, normal }

    normal.Y = zn
  }
 
  //return isHit, isInside, position , normal
  return true, false, position , normal
}

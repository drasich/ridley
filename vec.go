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

func Vec3Add(v1 Vec3, v2 Vec3) (v3 Vec3) {
  v3.X = v1.X + v2.X
  v3.Y = v1.Y + v2.Y
  v3.Z = v1.Z + v2.Z
  return v3
}

func Vec3Sub(v1 Vec3, v2 Vec3) (v3 Vec3) {
  v3.X = v1.X - v2.X
  v3.Y = v1.Y - v2.Y
  v3.Z = v1.Z - v2.Z
  return v3
}

func Vec3Mul(v1 Vec3, s float64) (v2 Vec3) {
  v2.X = v1.X*s
  v2.Y = v1.Y*s
  v2.Z = v1.Z*s
  return v2
}

func Vec3Cross(v1 Vec3, v2 Vec3) (v3 Vec3) {
  v3.X = v1.Y * v2.Z - v1.Z*v2.Y
  v3.Y = v1.Z * v2.X - v1.X*v2.Z
  v3.Z = v1.X * v2.Y - v1.Y*v2.X
  return v3
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
      return false, false, position, normal
    }
    xt /= ray.Direction.X
    inside = false
    xn = -1
  } else if (ray.Start.X > box.Max.X) {
    xt = box.Max.X - ray.Start.X
    if xt < ray.Direction.X {
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
      return false, false, position, normal
    }
    yt /= ray.Direction.Y
    inside = false
    yn = -1
  } else if (ray.Start.Y > box.Max.Y) {
    yt = box.Max.Y - ray.Start.Y
    if yt < ray.Direction.Y {
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
      return false, false, position, normal
    }
    zt /= ray.Direction.Z
    inside = false
    zn = -1
  } else if (ray.Start.Z > box.Max.Z) {
    zt = box.Max.Z - ray.Start.Z
    if zt < ray.Direction.Z {
      return false, false, position, normal
    }
    zt /= ray.Direction.Z
    inside = false
    zn = 1
  } else {
    zt = -1
  }

  if inside {
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

  position = Vec3Add(ray.Start, Vec3Mul(ray.Direction,t))
 
  return true, false, position , normal
}

func IntersectionRayObject(ray Ray, o *Object) (
  hit bool, inside bool, position Vec3, normal Vec3) {

  var box AABox

  if o.Box == nil {
    fmt.Println("no box return")
    return false, false, position, normal
  }

  box = o.Box.box

  fmt.Println("ray at first", ray)

  //transform the ray in box/object coord
  var newray Ray
  start :=  Vec3Sub(ray.Start, o.Position)
  iq := o.Orientation.conjugate()
  start = iq.RotateVec3(start)

  dir := iq.RotateVec3(ray.Direction)

  newray.Start = start
  newray.Direction = dir
  fmt.Println("newray", newray)
  fmt.Println("box", box)

  hit, inside, position, normal = IntersectionRayAABox(newray, box)

  //transform back
  position = o.Orientation.RotateVec3(position)
  position = Vec3Add(position, o.Position)

  normal = o.Orientation.RotateVec3(normal)

  return hit, inside, position, normal
}



package back

import (
	"math"
)

type Sphere struct {
	BaseObject
	r float64 // radius
}

func NewSphere(color int, radius float64) *Sphere {
   s := &Sphere{}

   s.color = color
   s.r = radius

   return s
}

func (s *Sphere) Intersection(baseline Coord, eye *Eye) {
   var a, b, c, r1, r2, delta float64

   a = baseline.x * baseline.x + baseline.y * baseline.y + baseline.z * baseline.z
   b = 2 * (eye.x * baseline.x + eye.y * baseline.y + eye.z * baseline.z)
   c = eye.x * eye.x + eye.y * eye.y + eye.z * eye.z - (s.r * s.r)
   delta = (b *b) - (4 * a *c);

   if delta < 0 {
   		s.k = 0
   		return
   }

   sqrtDelta := math.Sqrt(delta)

   r1 = (-b - sqrtDelta) / (2 * a)
   r2 = (-b + sqrtDelta) / (2 * a)

   if r1 > r2 && r2 > 0 {
   		s.k = r2
   } else {
   		s.k = r1
   }

   return
}

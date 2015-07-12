package back

import (
	"math"
   "container/list"
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

/* v = baseline
   s = obj
*/

// int   sphere_light(t_light *light)
// {
//   t_pt      l;
//   t_pt      n;
//   double k;

//   k = obj->k;
//   baseline->x -= obj->x;
//   baseline->y -= obj->y;
//   baseline->z -= obj->z;
//   n.x = (double) eye->x + baseline->x * k;
//   n.y = (double) eye->y + baseline->y * k;
//   n.z = (double) eye->z + baseline->z * k;
//   l.x = (double) spot->x - n.x;
//   l.y = (double) spot->y - n.y;
//   l.z = (double) spot->z - n.z;
//   return (calc_light(&l, &n, obj->color));
// }

func (s *Sphere) Light(baseline Coord, eye *Eye, lights *list.List) (color int) {
   var l, n Coord

   baseline.x -= s.x
   baseline.y -= s.y
   baseline.z -= s.z

   n.x = eye.x + baseline.x * s.k
   n.y = eye.y + baseline.y * s.k
   n.z = eye.x + baseline.z * s.k

   spot := lights.Front().Value.(*Light)
   l.x = spot.x - n.x
   l.y = spot.y - n.y
   l.z = spot.z - n.z

   return light(&l, &n, s.color)
}
package back

import (
	"log"
	"container/list"
	"math"
)

type Light struct {
	Coord
	color int
}

func NewLight(x float64, y float64, z float64, color int) (l *Light) {
	l = &Light{}

	l.x = x
	l.y = y
	l.z = z
	l.color = color
	return
}

// int			choose_light(t_light *light)
// {
//   int			i;
//   extern g_light	gl_light[];
//   int			color;
//   t_pt			v;

//   if (light->s == NULL)
//     return (0);

//   v = *(light->v);
//   rotate(light->v, light->s->rx, light->s->ry, light->s->rz);
//   i = 0;
//   while (i < NBINTERDEF)
//     {
//       if (light && light->s && light->s->id == gl_light[i].id)
// 	color = gl_light[i].funct(light);
//       i++;
//     }
//   *(light->v) = v;
//   return (color);
// }

func setLight(baseline Coord, eye *Eye, lights *list.List, obj Object) (color int) {

	// If there is no object, return dark
	if obj == nil {
		return 0
	}

	rotate(&baseline, obj.RX(), obj.RY(), obj.RZ())
	return obj.Light(baseline, eye, lights)
}


// int		calc_light(t_pt *l, t_pt *n, int color)
// {
//   double	cosa;
//   u_col		u_color;

//   cosa = calc_cosa(n, l);
//   if (cosa < 0.000001 || cosa > 1)
//     return (0);
//   u_color.col = color;
//   u_color.cha[0] *= cosa;
//   u_color.cha[1] *= cosa;
//   u_color.cha[2] *= cosa;
//   u_color.cha[3] *= cosa;
//   return (u_color.col);
// }

func light(l *Coord, n *Coord, color int) int {
	var cosinusa float64
	var res int

	cosinusa = cosa(n, l)
	if cosinusa < 0.000001 || cosinusa > 1 {
		log.Printf("cosinusa not valid : %v\n", cosinusa)
		return 0
	}

	tmp := make([]int8, 4)
	tmp[0] = int8(color)
	tmp[1] = int8(color >> 2)
	tmp[2] = int8(color >> 4)
	tmp[3] = int8(color >> 6)

	tmp[0] = int8(float64(tmp[0]) * cosinusa)
	tmp[1] = int8(float64(tmp[1]) * cosinusa)
	tmp[2] = int8(float64(tmp[2]) * cosinusa)
	tmp[3] = int8(float64(tmp[3]) * cosinusa)

	res = int(tmp[3])
	res = res << 2
	res += int(tmp[2])
	res = res << 2
	res += int(tmp[1])
	res = res << 2
	res += int(tmp[0])
	// log.Printf("res : %v\n", res)
	return color
}

// double		calc_cosa(t_pt *l, t_pt *n)
// {
//   double	normel;
//   double	normen;
//   double	a;

//   normel = sqrt(l->x * l->x + l->y * l->y + l->z * l->z);
//   normen = sqrt(n->x * n->x + n->y * n->y + n->z * n->z);
//   a = l->x * n->x + l->y * n->y + l->z * n->z;
//   return (a / (normel * normen));
// }


func cosa(l *Coord, n *Coord) float64 {
	var normel, normen, a float64

	normel = math.Sqrt(l.x * l.x + l.y * l.y + l.z * l.z)
	normen = math.Sqrt(n.x * n.x + n.y * n.y + n.z * n.z)
	a = l.x * n.x + l.y * n.y + l.z * n.z
	return (a / (normel * normen))
}


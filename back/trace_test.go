package back

import (
	"testing"
	"container/list"
)

func BenchmarkRay(b *testing.B) {
	objects := list.New()
	objects.PushBack(NewSphere(255, 35))

	eye := NewEye(-170, -50, 0)

	lights := list.New()
	lights.PushBack(Light{})

	for i := 0; i < b.N; i++ {
		Ray(150, 150, 300, 300, objects, eye, lights)
	}
}

func BenchmarkSphereP300(b *testing.B) {
	objects := list.New()
	objects.PushBack(NewSphere(255, 35))

	eye := NewEye(-170, -50, 0)

	lights := list.New()
	lights.PushBack(Light{})

	for i := 0; i < b.N; i++ {
		traceParallel(300, 300, objects, eye, lights)
	}	
}

func BenchmarkSphereS300(b *testing.B) {
	objects := list.New()
	objects.PushBack(NewSphere(255, 35))

	eye := NewEye(-170, -50, 0)

	lights := list.New()
	lights.PushBack(Light{})

	for i := 0; i < b.N; i++ {
		traceSequential(300, 300, objects, eye, lights)
	}
}

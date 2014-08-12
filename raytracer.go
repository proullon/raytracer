package main

import (
	"log"
	"os"
	"path"
	"fmt"
	"net/http"
	"encoding/json"
	"container/list"
	"github.com/proullon/raytracer/back"
)

func Raytracer(w http.ResponseWriter, r *http.Request) {
	objects := list.New()
	objects.PushBack(back.NewSphere(255, 35))

	eye := back.NewEye(-170, -50, 0)

	lights := list.New()
	lights.PushBack(back.Light{})

	image := back.Trace(300, 300, objects, eye, lights)
    data, _ := json.Marshal(image)
    fmt.Fprintf(w, "%s", string(data))
}

func main() {
	fmt.Println("Raytracer 0.1")

    static_directory, err := os.Getwd()
    if err != nil {
        log.Fatal("Cannot get working directory")
    }

    static_directory = path.Join(static_directory, "static")

	http.HandleFunc("/api/raytracer", Raytracer)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(static_directory))))
    err = http.ListenAndServe("127.0.0.1:8080", nil)
    if err != nil {
        log.Fatal("Cannot start http server : ", err)
    }
}
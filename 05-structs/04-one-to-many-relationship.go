package main

import "fmt"

type Course struct {
	Title  string
	Videos []Video
}

type Video struct {
	Title  string
	Course Course
}

func main() {
	video1 := Video{
		Title: "Intorducción",
	}

	video2 := Video{
		Title: "Instalación",
	}

	course1 := Course{
		Title:  "Curso de Go",
		Videos: []Video{video1, video2},
	}

	video1.Course = course1
	video2.Course = course1

	// Courses
	fmt.Println("Cursos disponibles")
	fmt.Println(course1)

	// Videos
	fmt.Printf("Capítulos de %s", course1.Title)
	for _, video := range course1.Videos {
		fmt.Println(video.Title)
	}

	// inverse relationship
	fmt.Printf("El video %s pertenece al curso %s", video1.Title, video1.Course.Title)
	fmt.Printf("El video %s pertenece al curso %s", video2.Title, video2.Course.Title)
}

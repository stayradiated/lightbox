package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Job struct {
	Source   string
	Filepath string
}

type Image struct {
	ID     int
	Source string
}

type ImageData struct {
	Shows    []Image
	Fanart   []Image
	Seasons  []Image
	Episodes []Image
}

func main() {

	// list all show posters
	// list all season images
	// list all episode images

	db := new(DB)
	db.Init()
	defer db.Close()

	images := db.GetAllImages()

	jobs := make(chan Job)
	done := make(chan bool)

	for w := 1; w <= 5; w++ {
		go worker(jobs, done)
	}

	// foreman("shows", images.Shows, jobs)
	// foreman("seasons", images.Seasons, jobs)
	// foreman("episodes", images.Episodes, jobs)
	foreman("fanart", images.Fanart, jobs)
	close(jobs)

	for i := 1; i <= 5; i++ {
		<-done
	}
}

func foreman(folder string, images []Image, c chan Job) {
	os.MkdirAll(folder, 0755)
	for _, image := range images {
		c <- Job{
			Source:   image.Source,
			Filepath: filepath.Join(folder, fmt.Sprintf("%d.jpg", image.ID)),
		}
	}
}

func worker(jobs <-chan Job, done chan bool) {
	for job := range jobs {
		if err := download(job.Source, job.Filepath); err != nil {
			log.Println(err)
			continue
		}
	}
	close(done)
}

func download(source string, filepath string) error {
	r, err := http.Get(source)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, r.Body); err != nil {
		return err
	}

	return nil
}

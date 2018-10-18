package main

import (
	"log"
	"os"

	"github.com/mikkyang/id3-go"
)

func main() {
	if len(os.Args) < 2 {
		panic("Missing file argument")
	}
	file := os.Args[1]
	mp3File, err := id3.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	popFrame := mp3File.Frame("POPM")
	log.Println(mp3File)
	log.Println("test")
	//println("").(*v2.UnsynchTextFrame)
	log.Println(popFrame)
}

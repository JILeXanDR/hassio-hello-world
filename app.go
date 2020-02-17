package main

import (
	"github.com/machinebox/sdk-go/facebox"
	"log"
)

func app() {
	faceboxClient := facebox.New("http://localhost:8080")

	billie := getBillie()

	faces, err := faceboxClient.Check(billie)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, face := range faces {
		log.Printf("%+v", face)
		if !face.Matched {
			if err := faceboxClient.TeachURL(getBillieURL(), getBillieURL().Path, "Billie Eilish"); err != nil {
				log.Printf("can't teach: %+v", err)
			}
		}
	}
}

package main

import (
	"image"
	"io/ioutil"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	files, err := ioutil.ReadDir("/Users/maks/go/src/github.com/supperdoggy/taro-pizda/taro-cards/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if len(file.Name()) > 4 && file.Name()[len(file.Name())-4:] == ".png" {
			fio, err := os.Open("/Users/maks/go/src/github.com/supperdoggy/taro-pizda/taro-cards/" + file.Name())
			if err != nil {
				log.Println(err)
				return
			}
			ip, str, err := image.Decode(fio)
			if err != nil {
				log.Println(str, err)
			}
			inverted := imaging.FlipV(ip)
			imaging.Save(inverted, "/Users/maks/go/src/github.com/supperdoggy/taro-pizda/taro-cards/"+file.Name()[:len(file.Name())-4]+"_reversed.png")
		}
	}
}

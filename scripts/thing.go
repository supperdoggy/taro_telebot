package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"strings"
//)
//
//func main() {
//	files, err := ioutil.ReadDir("/Users/maks/go/src/github.com/supperdoggy/taro-pizda/taro-cards/")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, file := range files {
//		if len(file.Name()) > 4 && file.Name()[len(file.Name())-4:] == ".png" {
//			i := strings.Index(file.Name(), "_")
//			name := file.Name()[:len(file.Name())-4]
//			name = name[i+1:]
//			fmt.Printf("\"%s\": \"\",\n", name)
//		}
//	}
//}

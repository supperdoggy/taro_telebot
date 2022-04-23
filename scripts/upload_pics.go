package main

import (
	"github.com/supperdoggy/taro-pizda/structs"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	logger, _ := zap.NewDevelopment()
	session, err := mgo.Dial("")
	if err != nil {
		logger.Fatal("error dialing db", zap.Error(err))
	}
	pics := session.DB("taro_bot_ebat").C("taro_pics")

	files, err := ioutil.ReadDir("/Users/maks/go/src/github.com/supperdoggy/taro-pizda/taro-cards/")
	if err != nil {
		logger.Fatal("error reading dir", zap.Error(err))
	}

	for _, file := range files {
		if len(file.Name()) > 4 && file.Name()[len(file.Name())-4:] == ".png" {
			fio, err := ioutil.ReadFile("/Users/maks/go/src/github.com/supperdoggy/taro-pizda/taro-cards/" + file.Name())
			if err != nil {
				log.Println(err)
				return
			}

			i := strings.Index(file.Name(), "_")
			name := file.Name()[:len(file.Name())-4]
			name = name[i+1:]

			p := structs.TaroPic{
				ID:      name,
				Data:    fio,
				Created: time.Now(),
			}

			err = pics.Insert(p)
			if err != nil {
				logger.Fatal("error inserting pic", zap.Error(err), zap.Any("pic", p))
			}

		}
	}
}

package main
//
//import (
//	"encoding/json"
//	"github.com/supperdoggy/taro-pizda/structs"
//	"go.uber.org/zap"
//	"gopkg.in/mgo.v2"
//	"io/ioutil"
//	"time"
//)
//
//func main() {
//	logger, _ := zap.NewDevelopment()
//	session, err := mgo.Dial("")
//	if err != nil {
//		logger.Fatal("error dialing db", zap.Error(err))
//	}
//	coll := session.DB("taro_bot_ebat").C("taro_ru_loc")
//	raw, err := ioutil.ReadFile("/Users/maks/go/src/github.com/supperdoggy/taro-pizda/day/ru_loc.json")
//	if err != nil {
//		logger.Fatal("error reading file with ru_loc", zap.Error(err))
//	}
//
//	data := map[string]string{}
//
//	err = json.Unmarshal(raw, &data)
//	if err != nil {
//		logger.Fatal("error unmarshalling data from json", zap.Error(err))
//	}
//
//	for k, v := range data {
//		loc := structs.TaroLoc{
//			ID: k,
//			Value: v,
//			Created: time.Now(),
//		}
//		logger.Info("data", zap.Any("loc", loc))
//		err = coll.Insert(loc)
//		if err != nil {
//			logger.Error("error inserting loc", zap.Any("loc", loc), zap.Error(err))
//		}
//	}
//
//}

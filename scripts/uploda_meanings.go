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
//	advice := session.DB("taro_bot_ebat").C("taro_advice_ru")
//	warning := session.DB("taro_bot_ebat").C("taro_warning_ru")
//	raw, err := ioutil.ReadFile("/Users/maks/go/src/github.com/supperdoggy/taro-pizda/day/meaning.json")
//	if err != nil {
//		logger.Fatal("error reading file with ru_loc", zap.Error(err))
//	}
//
//	data := map[string]map[string]string{}
//
//	err = json.Unmarshal(raw, &data)
//	if err != nil {
//		logger.Fatal("error unmarshalling data from json", zap.Error(err))
//	}
//
//	for k, v := range data["advice"] {
//		meaning := structs.TaroMeaning{
//			ID:      k,
//			Value:   v,
//			Created: time.Now(),
//		}
//		logger.Info("ADVICE meaning", zap.Any("meaning", meaning))
//		err = advice.Insert(meaning)
//		if err != nil {
//			logger.Error("ADVICE error uploading meaning", zap.Error(err))
//		}
//	}
//
//	for k, v := range data["warning"] {
//		meaning := structs.TaroMeaning{
//			ID:      k,
//			Value:   v,
//			Created: time.Now(),
//		}
//		logger.Info("ADVICE meaning", zap.Any("meaning", meaning))
//		err = warning.Insert(meaning)
//		if err != nil {
//			logger.Error("ADVICE error uploading meaning", zap.Error(err))
//		}
//	}
//
//}

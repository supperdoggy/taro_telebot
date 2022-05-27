package localization

import (
	"fmt"
	"sync"
)

type Localization struct {
	m   map[string]string
	mut sync.Mutex
}

var (
	loc = Localization{
		m: map[string]string{
			"start_message": "",
			"daily_taro":    "*Ваша карта дня*: %s\n\n*Совет дня*: %s\n\n*Предостережение дня*: %s",
		},
	}
)

func GetLoc(key string, args ...interface{}) string {
	loc.mut.Lock()
	defer loc.mut.Unlock()
	val := loc.m[key]
	if len(args) != 0 {
		val = fmt.Sprintf(val, args...)
	}
	return val
}

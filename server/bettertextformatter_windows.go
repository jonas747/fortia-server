package main

// Based on logrus.TextFormatter, a few changes: better colors, instead of current runtime it shows the actual time
import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"sort"
	"strings"
	"time"
)

type TextFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool
}

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var serialized []byte

	levelText := strings.ToUpper(entry.Data["level"].(string))[0:4]

	timeStr := entry.Data["time"].(string)
	when, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeStr)
	finalTime := when.Format("2006-01-02 15:04:05.9")

	serialized = append(serialized, []byte(fmt.Sprintf("[%s] %s %-45s ", finalTime, levelText, entry.Data["msg"]))...)

	keys := make([]string, 0)
	for k, _ := range entry.Data {
		if k != "level" && k != "time" && k != "msg" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	first := true
	for _, k := range keys {
		v := entry.Data[k]
		if first {
			first = false
		} else {
			serialized = append(serialized, ' ')
		}
		serialized = append(serialized, []byte(fmt.Sprintf("%s=%v", k, v))...)
	}

	return append(serialized, '\n'), nil
}

func (f *TextFormatter) AppendKeyValue(serialized []byte, key, value interface{}) []byte {
	if _, ok := value.(string); ok {
		return append(serialized, []byte(fmt.Sprintf("%v='%v' ", key, value))...)
	} else {
		return append(serialized, []byte(fmt.Sprintf("%v=%v ", key, value))...)
	}
}

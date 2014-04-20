package main

// Based on logrus.TextFormatter, a few changes: better colors, instead of current runtime it shows the actual time
import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/burke/ttyutils"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	col_nocolor = "\x1b[0m"
	col_black   = "\x1b[30m"
	col_red     = "\x1b[31m"
	col_green   = "\x1b[32m"
	col_yellow  = "\x1b[33m"
	col_blue    = "\x1b[34m"
	col_magenta = "\x1b[35m"
	col_cyan    = "\x1b[36m"
	col_white   = "\x1b[37m"

	col_bold_black   = col_black[:4] + ";1m"
	col_bold_red     = col_red[:4] + ";1m"
	col_bold_green   = col_green[:4] + ";1m"
	col_bold_yellow  = col_yellow[:4] + ";1m"
	col_bold_blue    = col_blue[:4] + ";1m"
	col_bold_magenta = col_magenta[:4] + ";1m"
	col_bold_cyan    = col_cyan[:4] + ";1m"
	col_bold_white   = col_white[:4] + ";1m"
)

type TextFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool
}

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var serialized []byte

	if f.ForceColors || ttyutils.IsTerminal(os.Stdout.Fd()) {
		levelText := strings.ToUpper(entry.Data["level"].(string))[0:4]

		levelColor := col_bold_blue

		switch entry.Data["level"] {
		case "warning":
			levelColor = col_bold_yellow
		case "error", "fatal", "panic":
			levelColor = col_bold_red
		case "debug":
			levelColor = col_magenta
		}

		timeStr := entry.Data["time"].(string)
		when, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeStr)
		finalTime := when.Format("2006-01-02 15:04:05.9")

		serialized = append(serialized, []byte(fmt.Sprintf("%s[%s] %s\x1b[0m %-45s ", levelColor, finalTime, levelText, entry.Data["msg"]))...)

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
			serialized = append(serialized, []byte(fmt.Sprintf("\x1b[%dm%s\x1b[0m=%v", levelColor, k, v))...)
		}
	} else {
		serialized = f.AppendKeyValue(serialized, "time", entry.Data["time"].(string))
		serialized = f.AppendKeyValue(serialized, "level", entry.Data["level"].(string))
		serialized = f.AppendKeyValue(serialized, "msg", entry.Data["msg"].(string))

		for key, value := range entry.Data {
			if key != "time" && key != "level" && key != "msg" {
				serialized = f.AppendKeyValue(serialized, key, value)
			}
		}
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

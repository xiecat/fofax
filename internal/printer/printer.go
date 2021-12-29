package printer

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/fatih/color"
)

const (
	ClearLn = "\r\x1b[2K"
	INFO    = "[INFO]"
	ERROR   = "[ERRO]"
	SUCCESS = "[SUCC]"
	DEBUG   = "[DEBUG]"
)

var (
	strC   = color.New(color.FgHiCyan).SprintfFunc()
	outPut = color.Error
)

type errInfo struct {
	fileName string
	line     int
	message  string
}

func HandlerLine(i interface{}) string {
	var message string
	switch i.(type) {
	case error:
		message = i.(error).Error()
	case string:
		message = i.(string)
	}
	_, file, line, _ := runtime.Caller(1)
	e := &errInfo{
		fileName: filepath.Base(file),
		line:     line,
		message:  message,
	}
	return e.Error()
}

func (e *errInfo) Error() string {
	return "[" + e.fileName + ":" + strconv.Itoa(e.line) + "] " + e.message
}

func Debugf(format string, a ...interface{}) {
	debug := color.New(color.FgYellow).SprintfFunc()
	prefix := debug(DEBUG)
	str := strC(fmt.Sprintf(format, a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
}

func Debug(a ...interface{}) {
	debug := color.New(color.FgYellow).SprintfFunc()
	prefix := debug(DEBUG)
	str := strC(fmt.Sprint(a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
}

func Infof(format string, a ...interface{}) {
	info := color.New(color.FgHiWhite).SprintfFunc()
	prefix := info(INFO)
	str := strC(fmt.Sprintf(format, a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
}

func Info(a ...interface{}) {
	info := color.New(color.FgHiWhite).SprintfFunc()
	prefix := info(INFO)
	str := strC(fmt.Sprint(a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
}

func Successf(format string, a ...interface{}) {
	succ := color.New(color.FgGreen).SprintfFunc()
	prefix := succ(SUCCESS)
	str := strC(fmt.Sprintf(format, a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
}
func Success(a ...interface{}) {
	succ := color.New(color.FgGreen).SprintfFunc()
	prefix := succ(SUCCESS)
	str := strC(fmt.Sprint(a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
}

func Errorf(format string, a ...interface{}) {
	err := color.New(color.FgRed).SprintfFunc()
	prefix := err(ERROR)
	str := strC(fmt.Sprintf(format, a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
}
func Error(a ...interface{}) {
	err := color.New(color.FgRed).SprintfFunc()
	prefix := err(ERROR)
	str := strC(fmt.Sprint(a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
}

func Fatalf(format string, a ...interface{}) {
	err := color.New(color.FgRed).SprintfFunc()
	prefix := err(ERROR)
	str := strC(fmt.Sprintf(format, a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
	os.Exit(-1)
}
func Fatal(a ...interface{}) {
	err := color.New(color.FgRed).SprintfFunc()
	prefix := err(ERROR)
	str := strC(fmt.Sprint(a...))
	fmt.Fprintln(outPut, formatPrint(prefix), str)
	os.Exit(-1)
}

func formatPrint(prefix string) string {
	format := time.Now().Format("2006/01/02 15:04:05") + " " + prefix
	return format

}

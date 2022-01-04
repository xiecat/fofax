package fofa

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/xiecat/fofax/internal/printer"
)

type FofaLine struct {
	Host     string `json:"id" csv:"host"`
	LastTime string `json:"lastupdatetime" csv:"host"`
}

func ImportFile(filename string, size int, ffi bool, res *sync.Map) {
	printer.Successf("Fetch Data From %s: [%d/%d]", filename, size, getLine(filename))
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)
	if strings.HasSuffix(filename, ".csv") {
		err = readCSV(rd, size, ffi, res)
		if err != nil {
			printer.Fatalf("file: %s,err: %s\n", filename, err.Error())
		}
	} else if strings.HasSuffix(filename, ".json") {
		readJson(rd, size, ffi, res)
	} else {
		printer.Fatal("the file suffix can only be csv or json")
	}
}

func readCSV(rd *bufio.Reader, size int, ffi bool, res *sync.Map) error {
	lnum := 0
	for {
		if lnum > size {
			return nil
		}
		line, _, err := rd.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
		fileds := strings.Split(string(line), ",")
		if len(fileds) != 2 {
			printer.Errorf("line: %d data: %s parser err\n", lnum+1, string(line))
			lnum++
			continue
		}
		if lnum == 0 {
			//检查开头
			if !strings.EqualFold("lastupdatetime", strings.TrimSpace(fileds[1])) {
				return errors.New("csv format err")
			}
			lnum++
			continue
		}
		if ffi {
			res.LoadOrStore(fixHost(fileds[0]), nil)
		} else {
			res.LoadOrStore(trimSchema(fileds[0]), nil)
		}
		lnum++
	}
	return nil
}

func readJson(rd *bufio.Reader, size int, ffi bool, res *sync.Map) {
	lnum := 1
	for {
		if lnum > size {
			return
		}
		line, _, err := rd.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
		fdata := FofaLine{}
		err = json.Unmarshal(line, &fdata)
		if err != nil {
			printer.Errorf("line: %d data: %s parser err\n", lnum, string(line))
			lnum++
			continue
		}
		if fdata.Host == "" {
			printer.Errorf("line: %d data: %s parser err\n", lnum, string(line))
			lnum++
			continue
		}
		if ffi {
			res.LoadOrStore(fixHost(fdata.Host), nil)
		} else {
			res.LoadOrStore(trimSchema(fdata.Host), nil)
		}
		lnum++
	}
}

func fixHost(furl string) string {
	if strings.HasPrefix(furl, "https://") || strings.HasPrefix(furl, "http://") {
		return furl
	}
	return fmt.Sprintf("http://%s", furl)
}

func trimSchema(furl string) string {
	if strings.HasPrefix(furl, "https://") {
		return strings.TrimPrefix(furl, "https://")
	}
	if strings.HasPrefix(furl, "http://") {
		return strings.TrimPrefix(furl, "http://")
	}
	return furl
}

func getLine(filename string) int {
	line := 1
	if strings.HasSuffix(filename, ".csv") {
		line = 0
	}
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)

	for {
		_, _, err := rd.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
		line++
	}
	return line
}

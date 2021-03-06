package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/zuiwuchang/go-3322/version"

	"gitlab.com/king011/king-go/timer"
)

const (
	// ConfigureFile 配置檔案 名稱
	ConfigureFile = "go-3322.jsonnet"
)

func main() {
	var v, h bool
	flag.BoolVar(&v, "v", false, "display version")
	flag.BoolVar(&h, "h", false, "display help")
	flag.Parse()
	if h {
		flag.PrintDefaults()
	} else if v {
		fmt.Println(runtime.GOOS, runtime.GOARCH)
		fmt.Println(version.Tag)
		fmt.Println(version.Commit)
		fmt.Println(version.Date)
	} else {
		cnf, e := initConfigure()
		if e != nil {
			log.Fatalln(e)
		}

		duration, e := timer.ToDuration(cnf.Timer)
		if e != nil {
			log.Fatalln(e)
		}

		for {
			doWork()
			time.Sleep(duration)
		}
	}
}

func doWork() {
	defer func() {
		fmt.Println()
	}()
	log.Println("***	update begin	***")

	cnf := getConfigure()
	var url string
	if cnf.A != "" {
		url = fmt.Sprintf(`http://members.3322.net/dyndns/update?hostname=%v&myip=%v`,
			cnf.Host,
			cnf.A,
		)
	} else {
		url = fmt.Sprintf(`http://members.3322.net/dyndns/update?hostname=%v`,
			cnf.Host,
		)
	}
	log.Println(url)
	req, e := http.NewRequest(http.MethodGet, url, nil)
	if e != nil {
		log.Println(e)
		return
	}

	str := base64.StdEncoding.EncodeToString([]byte(cnf.Name))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", str))

	var c http.Client
	res, e := c.Do(req)
	if e != nil {
		log.Println(e)
		return
	}
	b, e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Println(e)
		return
	}

	str = string(b)
	if strings.HasPrefix(str, "good") || strings.HasPrefix(str, "nochg") {
		log.Println("***	update success	***")
	} else {
		log.Println(str)
	}
}

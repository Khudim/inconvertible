package main

import (
	"bytes"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"gopkg.in/yaml.v2"
	"image/png"
	"io/ioutil"
	"log"
)

type Template struct {
	Name   string  `json:"name"`
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Button string  `json:"button"`
	Pause  float64 `json:"pause"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
}

type AppConfig struct {
	ProfitCount int         `yaml:"profitCount"`
	BotId       string      `yaml:"botId"`
	ClientId    string      `yaml:"clientId"`
	Trinkets    []*Template `yaml:"trinkets"`
	Buybacks    []*Template `yaml:"buybacks"`
	Logins      []*Template `yaml:"logins"`
	Merchants   []*Template `yaml:"merchants"`
	Logouts     []*Template `yaml:"logouts"`
	Options     []*Template `yaml:"options"`
	BuybackTabs []*Template `yaml:"buybackTabs"`
	GoldInBags  []*Template `yaml:"goldInBags"`
}

var gold map[string][]byte

func main() {
	var appConf *AppConfig

	if file, err := ioutil.ReadFile("./props.yaml"); err == nil {
		if err := yaml.Unmarshal(file, &appConf); err != nil {
			log.Println(err)
			panic(err)
		}
	} else {
		panic(err)
	}

	go startTelegramClient(appConf)

	gold = make(map[string][]byte)

	log.Println("Погнали ёпта")
	robotgo.MicroSleep(3000)

	go exitListener()

	for {
		for i := 0; i < appConf.ProfitCount; i++ {
			// login
			for _, v := range appConf.Logins {
				click(v)
			}
			// click merchant
			for _, v := range appConf.Merchants {
				click(v)
			}
			// sell trinkets
			for _, v := range appConf.Trinkets {
				click(v)
			}
			// options + logout
			for i, v := range appConf.Options {
				click(v)
				click(appConf.Logouts[i])
			}
		}
		// login
		for _, v := range appConf.Logins {
			click(v)
		}
		// click merchant
		for _, v := range appConf.Merchants {
			click(v)
		}
		// sell trinkets
		for _, v := range appConf.Trinkets {
			click(v)
		}
		// buyback button
		for _, v := range appConf.BuybackTabs {
			click(v)
		}
		// buyback trinkets
		for _, v := range appConf.Buybacks {
			click(v)
		}

		for _, v := range appConf.GoldInBags {
			body := makeScreenshot(v.X, v.Y, v.Width, v.Height)
			gold[v.Name] = body
		}

		// options + logout
		for i, v := range appConf.Options {
			click(v)
			click(appConf.Logouts[i])
		}
	}
}

func click(template *Template) {
	robotgo.MoveMouseSmooth(template.X, template.Y, 0.5, 0.5)
	robotgo.MouseClick(template.Button)
	robotgo.MicroSleep(template.Pause)
}

func exitListener() {
	for {
		if ok := robotgo.AddEvent("f4"); ok {
			panic("exit")
		}
	}
}

func makeScreenshot(x, y, width, height int) []byte {
	img, _ := screenshot.Capture(x, y, width, height)
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return []byte{}
	}
	return buf.Bytes()
}

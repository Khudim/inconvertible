package main

import (
	"github.com/go-vgo/robotgo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Template struct {
	Name   string  `json:"name"`
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Button string  `json:"button"`
	Pause  float64 `json:"pause"`
}

type AppConfig struct {
	ProfitCount int         `yaml:"profitCount"`
	Trinkets    []*Template `yaml:"trinkets"`
	Buybacks    []*Template `yaml:"buybacks"`
	Logins      []*Template `yaml:"logins"`
	Merchants   []*Template `yaml:"merchants"`
	Logouts     []*Template `yaml:"logouts"`
	BuybackTabs []*Template `yaml:"buybackTabs"`
}

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

	log.Println("Погнали ёпта")
	robotgo.MicroSleep(2000)

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
			// logout
			for _, v := range appConf.Logouts {
				click(v)
				tripleEsc()
				click(v)
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
		// logout
		for _, v := range appConf.Logouts {
			click(v)
			tripleEsc()
			click(v)
		}
		robotgo.MicroSleep(1000)
	}
}

func tripleEsc() {
	robotgo.KeyTap("escape")
	robotgo.MicroSleep(50)
	robotgo.KeyTap("escape")
	robotgo.MicroSleep(50)
	robotgo.KeyTap("escape")
	robotgo.MicroSleep(50)
}

func click(template *Template) {
	robotgo.MoveMouseSmooth(template.X, template.Y, 0.9, 0.9)
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

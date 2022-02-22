package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

type Settings struct {
	Name  string
	Value string
}

var prefs = [12]Settings{
	Settings{"browser.newtabpage.activity-stream.feeds.telemetry", "false"},
	Settings{"browser.newtabpage.activity-stream.telemetry", "false"},
	Settings{"browser.ping-centre.telemetry", "false"},
	Settings{"toolkit.telemetry.archive.enabled", "false"},
	Settings{"toolkit.telemetry.bhrPing.enabled", "false"},
	Settings{"toolkit.telemetry.firstShutdownPing.enabled", "false"},
	Settings{"toolkit.telemetry.newProfilePing.enabled", "false"},
	Settings{"toolkit.telemetry.reportingpolicy.firstRun", "false"},
	Settings{"toolkit.telemetry.server", "\"\""},
	Settings{"toolkit.telemetry.shutdownPingSender.enabled", "false"},
	Settings{"toolkit.telemetry.unified", "false"},
	Settings{"toolkit.telemetry.updatePing.enabled", "false"}}

type User struct {
	Uid      string
	Gid      string
	Username string
	Name     string
	HomeDir  string
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func processingUserDir(path string) {
	files, err := ioutil.ReadDir(path + "/.mozilla/firefox")
	check(err)
	for _, f := range files {
		if f.IsDir() && strings.Index(f.Name(), ".") == 8 {
			if _, err := os.Stat(path + "/.mozilla/firefox/" + f.Name() + "/prefs.js"); err == nil {
				processingPrefsFile(path + "/.mozilla/firefox/" + f.Name() + "/prefs.js")
			}
		}
	}
}

func existSetting(line string) bool {
	for _, seting := range prefs {
		if strings.Index(line, seting.Name) == 11 {
			return true
		}
	}
	return false
}

func addSetting(textPrefs string) string {
	for _, setting := range prefs {
		textPrefs += "user_pref(\"" + setting.Name + "\", " + setting.Value + ");\n"
	}
	return textPrefs
}

func processingPrefsFile(path string) {
	dat, err := ioutil.ReadFile(path)
	var resultPrefs string
	check(err)
	for _, line := range strings.Split(string(dat), "\n") {
		if !existSetting(line) && len(line) != 0 {
			resultPrefs += line + "\n"
		}
	}
	resultPrefs = addSetting(resultPrefs)
	resultDat := []byte(resultPrefs)
	err = ioutil.WriteFile(path, resultDat, 0600)
	check(err)
}

func main() {
	currentUser, err := user.Current()
	check(err)
	if currentUser.Uid == "0" {
		files, err := ioutil.ReadDir("/home")
		check(err)
		for _, f := range files {
			if f.IsDir() {
				processingUserDir("/home/" + f.Name())
			}
		}
	} else {
		processingUserDir(currentUser.HomeDir)
	}
}

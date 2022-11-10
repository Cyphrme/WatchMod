package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"time"

	"github.com/DisposaBoy/JsonConfigReader"
	"github.com/fsnotify/fsnotify"
)

//

// flags holds setting settable by flag.
// Example with flags:
//
//	go run watch.go -config=custom.json5 -daemon=false
type flags struct {
	configPath string
	daemon     bool
}

type Config struct { // Config options setable by config file.
	WatchCommand   map[string]string
	ExcludeFiles   []string
	ExcludeStrings []string
	RunCmdOnStart  bool
}

var fc flags
var c Config

var regexes []*regexp.Regexp

func main() {
	parseFlags()
	run()
}

func parseFlags() {
	flag.StringVar(&fc.configPath, "config", "watch.json5", "Path for the watch config.")
	flag.BoolVar(&fc.daemon, "daemon", true, "Run as daemon.  If false, runs command and shuts down.")
	flag.Parse()
	// fmt.Printf("Flag config path: %s, %v\n", fc.configPath, fc.daemon)
}

func run() {
	parseConfig(&c)

	sort.Strings(c.ExcludeFiles) // must be sorted for search
	setStringRegexes()
	// log.Printf("Config: %+v\n", c)
	// log.Printf("WatchCommand: %+v\n", c.WatchCommand)
	// log.Printf("Exclude Files: %+v\n", c.ExcludeFiles)
	// log.Printf("Exclude Strings: %+v\n", c.ExcludeStrings)

	if !fc.daemon {
		fmt.Println("Flag `daemon` set to false.  Running commands in config and exiting.")
		c.RunCmdOnStart = true
	}

	var expanded = make(map[string]string)
	for k, v := range c.WatchCommand {
		var err error
		// For windows slashes
		k, err = filepath.Abs(os.ExpandEnv(k))
		if err != nil {
			panic(err)
		}

		// For windows slashes
		v, err = filepath.Abs(os.ExpandEnv(v))
		if err != nil {
			panic(err)
		}

		expanded[k] = v
		if c.RunCmdOnStart {
			runCmd(v)
		}
	}

	if !fc.daemon {
		return
	}

	done := make(chan bool)
	for dir, cmd := range expanded {
		go Watch(dir, cmd)
	}
	<-done
}

// Watch is for each dir/cmd to watch/run.
func Watch(dir, cmd string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Rename == fsnotify.Rename {
					// TODO, if a write just happened, sometimes rename is also triggered.
					// Should only trigger once.

					log.Printf("File changed: %+s, event: %s\n", event.Name, event.Op.String())
					fileName := filepath.Base(event.Name)
					if Excluded(fileName) {
						continue
					}
					runCmd(cmd)
				}
			case err := <-watcher.Errors:
				log.Fatal(err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done setting up watch for " + dir)
	<-done
}

func runCmd(cmd string) {
	log.Printf("Start run %q\n", cmd)
	start := time.Now()

	c := exec.Command(cmd)
	if stdoutStderr, err := c.CombinedOutput(); err != nil {
		log.Printf("Watch Error: %s; On cmd: %s; Error: \n%s\n", err, cmd, stdoutStderr)
	}

	elapsed := time.Since(start)
	log.Printf("End   run %q in %s\n", cmd, elapsed)
}

func parseConfig(i interface{}) {
	expand := os.ExpandEnv(fc.configPath)

	// For windows slashes
	expand, err := filepath.Abs(expand)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(expand)
	if err != nil {
		panic(err)
	}

	// wrap reader before passing it to the json decoder for comment stripping
	r := JsonConfigReader.New(file)

	decoder := json.NewDecoder(r)
	err = decoder.Decode(i)
	if err != nil {
		panic(err)
	}
}

// Essentially just a search function.
func excludeByFileName(fileName string) (excluded bool) {
	i := sort.SearchStrings(c.ExcludeFiles, fileName) // binary search
	if len(c.ExcludeFiles) == i {
		return false
	}
	if c.ExcludeFiles[i] != fileName { // Is the element the thing?  If not, it's new.
		return false
	}
	return true
}

func setStringRegexes() {
	for _, ext := range c.ExcludeStrings {
		escaped := regexp.QuoteMeta(ext)
		reg := regexp.MustCompile(escaped)
		regexes = append(regexes, reg)
	}
}

func matchExcludeString(filename string) (excluded bool) {
	for _, reg := range regexes {
		matched := reg.Match([]byte(filename))
		if matched {
			log.Printf("Input Matched, FileName and ExludeExt: [%s], [%s]\n", filename, reg)
			return true
		}
	}
	return false
}

func Excluded(fileName string) (excluded bool) {
	excluded = excludeByFileName(fileName)
	if excluded {
		log.Printf("Excluded by file name: %s.  Doing nothing.  \n", fileName)
		return true
	}

	excluded = matchExcludeString(fileName)
	if excluded {
		log.Printf("Excluded by string: %s.  Doing nothing.  \n", fileName)
		return true
	}

	return false
}

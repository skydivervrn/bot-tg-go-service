package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	appOutput     = os.Getenv("APP_OUTPUT_FILE")
	serviceOutput = os.Getenv("CMD_OUTPUT_FILE")
	buffer        []byte
	bp            = &buffer
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		if fileExists(appOutput) {
			err := readFile()
			tmp := strings.Split(string(*bp), "\n")
			if err != nil {
				log.Printf("Error is: %s", err)
				writeFile(strings.Join([]string{tmp[0], "READ File error", err.Error()}, "\n"))
				os.Remove(appOutput)
				continue
			}
			log.Printf("DEBUG: %v", tmp[1])
			out, errw := runCommand(parseStringList(tmp[1]))
			if errw != nil {
				log.Printf("Command finished with error: %v", errw)
			}
			writeFile(strings.TrimSuffix(strings.Join([]string{tmp[0], out, ""}, "\n"), "\n"))
			os.Remove(appOutput)
		}
	}
}

func readFile() error {
	log.Printf("Start reading command from file...")
	var err error
	*bp, err = ioutil.ReadFile(appOutput)
	if err != nil {
		log.Println("Unable to open file:", err)
		return err
	}
	log.Printf("File read")
	return err
}

func writeFile(data string) bool {
	log.Printf("Starting write file and waiting for it to finish...")
	d1 := []byte(data)
	err := ioutil.WriteFile(serviceOutput, d1, 0644)
	if err != nil {
		log.Println("Unable to write file:", err)
		return true
	}
	log.Printf("File written")
	return false
}

func runCommand(command string, args []string) (string, error) {
	log.Printf("Running command and waiting for it to finish...")
	out, err := exec.Command(command, args...).Output()
	log.Printf("Command is %s", exec.Command(command, args...).String())
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return string(out), err
	}
	log.Printf("Command execution finished")
	return string(out), err
}

func parseStringList(str string) (string, []string) {
	a := strings.Split(str, " ")
	return a[0], a[1:]
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

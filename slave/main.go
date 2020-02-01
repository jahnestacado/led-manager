package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	ledPath         = "/sys/class/leds/orangepi"
	redLedCommand   = "status/brightness"
	greenLedCommand = "pwr/brightness"
	colors          = []string{"green", "red"}
	defaultPort     = "3333"
)

func resetState(color string) error {
	for _, c := range colors {
		if color != c {
			err := executeCommand(c, "0")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// echo 1 > /sys/class/leds/orangepi:red:status/brightness
// echo 0 > /sys/class/leds/orangepi:green:pwr/brightness

func executeCommand(color, state string) error {
	ledCommand := redLedCommand
	if color == "green" {
		ledCommand = greenLedCommand
	}

	ledPath := fmt.Sprintf("%s:%s:%s", ledPath, color, ledCommand)
	ledFile, err := os.Create(ledPath)
	if err != nil {
		fmt.Println("os.Create", err)
		return err
	}
	defer ledFile.Close()

	cmd := exec.Command("echo", state)
	cmd.Stdout = ledFile

	err = cmd.Run()
	if err != nil {
		fmt.Println("cmd.Run", err)
		return err
	}

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path[1:], "/")
	color := params[0]
	state := params[1]

	err := resetState(color)
	if err != nil {
		fmt.Println("resetState", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = executeCommand(color, state)
	r.Close = true
	r.Header.Set("Connection", "close")
	if err != nil {
		fmt.Println("executeCommand", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "OK")
}

func main() {
	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = defaultPort
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

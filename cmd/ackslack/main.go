package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/BurntSushi/toml"
	"github.com/jessevdk/go-flags"
)

const APP_NAME = "ackslack"

type Options struct {
	ConfigPath string `short:"c" long:"config" description:"path to config file"`
	Args struct {
		Command []string `positional-arg-name:"COMMAND"`
	} `positional-args:"yes" required:"yes"`
}

type Config struct {
	WebhookUrl string `toml:"webhook_url"`
}

type SlackAttachment struct {
	Color string `json:"color"`
	Text string `json:"text"`
}

type SlackMessage struct {
	Text string `json:"text"`
	Attachments []SlackAttachment `json:"attachments"`
}

func CreateOkMessage(text string) SlackMessage {
	return SlackMessage{
		Text: "",
		Attachments: []SlackAttachment{
			{Color: "good", Text: text},
		},
	}
}

func CreateBadMessage(text string) SlackMessage {
	return SlackMessage{
		Text: "",
		Attachments: []SlackAttachment{
			{Color: "danger", Text: text},
		},
	}
}

func CheckFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func main() {

	var options Options
	_, err := flags.Parse(&options)

	if err != nil {
		fmt.Println("[!] invalid options")
		os.Exit(1)
	}

	if len(options.Args.Command) == 0 {
		fmt.Println("[!] command not specified")
		os.Exit(1)
	}

	config_path := options.ConfigPath

	if config_path == "" {

		paths := []string{}
		app_dir := APP_NAME
		app_toml := fmt.Sprintf("%s.toml", APP_NAME)
		dotapp_toml := fmt.Sprintf(".%s.toml", APP_NAME)

		if conf_dir, err := os.UserConfigDir(); err == nil {
			paths = append(paths, filepath.Join(conf_dir, app_dir, app_toml))
			paths = append(paths, filepath.Join(conf_dir, app_dir, "config.toml"))
		}
		if home_dir, err := os.UserHomeDir(); err == nil {
			paths = append(paths, filepath.Join(home_dir, dotapp_toml))
		}
		paths = append(paths, app_toml)
		for _, path := range paths {
			if CheckFileExists(path) {
				config_path = path
				break
			}
		}
		if config_path == "" {
			fmt.Println("[!] config file not found")
			os.Exit(1)
		}
	}

	var config Config
	_, err = toml.DecodeFile(config_path, &config)

	if err != nil {
		fmt.Println("[!] cannot decode config file")
		os.Exit(1)
	}

	cmd := options.Args.Command
	proc := exec.Command(cmd[0], cmd[1:]...)
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	proc.Stdin = os.Stdin
	err = proc.Run()

	var msg SlackMessage

	if err != nil {
		text := fmt.Sprintf("command `%s` failed", proc.String())
		fmt.Printf("[*] %s\n", text)
		msg = CreateBadMessage(text)
	} else {
		text := fmt.Sprintf("command `%s` succeeded", proc.String())
		fmt.Printf("[*] %s\n", text)
		msg = CreateOkMessage(text)
	}

	encoded, err := json.Marshal(msg)

	if err != nil {
		fmt.Println("[!] cannot execute command")
		os.Exit(1)
	}

	resp, err := http.Post(config.WebhookUrl, "application/json", bytes.NewBuffer(encoded))

	if err != nil {
		fmt.Printf("[!] post failed (err = %s)\n", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
}

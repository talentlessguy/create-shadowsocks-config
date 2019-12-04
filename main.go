package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

type option struct {
	name string
	desc string
}

func greet() {
	color.New(color.FgBlue).Add(color.Bold).Println("ShadowSocks config generator")
}

func setOption(option option) (n int, err error) {
	return color.New(color.FgGreen).Printf("%s: ", option.desc)
}

func choosePath() string {
	color.New(color.FgHiMagenta).Print("Config path (default: ~/ss-config.json): ")
	path := ReadInput()

	if path == "" {
		path, _ = tilde.Expand("~")
	}

	return path
}

func getConfig() map[string]interface{} {

	config := map[string]interface{}{
		"server":      "127.0.0.1",
		"server_port": 8080,
		"local_port":  1080,
		"password":    "Hello World",
		"timeout":     600,
		"method":      "aes-256-cfb",
	}

	options := []option{
		option{name: "server", desc: "Server address (IP / Host)"},
		option{name: "server_port", desc: "Server port"},
		option{name: "local_port", desc: "Local port"},
		option{name: "password", desc: "Password"},
		option{name: "timeout", desc: "Request timeout (default: 600)"},
		option{name: "method", desc: "Encryption method (default: aes-256-cfb)"},
	}

	// TODO: Make a loop but Idk how to properly set field in struct by a variable

	for _, opt := range options {
		setOption(opt)

		text := ReadInput()

		if text == "" {
			continue
		}

		config[opt.name] = text
	}

	return config
}

func main() {

	app := &cli.App{
		Name:    "create-shadowsocks-config",
		Version: "0.0.1",
		Usage:   "Generate config for shadowsocks by v1rtl (v1rtl.site)",
		Action: func(c *cli.Context) error {
			greet()

			path := choosePath()

			conf := getConfig()

			color.New(color.FgBlue).Println("\nGenerated config:")

			fmt.Println(PrettyPrint(conf))

			SaveMapAsJSONToFile(conf, path)

			return nil
		},
	}

	app.Run(os.Args)
}

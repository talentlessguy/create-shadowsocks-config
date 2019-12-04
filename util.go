package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	tilde "gopkg.in/mattes/go-expand-tilde.v1"
)

// ReadInput - read input from stdin
func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	text = strings.Replace(text, "\n", "", -1)

	return text
}

// PrettyPrint - pretty-print map
func PrettyPrint(m map[string]interface{}) string {
	b, err := json.MarshalIndent(m, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

// SaveMapAsJSONToFile - save map to JSON file
func SaveMapAsJSONToFile(m map[string]interface{}, path string) {

	b := []byte(PrettyPrint(m))

	path, _ = tilde.Expand(path + "/ss-config.json")

	err := ioutil.WriteFile(path, b, 0644)

	if err != nil {
		panic(err)
	}
}

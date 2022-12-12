package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetCurrentDirEnv() string {
	if _, err := os.Stat(".env"); err != nil {
		log.Fatal("Failed to locate the env file.")
	} else {
		file, err := ioutil.ReadFile(".env")
		if err != nil {
			log.Fatal("Failed to open the .env file.")
		}
		return string(file)
	}
	return ""
}

func Parse(str string) []string {
	split := strings.Split(str, "\n")
	arr := []string{}
	for i := 0; i < len(split); i++ {
		x := strings.Split(split[i], "=")
		arr = append(arr, x[0])
	}
	return arr
}

func main() {
	log.Println("Generating file...")
	fileContents := GetCurrentDirEnv()
	if fileContents == "" {
		log.Fatal("Failed to get the .env's file content.")
	}
	parsed := Parse(fileContents)
	err := ioutil.WriteFile(".env.example", []byte(fmt.Sprintf("%v=VALUE\n", strings.Join(parsed, "=VALUE\n"))), 0666)
	if err != nil {
		log.Fatal("Failed to create env example file.")
	}
	log.Println("Created .env.example file.")
}

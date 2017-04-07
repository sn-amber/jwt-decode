package main

import (
	"io/ioutil"
	"os"
	"encoding/base64"
	"strings"
	"encoding/json"
	"fmt"
	"log"
	"flag"
)

var (
	token string
)

func main() {

	flag.StringVar(&token, "token", "", "token to be decoded")
	flag.Parse()

	if token == "" {
		var err error
		token, err = readTokenFromStdin()
		fatalOnErr(err, "cannot read token from stdin")
	}

	// split token into three parts
	segments := strings.Split(token, ".")

	for _, segment := range segments[:len(segments)-1] {

		data, err := base64ToJson(segment)
		fatalOnErr(err, "cannot decode from base64")

		err = prettyPrintJson(data)
		fatalOnErr(err, "cannot pretty-print json")
	}
}

func readTokenFromStdin() (string, error) {
	bt, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

func base64ToJson(str string) (interface{}, error) {

	// decode
	decoded, err := base64.RawStdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	// unmarshal
	var parsed interface{}
	err = json.Unmarshal(decoded, &parsed)
	if err != nil {
		return "", err
	}
	return parsed, nil
}

func prettyPrintJson(data interface{}) error {
	bt, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(bt))
	return nil
}

func fatalOnErr(err error, desc string) {
	if err != nil {
		log.Fatalln(desc, err)
	}
}

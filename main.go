package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	token    string
	datetime bool

	errValWrongType   = errors.New("Value of 'exp' not convertible to int")
	errTokenWrongType = errors.New("Token not of shape map[string]interface{}")
	errKeyNotFound    = errors.New("Key 'exp' not found")
)

func main() {

	flag.StringVar(&token, "token", "", "token to be decoded")
	flag.BoolVar(&datetime, "datetime", false, "whether to print timestamp, instead of unix time for \"exp\"")
	flag.Parse()

	if token == "" {
		var err error
		token, err = readTokenFromStdin()
		fatalOnErr(err, "cannot read token from stdin")
	}

	// split token into three parts
	segments := strings.Split(token, ".")

	for i, segment := range segments[:len(segments)-1] {

		data, err := base64ToJSON(segment, datetime)
		if err != nil {
			if err != errTokenWrongType && err != errKeyNotFound && err != errValWrongType { // external error
				fatalOnErr(err, "cannot decode from base64")
			} else if i == 1 { // exp error
				logrus.Warn("could not extract expiration datetime: ", err)
			}
		}
		err = prettyPrintJSON(data)
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

func base64ToJSON(str string, utcFlag bool) (interface{}, error) {

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
	if utcFlag {
		if p, ok := parsed.(map[string]interface{}); ok {

			found := false
			for key, val := range p {
				if key == "exp" {
					found = true
					if v, ok := val.(float64); ok {
						p[key] = time.Unix(int64(v), 0)
					} else {
						return parsed, errValWrongType
					}
				}
			}

			if !found {
				return parsed, errKeyNotFound
			}
			return p, nil
		}
		return parsed, errTokenWrongType
	}
	return parsed, nil
}

func prettyPrintJSON(data interface{}) error {
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

package main

import (
	"io/ioutil"
	"os"
	"encoding/base64"
	"strings"
	"encoding/json"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	segments := strings.Split(string(input), ".")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")

	for _, segment := range segments[:len(segments)-1] {
		decoded, err := base64.RawStdEncoding.DecodeString(segment)
		if err != nil {
			panic(err)
		}

		var parsed interface{}
		err = json.Unmarshal(decoded, &parsed)
		if err != nil {
			panic(err)
		}

		err = enc.Encode(parsed)
		if err != nil {
			panic(err)
		}
	}
}

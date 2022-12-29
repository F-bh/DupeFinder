package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal("could not read dir")
	}

	lenMap := make(map[int64]string)
	dupeLenMap := make(map[string]string)
	hashMap := make(map[string]string)
	dupeHashMap := make(map[string]string)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			log.Fatal("could not read file info for: " + entry.Name())
		}

		name, ok := lenMap[info.Size()]
		if !ok {
			lenMap[info.Size()] = entry.Name()
		} else {
			dupeLenMap[entry.Name()] = name
		}

		file, err := os.Open("./" + entry.Name())
		if err != nil {
			panic("could not read file: " + entry.Name())
		}

		h := sha256.New()
		if _, err := io.Copy(h, file); err != nil {
			log.Fatal(err)
		}

		//hash
		name, ok = hashMap[hex.EncodeToString(h.Sum(nil))]

		if !ok {
			hashMap[hex.EncodeToString(h.Sum(nil))] = entry.Name()
		} else {
			dupeHashMap[entry.Name()] = name
			continue
		}

		for k, v := range dupeHashMap {
			fmt.Printf("found duplicate hash for: '%s' - '%s'\n", k, v)
		}

		for k, v := range dupeLenMap {
			fmt.Printf("found duplicate length for: '%s' - '%s'\n", k, v)
		}
	}
}

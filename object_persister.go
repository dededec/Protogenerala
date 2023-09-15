package main

import (
	"log"
	"os"
)

func save_object(filename string, s Serializer) error {
	return os.WriteFile(filename, s.Serialize(), 0666)
}

func load_object(filename string, s Serializer) Serializer {
	bs, err := os.ReadFile(filename)

	if err != nil {
		log.Println(err)
		return nil
	}

	return s.Deserialize(bs)
}

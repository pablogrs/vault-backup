package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJson(t *testing.T) {

	filename := "vault.test.json"
	client, err := NewBackup()
	if err != nil {
		log.Fatal(err)
	}

	jsonMap := make(map[string]string)

	jsonMap["deployment/secret1/key1"] = "1234"
	jsonMap["deployment/secret1/key2"] = "abc"
	jsonMap["deployment/secret2/key1"] = "5678"
	jsonMap["deployment/secret2/key2"] = "efg"

	client.secrets = jsonMap
	client.output = "json"
	client.filename = filename

	if err := client.write(); err != nil {
		log.Printf("Could not write file\n")
		log.Printf("%v", err)
		t.Fail()
	}

	mapSecrets, err := client.readJson("vault.test.json")

	if err != nil {
		log.Printf("Could not read Json file\n")
		log.Printf("%v", err)
		t.Fail()
	}

	assert.Equal(t, mapSecrets["deployment/secret1/key1"], jsonMap["deployment/secret1/key1"])

	os.Remove(filename)
}

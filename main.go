package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var desiredPrefixes = []string{"CAT", "42/"}

const numberOfKeys = 5

func main() {
	//Uncomment the function you want to run
	findFirst()
	//findX()
}

func generateKeyPair() (string, string) {
	privateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}
	publicKey := privateKey.PublicKey()
	return privateKey.String(), publicKey.String()
}

func hasDesiredPrefix(publicKey string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(publicKey, prefix) {
			return true
		}
	}
	return false
}

func findFirst() {
	count := 0
	found := false

	for !found {
		privateKey, publicKey := generateKeyPair()
		count++

		if hasDesiredPrefix(publicKey, desiredPrefixes) {
			fmt.Printf("Found a matching key after %d attempts\n", count)
			fmt.Printf("Private key: %s\nPublic key: %s\n", privateKey, publicKey)
			found = true
		}

		if count%10000 == 0 {
			fmt.Printf("Checked %d keypairs\n", count)
		}
	}
}

func findX() {
	count := 0
	foundKeys := 0
	var foundPairs []string

	for foundKeys < numberOfKeys {
		privateKey, publicKey := generateKeyPair()
		count++

		if hasDesiredPrefix(publicKey, desiredPrefixes) {
			foundPairs = append(foundPairs, fmt.Sprintf("Private key: %s\nPublic key: %s\n\n", privateKey, publicKey))
			foundKeys++
			fmt.Printf("Found %d/%d keys\n", foundKeys, numberOfKeys)
		}

		if count%10000 == 0 {
			fmt.Printf("Checked %d keypairs\n", count)
		}
	}

	file, err := os.Create("found.txt")
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
	defer func() { _ = file.Close() }()

	for _, pair := range foundPairs {
		_, err = file.WriteString(pair)
		if err != nil {
			log.Fatalf("Unable to write to file: %v", err)
		}
	}

	fmt.Printf("Found %d keys after %d tries. Keys saved to 'found.txt'\n", foundKeys, count)
}

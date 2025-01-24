package utils

import (
	"encoding/json"
	"os"

	"github.com/gocarina/gocsv"
)

// Title,URL,Username,Password,Notes,OTPAuth
type KeychainEntry struct {
	Title    string `cvs:"Title"`
	URL      string `cvs:"URL"`
	Username string `cvs:"Username"`
	Password string `cvs:"Password"`
	Notes    string `cvs:"Notes"`
	OTPAuth  string `cvs:"OTPAuth"`
}

func ReadKeychainFile(filepath string) ([]*KeychainEntry, error) {
	in, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	entries := []*KeychainEntry{}

	if err := gocsv.UnmarshalFile(in, &entries); err != nil {
		return nil, err
	}

	return entries, nil
}

func ConvertKeychainToJSON(importFilepath string, exportFilepath string) error {
	keychainEntries, err := ReadKeychainFile(importFilepath)

	if err != nil {
		return err
	}

	bitwardenItems := []*BitwardenItem{}

	for _, entry := range keychainEntries {
		item := &BitwardenItem{
			Type: 1,
			Name: entry.Title,
			Login: BitwardenItemLogin{
				Uris: []BitwardenItemLoginUris{
					{
						Uri: entry.URL,
					},
				},
				Username: entry.Username,
				Password: entry.Password,
				OTPAuth:  entry.OTPAuth,
			},
		}

		bitwardenItems = append(bitwardenItems, item)
	}

	file, err := json.MarshalIndent(&BitwardenFile{
		Items: bitwardenItems,
	}, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile(exportFilepath, file, 0644)

	if err != nil {
		return err
	}

	return nil
}

package proxy

import (
	"fmt"
	"os"

	"github.com/mmanjoura/niya-voyage/backend/pkg/amadeus"
)

func GetAmadiusClient() (*amadeus.Amadeus, error) {
	client, err := amadeus.New(
		os.Getenv("AMADEUS_CLIENT_ID"),
		os.Getenv("AMADEUS_CLIENT_SECRET"),
		os.Getenv("AMADEUS_ENV"), // set to "TEST"
	)
	if err != nil {
		fmt.Println("not expected error while creating client", err)
		return nil, err
	}
	return client, nil
}

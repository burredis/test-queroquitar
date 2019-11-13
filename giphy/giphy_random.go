package giphy

import (
	"fmt"
	"os"

	libgiphy "github.com/sanzaru/go-giphy"
)

// GetRandom return random gif
func GetRandom(t string) (string, error) {
	giphy := libgiphy.NewGiphy(os.Getenv("GIPHY_API_KEY"))

	random, err := giphy.GetRandom(t)
	if err != nil {
		fmt.Println("error:", err)
	}

	return random.Data.Image_url, err
}

package helpers

import (
	"os"

	"github.com/mdp/qrterminal/v4"
)

func GenerateQRCode(url string) {
	qrterminal.Generate(url, qrterminal.M, os.Stdout)
}

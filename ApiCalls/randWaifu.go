package ApiCalls

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetRandom() {
	// URL of the image you want to download
	imageUrl := "https://api.devgoldy.xyz/aghpb/v1/random"

	// Make a GET request to the URL
	response, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", response.Status)
		return
	}

	// Create a new file to save the image
	file, err := os.Create("image.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Copy the image data from the response to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Image saved successfully.")
}

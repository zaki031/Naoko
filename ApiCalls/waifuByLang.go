package ApiCalls
import (
	
	"net/http"
	"io/ioutil"
	"strings"
	"math/rand"
	"io"
    "fmt"
    "encoding/json"
	"os"
)

type WaifuBook struct {
    SearchID   string `json:"search_id"`
    Name      string `json:"name"`
    Category  string `json:"category"`
    DateAdded string `json:"date_added"`
}

func WaifuByLang(lang string){
	url := "https://api.devgoldy.xyz/aghpb/v1/search?query="+lang

	// Make a GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

    var WaifuBooks []WaifuBook
    if err := json.Unmarshal([]byte(string(body)), &WaifuBooks); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }
	var searchIDs []string

    for _, WaifuBook := range WaifuBooks {
		searchID := WaifuBook.SearchID
		if err != nil{

			fmt.Println(err)
		  }
		if strings.ToLower(lang) == strings.ToLower(WaifuBook.Category){
			searchIDs = append(searchIDs, searchID)

		}

   
    } 
	
	
	if len(searchIDs) > 0 {
		randomIndex := rand.Intn(len(searchIDs))
		searchByID(searchIDs[randomIndex])

	}





}




func searchByID(id string){
	imageUrl := "https://api.devgoldy.xyz/aghpb/v1/get/id/"+ id

	// Make a GET request to the api url
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

	// create new file to save the image
	file, err := os.Create("image.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// copy the image data from the response to the file created above
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Image saved successfully.")
}

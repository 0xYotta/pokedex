package pokeapi

import "fmt"

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func DisplayAll(locationList RespShallowLocations) {
	fmt.Println(locationList.Count)
	if locationList.Next == nil {
		fmt.Println(locationList.Next)
	} else {
		fmt.Println(*locationList.Next)
	}
	if locationList.Previous == nil {
		fmt.Println(locationList.Previous)
	} else {
		fmt.Println(*locationList.Previous)
	}
	for _, obj := range locationList.Results {
		fmt.Printf("name: %s, url: %s\n", obj.Name, obj.Url)
	}
}

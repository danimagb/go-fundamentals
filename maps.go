package main

import "fmt"

func main() {
	//Maps are reference types and dynamically resizable
	//They do not have order
	//Unsafe for concurrency
	//Cheap to pass
	//We can specify size for maps make(map[<keyType>]<valueType>, size) -> to increase performance on big maps

	leagueTitles := make(map[string]int)
	leagueTitles["Porto"] = 31
	leagueTitles["Benfica"] = 0
	leagueTitles["Sporting"] = 0

	/*recentHead2Head := map[string]int{
		"Porto" : 5,
		"Benfica" : 0
	}*/

	for key, value := range leagueTitles {
		fmt.Printf("\n Key is: %v Value is: %v", key, value)
	}

	fmt.Println(leagueTitles["Porto"])

	leagueTitles["Porto"] = 32

	// Add if the key does not exists
	leagueTitles["Famalicao"] = 1

	fmt.Println(leagueTitles)

	delete(leagueTitles, "Benfica")

	fmt.Println(leagueTitles)
}

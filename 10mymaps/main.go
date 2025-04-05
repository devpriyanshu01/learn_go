package main

import "fmt"

func main() {
	fmt.Println("---------------   Maps in Golang    -----------------")

	//initialize some space for a map ds
	var techs = make(map[string]string)

	techs["JS"] = "JavaScript"
	techs["K8S"] = "Kubernetes"
	techs["PY"] = "Python"
	techs["TS"] = "TypeScript"

	fmt.Println("Printing techs (a map) --> \n", techs)

	delete(techs, "TS") //delete TS key and it's value
	fmt.Println("Printing techs (a map) --> \n", techs)
	

	//loops are interesting in Golang
	for key, value := range techs{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}

package main

import "fmt"

func add(key,value string,medal map[string]string){

	medal[key]=value
}

func update(key,value string,medal map[string]string){

	medal[key]=value
}

func get(key,value string,medal map[string]string)string{

	return medal[key]
}

func getall(medal map[string]string){
	for key, value := range medal{
		fmt.Println("Key:", key, "Value:", value)
	}
}



func main(){

	var medal map[string]string
	medal = make(map[string]string)

	fmt.Println("\n Now Using add function")
	add("Gold","1",medal)
	add("Silver","2",medal)
	add("Platinum","3",medal)

	getall(medal)
	update("Platinum","5",medal)
	fmt.Println("\n Now using update")
	getall(medal)
	add("Platinum","3",medal)
	delete(medal,"Gold",)
	fmt.Println("\n Now using delete")

	getall(medal)

}
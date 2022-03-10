package fungsi

import "fmt"

func Channel() {
	channel2()
}

func channel2() {
	var messages = make(chan string)
	names := []string{"Arum", "Novita"}
	for _, name := range names {
		go func(nam string) {
			var data = fmt.Sprintf("my name is %s", nam)
			messages <- data
		}(name)
	}

	// fmt.Println(messages)
	for i := 0; i < len(names); i++ {
		printName(messages)
	}
}

func printName(namePrint chan string) {
	fmt.Println(<-namePrint)
}

func channel() {
	var messages = make(chan string)

	var sayHello = func(name string) {
		var data = fmt.Sprintf("Hello my name is %s", name)
		messages <- data
	}

	go sayHello("Arum")
	go sayHello("Reza")

	var message1 = <-messages
	fmt.Println(message1)
	var message2 = <-messages
	fmt.Println(message2)
}

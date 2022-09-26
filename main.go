package main

import (
	"fmt"
	"os"
)

const green = "\033[32m"
const blue = "\033[34m"

func main() {
	argsWithProg := os.Args[1:]

	commands := make(map[string]string)
	commands["make:migration"] = "test command"
	commands["list"] = "Lists all available commands"

	if _, ok := commands[argsWithProg[0]]; ok {
		
	}
}

func list(commands map[string]string) {

	logo := `
	#     #                        #                                        
	#     # #    # ######         # #   #####  ##### #  ####    ##   #    # 
	#     # #    # #             #   #  #    #   #   # #       #  #  ##   # 
	#     # #    # #####  ##### #     # #    #   #   #  ####  #    # # #  # 
	 #   #  #    # #            ####### #####    #   #      # ###### #  # # 
	  # #   #    # #            #     # #   #    #   # #    # #    # #   ## 
	   #     ####  ######       #     # #    #   #   #  ####  #    # #    # 
																			
	`
	fmt.Println(blue + logo)
	for key, element := range commands {
		fmt.Println(green + key + " " + blue + element)
	}
}

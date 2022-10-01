package main

import (
	"fmt"
	"os"
)

const green = "\033[32m"
const blue = "\033[34m"
const red = "\033[31m"

type Artisan struct{}

func Commands() map[string]string {
	commands := make(map[string]string)
	commands["make:component"] = "make component"
	commands["list"] = "Lists all available commands"
	commands["make:project"] = "Creates a new vite project; add ts argument to create a typecript vite project"
	commands["make:view"] = "Creates a view file"

	return commands
}

func main() {
	argsWithProg := os.Args[1:]

	artisan := Artisan{}

	switch argsWithProg[0] {
	case "make:component":
		artisan.MakeComponent(argsWithProg[1])
		break
	case "list":
		artisan.List()
	case "make:project":
		artisan.MakeProject(argsWithProg)
	case "make:view":
		artisan.MakeView(argsWithProg)
	default:
		fmt.Println("Unknown command, use list for all available commands")
	}

}

func (a Artisan) List() {
	commands := Commands()
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

func (a Artisan) MakeComponent(name string) {
	fmt.Println("Making component...", name)
	filename := fmt.Sprintf("src/components/%s.vue", name)
	if _, err := os.Stat("src/components/"); os.IsNotExist(err) {
		fmt.Println("Directory does not exist, creating...")
		err := os.MkdirAll("/src/components/", 0755)

		if err != nil {
			fmt.Println("Error creating directory: ", err)
		}
	} else {
		fmt.Println("Directory exists")
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Fprint(file, "<template>\n\t<div>\n\n\t</div>\n</template>\n")
	fmt.Fprint(file, "<script lang=\"ts\">\nimport { defineComponent } from \"vue\";\n", "export default defineComponent({\n\n})", "\n</script>\n")
	fmt.Fprint(file, "<style scoped>\n\n</style>")
	fmt.Println("File ", filename, " created successfully")
}

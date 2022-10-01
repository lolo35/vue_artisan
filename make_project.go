package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func (a Artisan) MakeProject(arguments []string) {
	typescript := false
	project_name := "vue_app"

	//npm create vite@latest my-vue-app -- --template vue

	for key := range arguments {
		if arguments[key] == "ts" {
			typescript = true
		}
		if strings.Contains(arguments[key], "--name=") {
			strs := strings.Split(arguments[key], "--name=")
			project_name = strs[1]
		}
	}

	template := "vue"
	program := "npm"
	arg1 := "create"
	arg2 := "vite@latest"
	// arg3 := "."
	arg4 := "--"
	arg5 := "--template"

	if typescript {
		fmt.Println("Creating vite typescript project...")
		template = "vue-ts"
	}

	cmd := exec.Command(program, arg1, arg2, project_name, arg4, arg5, template)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Print(string(stdout))

	if strings.Contains(string(stdout), "Done") {
		fmt.Println("Vite scaffold created, installing packages")
		install := exec.Command("npm", "install")
		install.Dir = path.Join(project_name)

		_, ierr := install.Output()

		if ierr != nil {
			fmt.Println(ierr.Error())
			return
		}

		// fmt.Print(string(stdout))
		a.AddRouter(project_name, typescript)
		a.MoveArtisan(project_name)
	}
}

func (a Artisan) AddRouter(project_name string, typescript bool) {
	router_question := "Would you like to add vue router to the project? (y/n)"
	var router_answer string
	fmt.Println(router_question)
	fmt.Scan(&router_answer)

	if router_answer == "y" || router_answer == "yes" {
		fmt.Println("Adding router to project")

		installCmd := exec.Command("npm", "install", "vue-router@4")
		installCmd.Dir = path.Join(project_name)

		_, err := installCmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		mkdir := a.CreateDirectory(path.Join(project_name), "routes")

		if !mkdir {
			return
		}

		file := "index.js"

		if typescript {
			file = "index.ts"
		}

		createFile, err := os.Create(path.Join(project_name, "/routes/", file))

		if err != nil {
			fmt.Print(err.Error())
			return
		}

		defer createFile.Close()

		fmt.Fprint(createFile, `import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'
const routes: Array<RouteRecordRaw> = [
	{
		// add routes here
		// path: "/",
		// name: "Home",
		// component: HomeView,
	}
];
		
const router = createRouter({
	history: createWebHashHistory(),
	routes
});
		
export default router;`)

		fmt.Println("Creating views directory...")
		if !a.CreateDirectory(path.Join(project_name), "views") {
			return
		}
	}
}

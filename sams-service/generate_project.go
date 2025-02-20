package main

import (
	"fmt"
	"os"
)

func createDir(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %s", dir)
	}
	return nil
}

func createFile(file string, content string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", file)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write content to file: %s", file)
	}
	return nil
}

func main() {
	root := "my_project"
	dirs := []string{
		root + "/app/controller",
		root + "/app/service",
		root + "/app/model",
		root + "/config",
		root + "/routes",
		root + "/middleware",
		root + "/repository",
		root + "/util",
		root + "/public",
	}

	// Create directories
	for _, dir := range dirs {
		if err := createDir(dir); err != nil {
			fmt.Println(err)
			return
		}
	}

	// Create main.go
	mainContent := `package main

import (
 "github.com/gin-gonic/gin"
 "my_project/routes"
 "my_project/middleware"
)

func main() {
 r := gin.Default()
 r.Use(middleware.Logger())
 r.Use(middleware.Recovery())

 routes.SetupRoutes(r)

 r.Run(":8080")
}
`
	if err := createFile(root+"/main.go", mainContent); err != nil {
		fmt.Println(err)
		return
	}

	// Create routes/routes.go
	routesContent := `package routes

import (
 "github.com/gin-gonic/gin"
 "my_project/app/controller"
)

func SetupRoutes(r *gin.Engine) {
 r.GET("/users", controller.GetUsers)
 r.POST("/users", controller.CreateUser)
}
`
	if err := createFile(root+"/routes/routes.go", routesContent); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Project structure created successfully!")
}

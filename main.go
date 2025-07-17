package main

// imports
import (
	"fmt";
	"log";
	"github.com/natnael-eyuel-dev/Enhanced-Task-Manager-REST-API/data";
	"github.com/natnael-eyuel-dev/Enhanced-Task-Manager-REST-API/router";
)

// entry point of the Enhanced Task Manager REST API application
func main() {
	fmt.Println("Enhanced Task Manager REST API Project")      // print startup message

	// initialize service and controller layers
	taskService, err := data.NewMongoDBTaskManager (      // create persistent task service instance using mongodb go driver 
			"mongodb://localhost:27017",
			"taskdb",
			"tasks",
	)        
	
	if err != nil {
		log.Fatal(err)
	}
	defer taskService.Close()

	router := router.SetupRouter(taskService)	  // initialize the router with all configured routes
	
	router.Run(":8080")                        // start the server on port 8080
	log.Println("Starting server on :8080")
}
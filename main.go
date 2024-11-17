package main

import(
	"github.com/ChanchalS7/practice_api/database"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	//load env file
	loadEnv()

	//load database configuration and connection
	loadDatabase()
	//start server
	serveApplication()
}
func loadEnv(){
	err:=godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase(){
	database.InitDb()
}

func serveApplication(){
	router:=gin.Default()
	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
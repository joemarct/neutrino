package backend

import (
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// Load the sqlite dialect for gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dbPath := path.Join(dir, "sqlite3.db")
	db, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&todoModel{})
}

// RunServer runs the backend API server
func RunServer(port int) {
	router := gin.Default()
	router.Use(cors.Default())

	gin.SetMode(gin.ReleaseMode)

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
	}
	router.Run(":" + strconv.Itoa(port))
}

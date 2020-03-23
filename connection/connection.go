package connection

import (
	// "fmt"
	"database/sql"
    // "html/template"
    _ "github.com/lib/pq"
    "os"
    "github.com/joho/godotenv"
    "log"
    "regexp"
    "fmt"
)
func init () {
    const projectDirName = "news-app"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
    cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

    err := godotenv.Load(string(rootPath) + `/.env`)

    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func Connect() (*sql.DB) {
    user     := os.Getenv("DB_USER")
    host     := os.Getenv("DB_HOST")
    pass     := os.Getenv("DB_PASSWORD")
    dbname   := os.Getenv("DB_NAME")
    port     := os.Getenv("DB_PORT")

    var err error
    dbinfo := fmt.Sprintf("host=%s port=%s user=%s "+
            "password=%s dbname=%s sslmode=disable",
            host, port, user, pass, dbname)
    db, err := sql.Open("postgres", dbinfo)

    if err != nil {
        panic(err)
    }

    if err = db.Ping(); err != nil {
        panic(err)
    }

	return db
}
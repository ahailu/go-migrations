package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"
    "net/http"
    "database/sql"

    _ "github.com/lib/pq"
    "github.com/DavidHuie/gomigrate"
    "os"
    "math"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "migrations! migrations!")
}


func main() {
    db, _ := initDB()

    // carry on and run migrations
    runMigrations(db)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func initDB() (*sql.DB, error) {
    driver := "postgres"
    dbUri := os.Getenv("DATABASE_URL")
    db, err := sql.Open("postgres", dbUri)
    if err != nil {
        log.Println(err)
    }
    fmt.Println(driver, dbUri)

    if err != nil {
        log.Println("DB connection error: ", err.Error())
    }

    return db, err
}

func runMigrations(db *sql.DB) {
    currDelay := 8
    tick := time.NewTicker(time.Duration(currDelay) * time.Millisecond)
    rand.Seed(time.Now().UTC().UnixNano())
    for {
        select {
        case <-tick.C:
            if err := db.Ping(); err == nil {
                migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./db/migrations")
                if err := migrator.Migrate(); err == nil {
                    log.Println("Connected to and migrated database")
                    tick.Stop()
                    return
                }
                log.Println("Failed migration")
            }

            currDelay *= 2
            if currDelay > 250 {
                currDelay = 250
            }
            currDelay = int(math.Abs(float64(currDelay+rand.Intn(20)-10))) + 1
            tick.Stop()
            tick = time.NewTicker(time.Duration(currDelay) * time.Millisecond)
            log.Println("Retrying migration with delay of %dms", currDelay)
        }
    }
}

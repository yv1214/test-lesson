package main

import (
    "fmt"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "strings"
)

func queryString(w http.ResponseWriter, r *http.Request) {
  fmt.Println("GET params were:", r.URL.Query())
  queryMap := r.URL.Query()
  if len(queryMap) > 0 {
    for key, element := range queryMap {
      fmt.Fprintf(w, "%v: %v\n", key, element)
    }

    // Create the database handle, confirm driver is present
    db, _ := sql.Open("mysql", "root:lolol@tcp/test")
    defer db.Close()

    var err error

    username := strings.Join(queryMap["username"], "")
    password := strings.Join(queryMap["password"], "")

    sqlStatement := `INSERT INTO test (username, password, created) VALUES (?,?,NOW()) `
    _, err = db.Exec(sqlStatement, username, password)
    fmt.Println("inserting: ", username, password)
    if err != nil {
      panic(err)
    }
  }
}

func main() {
    // Create the database handle, confirm driver is present
    db, _ := sql.Open("mysql", "root:lolol@tcp/test")
    defer db.Close()

    // Connect and check the server version
    var version string
    db.QueryRow("SELECT VERSION()").Scan(&version)
    fmt.Println("Connected to:", version)

    http.HandleFunc("/lol", queryString)
    http.ListenAndServe("127.0.0.1:8090", nil)

}

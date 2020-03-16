package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5433
    user     = "tester"
    password = "tester123"
    dbname   = "tester"
)

func main() {
    db := ConnectToDB()
    defer db.Close()
    stat := db.Stats()

    fmt.Printf("db stat waitDuration: %v\n", stat.WaitDuration)
    fmt.Printf("db stat WaitCount: %v\n", stat.WaitCount)

    log.Println("test begin...")

    DoSomethingWithDBNotClose(db)
    DoSomethingWithDBNotClose(db)

    log.Println("test is done...")
}

func ConnectToDB() *sql.DB{
    // DB config.
    psqlCfg := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // connect DB.
    db, err := sql.Open("postgres", psqlCfg)
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }
    db.SetMaxOpenConns(1);
    db.SetMaxIdleConns(1);

    log.Println("Successfully connected to DB!")

    return db
}

func DoSomethingWithDBNotClose(db *sql.DB) {
    log.Println("DoSomethingWithDBNotclose is called")

    stat := db.Stats()
    fmt.Printf("db stat OpenConnections: %v\n", stat.OpenConnections)
    fmt.Printf("db stat InUse: %v\n", stat.InUse)
    fmt.Printf("db stat Idle: %v\n", stat.Idle)

    sqlStatement :=
        `select * from test`

    _, err := db.Query(sqlStatement)
    if err != nil {
        fmt.Println(err)
    }

    log.Println("DoSomethingWithDBNotclose end...")
    // I'm not gonna close connection.
    // for rows.Next() {}
    // rows.Close()

    /*
when I don't use rows.Next(), next statement doesn't end for 1 hour...
maybe it will not end forever...

2020/03/16 17:23:38 Successfully connected to DB!
db stat waitDuration: 0s
db stat WaitCount: 0
2020/03/16 17:23:38 test begin...
2020/03/16 17:23:38 DoSomethingWithDBNotclose is called
db stat OpenConnections: 1
db stat InUse: 0
db stat Idle: 1
2020/03/16 17:23:38 DoSomethingWithDBNotclose end...
2020/03/16 17:23:38 DoSomethingWithDBNotclose is called
db stat OpenConnections: 1
db stat InUse: 1
db stat Idle: 0
*/

}

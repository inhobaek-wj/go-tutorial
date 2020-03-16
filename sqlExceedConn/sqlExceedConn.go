package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

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

    go func() {
        DoSomethingWithDBNotClose(db, "first")
    }()

    go func() {
        DoSomethingWithDBNotClose(db, "second")
    }()

    time.Sleep(time.Second * 6)
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

func DoSomethingWithDBNotClose(db *sql.DB, text string) {
    log.Println("DoSomethingWithDBNotclose is called")

    stat := db.Stats()
    fmt.Printf("db stat OpenConnections: %v\n", stat.OpenConnections)
    fmt.Printf("db stat InUse: %v\n", stat.InUse)
    fmt.Printf("db stat Idle: %v\n", stat.Idle)

    sqlStatement :=
        `select * from test`

    rows, err := db.Query(sqlStatement)
    if err != nil {
        fmt.Println(err)
    }

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
    time.Sleep(time.Second * 2)
    for rows.Next() {
        fmt.Println("row is scaned from " + text)
    }

    log.Println("DoSomethingWithDBNotclose end...")

    /*
return connection after 2 seconds, then second statement is excuted.

    2020/03/16 18:46:36 Successfully connected to DB!
db stat waitDuration: 0s
db stat WaitCount: 0
2020/03/16 18:46:36 test begin...
2020/03/16 18:46:36 DoSomethingWithDBNotclose is called
db stat OpenConnections: 1
db stat InUse: 0
db stat Idle: 1
2020/03/16 18:46:36 DoSomethingWithDBNotclose is called
db stat OpenConnections: 1
db stat InUse: 1
db stat Idle: 0
row is scaned from first
row is scaned from first
row is scaned from first
row is scaned from first
row is scaned from first
row is scaned from first
row is scaned from first
row is scaned from first
2020/03/16 18:46:38 DoSomethingWithDBNotclose end...
row is scaned from second
row is scaned from second
row is scaned from second
row is scaned from second
row is scaned from second
row is scaned from second
row is scaned from second
row is scaned from second
2020/03/16 18:46:40 DoSomethingWithDBNotclose end...
2020/03/16 18:46:42 test is done...
    */
}

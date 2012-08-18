package main

import (
  "fmt"
  "time"
  _ "github.com/bmizerany/pq"
  "database/sql"
)

type User struct {
  Id      int64
  Name    string
  Created time.Time
} 

func main() {
  db, _ := sql.Open("postgres", "dbname=swift_test sslmode=disable")
  db.Exec("drop table if exists users")
  db.Exec("create table users(id serial, name text, created_at timestamp with time zone, primary key(id))")

  now := time.Now()
  for i := 0; i < 1000; i++ {
    db.Exec("insert into users(name, created_at) values($1, $2)", fmt.Sprintf("test %d", i), time.Now())
  } 
  fmt.Printf("pg #insert %v\n", time.Since(now))
      
  now = time.Now()
  for i := 0; i < 100; i++ {
    rows, _ := db.Query("select id, name, created_at from users")
    for rows.Next() {
      user := new(User)
      rows.Scan(&user.Id, &user.Name, &user.Created)
      // fmt.Printf("select %+v\n", user)
    }
  }
  fmt.Printf("pg #select %v\n", time.Since(now))
}

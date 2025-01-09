package main

import "fmt"

type DB struct {
	Dsn string
}

var Gdb *DB

func InitDB(dsn string) *DB {
	return &DB{
		Dsn: dsn,
	}
}

func main() {
	//Gdb = InitDB("fengfeng")
	fmt.Println(Gdb.Dsn)
}

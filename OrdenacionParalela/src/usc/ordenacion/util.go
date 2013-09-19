// María Sanmiguel Suarez. 2013

// Package ordenacion provides functions to generate sample data and to help
// testing the sorting functions.
package ordenacion

import (
	"math/rand"
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/mattn/go-sqlite3"
	"time"
)

// IsSorted checks wether a slice of integers is sorted. It does so by
// comparing each element with the next, and checking if it is smaller.
func IsSorted(a []int) bool{
	i:=0
	for i=0; i<len(a)-1; i++{
		if a[i]>a[i+1]{
		
			return false;
		}
	}
	return true

}

// Generates an array of integers with random values, with the size received
// as a parameter.
func CreateRandomArray(n int) []int{
	a := make([]int, n)
	t := time.Now().Nanosecond()
	rand.Seed(int64(t))
	for i:= 0; i < n; i++ {
		a[i]=rand.Int()
	}
	return a
}

// func AbrirBasePruebas() {
// 	db, err := sql.Open("sqlite3", "./prueba.db")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	db.Exec("create table foo (id integer not null primary key, name text)")
// }

// Generates an array of integers with descending values, with the size
// received as a parameter.
func CreateDescendingArray(n int) []int{
	var m int = n
	a := make([]int,n) 
	for i:= 0; i<n; i++{
		a[i] = m 
		m--
	}
	return a
}

// Generates an array of integers with ascending values, with the size
// received as a parameter
func CreateAscendingArray(n int) []int{
	a := make([]int,n)
	for i:= 0; i<n; i++{
		a[i] = i+1 
	}
	return a
}



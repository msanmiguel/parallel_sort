package ordenacion

import (
	"math/rand"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func EstaOrdenado(a []int) bool{
	i:=0
	for i=0; i<len(a)-1; i++{
		if a[i]>a[i+1]{
		
			return false;
		}
	}
	return true

}


func CrearArrayAleatorio(n int) []int{
	a := make([]int, n)
	t := time.Now().Nanosecond()
	rand.Seed(int64(t))
	for i:= 0; i < n; i++ {
		a[i]=rand.Int()
	}
	return a
}

func AbrirBasePruebas() {
	db, err := sql.Open("sqlite3", "./prueba.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Exec("create table foo (id integer not null primary key, name text)")
}

func CrearArrayDescendente(n int) []int{
	var m int = n
	a := make([]int,n) 
	for i:= 0; i<n; i++{
		a[i] = m 
		m--
	}
	return a
}


func CrearArrayAscendente(n int) []int{
	a := make([]int,n)
	for i:= 0; i<n; i++{
		a[i] = i+1 
	}
	return a
}



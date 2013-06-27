package main
import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

func main() { 
	fmt.Println(os.Getwd())
	err := os.Remove("./bdTest.db")
	if err != nil {
		fmt.Println(err)
		//return
	}

	db, err := sql.Open("sqlite3", "./bdTest.db")
	if err != nil {
		fmt.Println(err)
		return 
	}
	defer db.Close()

	sqls := []string{
		"CREATE TABLE Test (id INTEGER PRIMARY KEY, nombre TEXT, fecha DATETIME, nCores INTEGER)",
		"CREATE TABLE ResultadoTest (idResultado INTEGER PRIMARY KEY, idTest INTEGER, algoritmo TEXT, tamanhoEntrada INTEGER, tiempo INTEGER, FOREIGN KEY (idTest) REFERENCES Test(idTest))",
		"CREATE INDEX indiceTest ON Test(id)",
		"CREATE INDEX indiceResultado ON ResultadoTest(idResultado)",
		
	}
	for _, sql := range sqls {
		_, err = db.Exec(sql)
		if err != nil {
			fmt.Printf("%q: %s\n", err, sql)
			return
		}
	}
	tx, err := db.Begin()
	if err != nil { 
		fmt.Println(err)
		return
	}
	t := time.Now()
	//fecha := t.String()
		stmt, err := tx.Prepare("INSERT INTO Test(nombre, fecha, nCores) VALUES(?, ?, ?)")
		if err != nil {
			fmt.Println(err)
			return
		}
		_,err = stmt.Exec("test Prueba",t , 2)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer stmt.Close()

	tx.Commit()   
	fmt.Println("funciona")
	rows, err := db.Query("select id, nombre, fecha from Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var nombre string
		var fecha time.Time
		rows.Scan(&id, &nombre, &fecha)
		fmt.Println(id, nombre, fecha)
	}

}

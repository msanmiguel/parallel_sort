package daos

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_"os"
	"fmt"
	"time"
	//"strconv"
)

//var rutaBd string = "C:\\Users\\maria\\git\\OrdenacionParalela\\OrdenacionParalela\\"
var rutaBd string = "./"
var nombreBd string =  rutaBd + "bdTest.db"
type Test struct{
	IdTest int
	Fecha time.Time
	Nombre string
	NCores int
 
}
type ResultadoTest struct{
	IdResultado int
	IdTest int //clave foranea
	Algoritmo string
	TamanhoEntrada int
	Tiempo int
	
}


type TestDAO interface{
	CrearTest(test *Test) 
	BorrarTest(idTest int) 
	ObtenerTest(idTest int) *Test
	ActualizarTest(test *Test)
	
}

type ResultadoTestDAO interface{
	CrearResultado(resultado *ResultadoTest)
	BorrarResultado(idResultado int) 
	ObtenerResultado(idResultado int) *ResultadoTest
	ActualizarResultado(resultado *ResultadoTest)
	
}

type SqliteTestDAO struct {}
type SqliteResultadoDAO struct{}

func (s *SqliteTestDAO) CrearTest(test *Test){
	db, err := sql.Open("sqlite3", nombreBd)
	transaction, err := db.Begin()  // el tipo de transaction es Tx
	if err != nil { 
		fmt.Println(err)
		return
	}
	statement, err := transaction.Prepare("INSERT INTO Test(nombre, fecha, ncores)  VALUES( ?, ?, ? )") // el tipo de statemente es Stmt
	if err != nil{
		fmt.Println(err)
	}
	_, err = statement.Exec(test.Nombre, test.Fecha, test.NCores)
	if err != nil{
		fmt.Println(err)
	}
	row := transaction.QueryRow("SELECT last_insert_rowid()")
	row.Scan(&test.IdTest)
	statement.Close();
	transaction.Commit();
	db.Close();
}

func (s *SqliteTestDAO) BorrarTest(idTest int){
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	sentencia, err := db.Prepare("DELETE * FROM Test WHERE idTest = ?")
	if err != nil { 
		fmt.Println(err)
		return
	}
	_, err = sentencia.Exec(idTest)
	if err != nil{
		fmt.Println(err)
		return
	}
	db.Close();
}

func (s *SqliteTestDAO) ObtenerTest(idTest int) *Test{
	var test *Test = new(Test)
	test.IdTest = idTest
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	row:= db.QueryRow("SELECT nCores, nombre, fecha FROM Test WHERE id = ?", idTest)
	row.Scan(&test.NCores, &test.Nombre, &test.Fecha)
	test.Fecha = test.Fecha.Local()
	db.Close();
	return test
}

func (s *SqliteTestDAO) ActualizarTest(test *Test){
	db, err := sql.Open("sqlite3", nombreBd)
	sentencia, err := db.Prepare("UPDATE  Test SET fecha = ?, nombre = ?, nCores = ? WHERE id = ?")
	if err != nil{
		fmt.Println(err)
		return
	}
	_, err = sentencia.Exec(test.Fecha, test.Nombre, test.NCores, test.IdTest)
	if err != nil{
		fmt.Println(err)
		return
	}
	db.Close();
}

func (s *SqliteResultadoDAO) CrearResultado(resultado *ResultadoTest){
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	transaccion, err := db.Begin()
	if err != nil{
		fmt.Println(err)
	}
	sentencia, err := transaccion.Prepare("INSERT INTO ResultadoTest(idTest, algoritmo, tamanhoEntrada, tiempo) VALUES( ?, ?, ?, ?)")
	if err != nil{
		fmt.Println(err)
	}
	_, err = sentencia.Exec(resultado.IdTest, resultado.Algoritmo, resultado.TamanhoEntrada, resultado.Tiempo)
	if err != nil{
		fmt.Println(err)
	}
	row := transaccion.QueryRow("SELECT last_insert_rowid()")
	row.Scan(&resultado.IdResultado)
	sentencia.Close();
	transaccion.Commit();
	db.Close()
}

func (s *SqliteResultadoDAO) ObtenerResultadoPorIdtest(idTest int) []*ResultadoTest{
	var resultados []*ResultadoTest = []*ResultadoTest{}
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT idTest, algoritmo, tamanhoEntrada, tiempo FROM ResultadoTest WHERE idTest = ? ORDER BY idResultado", idTest)
	
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		resultado := &ResultadoTest {}
		rows.Scan(&resultado.IdTest, &resultado.Algoritmo, &resultado.TamanhoEntrada, &resultado.Tiempo)
		resultados = append(resultados, resultado)
	}
	
	db.Close()
	return resultados
}

func (s *SqliteResultadoDAO) ObtenerResultado(idResultado int) *ResultadoTest{
	var resultado *ResultadoTest = new(ResultadoTest)
	resultado.IdResultado = idResultado
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	row := db.QueryRow("SELECT idTest, algoritmo, tamanhoEntrada, tiempo FROM ResultadoTest WHERE idResultado = ?", idResultado)
	row.Scan(&resultado.IdTest, &resultado.Algoritmo, &resultado.TamanhoEntrada, &resultado.Tiempo)
	
	db.Close()
	return resultado
}

func (s *SqliteResultadoDAO) BorrarResultado(idResultado int){
	db, err:= sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	sentencia, err := db.Prepare("DELETE * FROM ResultadoTest WHERE idResultado = ?")
	_, err = sentencia.Exec()
	
	db.Close()
}

func (s *SqliteResultadoDAO) ActualizarResultado(resultado *ResultadoTest){

	bd, err:= sql.Open("sqlite3", nombreBd)
	if err !=  nil {
		fmt.Println(err)
		return
	}
	sentencia, err := bd.Prepare("UPDATE ResultadoTest SET idTest = ?, algoritmo = ?, tamanhoEntrada = ?, tiempo = ?  WHERE idResultado = ?")
	if err != nil{
		fmt.Println(err)
	}
	_, err = sentencia.Exec(resultado.IdTest, resultado.Algoritmo, resultado.TamanhoEntrada, resultado.Tiempo, resultado.IdResultado)
	if err != nil{
		fmt.Println(err)
	}
	bd.Close()
}
 

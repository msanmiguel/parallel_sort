package daos

import (
	"testing"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_"os"
	"fmt"
	"time"
)

func TestCrearTestDao(t *testing.T){
	var fecha time.Time
	fecha = time.Now().Local()
	test := &Test{-1, fecha, "test prueba", 2}
	dao := new(SqliteTestDAO)
	dao.CrearTest(test)
	var nCores int
	var nombre string
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	row := db.QueryRow("SELECT  fecha, nCores, nombre FROM Test WHERE id = ?", test.IdTest )
	err = row.Scan(&fecha, &nCores, &nombre)
	if err != nil {
		t.Fatalf("No se ha podido recuperar la fila insertada %s", err.Error())
	}
	if nombre != test.Nombre {
		t.Errorf("El nombre recuperado es incorrecto. %s != %s", nombre, test.Nombre)
	}
	if nCores != test.NCores{
		t.Errorf("El numero de cores recuperado es incorrecto. %d != %d", nCores, test.NCores)
	}
	fmt.Println(fecha)
	if fecha.Local() != test.Fecha {
		zona1, _ := fecha.Zone()
		zona2, _ := test.Fecha.Zone()
		t.Errorf("La fecha recuperada es incorrecta. %s (%s) != %s (%s)", fecha, zona1, test.Fecha, zona2)
	}
	if err != nil { 
		fmt.Println(err)
	}
	db.Close();
}

func TestActualizarTestDao(t *testing.T){
var fecha time.Time
	fecha = time.Now()
	test := &Test{1, fecha, "test probar actualizar test", 6}
	dao := new(SqliteTestDAO)
	dao.ActualizarTest(test)
	var nCores int
	var nombre string
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	//compruebo que se ha actualizado
	row := db.QueryRow("SELECT fecha, nCores, nombre FROM Test WHERE id = ?", test.IdTest)
	row.Scan(&fecha, &nCores, &nombre)
	
	if nombre != test.Nombre {
		t.Error("El nombre recuperado es incorrecto.")
	}
	if nCores != test.NCores{
		t.Error("El numero de cores recuperado es incorrecto.")
	}
	if fecha.Local() != test.Fecha{
		t.Error("La fecha recuperada es incorrecta." )
	}
	db.Close();
}

func TestObtenerTestDao(t *testing.T){
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	transaction, err := db.Begin()  // el tipo de transaction es Tx
	if err != nil { 
		fmt.Println(err)
		return
	}
	statement, err := transaction.Prepare("INSERT INTO Test(nombre, fecha, ncores)  VALUES( ?, ?, ? )") 
	if err != nil{
		fmt.Println(err)
	}
	var testId int
	var nombre string = "Test probar obtener test"
	fecha := time.Now()
	var nCores int = 2
	_, err = statement.Exec(nombre, fecha, nCores)
	if err != nil{
		fmt.Println(err)
	}
	row := transaction.QueryRow("SELECT last_insert_rowid()")
	row.Scan(&testId)
	
	statement.Close();
	transaction.Commit();
	db.Close();
	
	dao := new(SqliteTestDAO)
	test := &Test{}
	fmt.Printf("Obtener %d\n", testId)
	test = dao.ObtenerTest(testId)
	
	if nombre != test.Nombre {
		t.Errorf("El nombre recuperado es incorrecto. %s != %s\n", nombre, test.Nombre)
	}
	if nCores != test.NCores{
		t.Errorf("El numero de cores recuperado es incorrecto. %d != %d\n", nCores, test.NCores)
	}
	if fecha != test.Fecha{
		t.Error("La fecha recuperada es incorrecta.")
	}
}

func TestBorrarTest(t *testing.T){
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	transaction, err := db.Begin()  // el tipo de transaction es Tx
	if err != nil { 
		fmt.Println(err)
		return
	}
	statement, err := transaction.Prepare("INSERT INTO Test(nombre, fecha, ncores)  VALUES( ?, ?, ? )") 
	if err != nil{
		fmt.Println(err)
	}
	var testId int
	var nombre string = "Test Probar Borrar test"
	fecha := time.Now()
	var nCores int = 2
	_, err = statement.Exec(nombre, fecha, nCores)
	if err != nil{
		fmt.Println(err)
	}
	row := db.QueryRow("SELECT last_insert_rowid()")
	row.Scan(testId)
	
	statement.Close();
	transaction.Commit();
	
	
	dao := new(SqliteTestDAO)
	dao.BorrarTest(testId) 
	
	row = db.QueryRow("SELECT fecha, nCores, nombre FROM Test WHERE id = ?", testId)
	err = row.Scan(&fecha, &nCores, &nombre )
	if err == nil{
		t.Error("La fila no a sido borrada correctamente.\n")
	}
	
	
	db.Close();

}

func TestCrearResultado(t *testing.T){
	resultado:= &ResultadoTest{-1 ,1, "quicksort", 100000, 123444} // idResultadoTest, idTest, algoritmo, tamaño entrada, tiempo
	dao := new(SqliteResultadoDAO)
	dao.CrearResultado(resultado)
	
	var idTest, tamanhoEntrada, tiempo int
	var algoritmo string
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	row := db.QueryRow("SELECT idTest, algoritmo, tamanhoEntrada, tiempo FROM ResultadoTest WHERE idResultado = ?", resultado.IdResultado )
	err = row.Scan(&idTest, &algoritmo, &tamanhoEntrada, &tiempo)
	if err != nil {
		t.Fatalf("No se ha podido recuperar la fila insertada de id %d  %s",resultado.IdResultado, err.Error())
	}
	if idTest != resultado.IdTest {
		t.Errorf("El idTest recuperado es incorrecto. %d != %d", idTest, resultado.IdTest)
	}
	if algoritmo != resultado.Algoritmo{
		t.Errorf("El algoritmo recuperado es incorrecto %s != %s", algoritmo, resultado.IdResultado)
	}
	if  tamanhoEntrada != resultado.TamanhoEntrada{
		t.Error("El tamaño de la entrada recuperado es incorrecto %d != &d", tamanhoEntrada, resultado.TamanhoEntrada)
	}
	if  tiempo != resultado.Tiempo{
		t.Error("El tamaño de la entrada recuperado es incorrecto %d != &d", tiempo, resultado.Tiempo)
	}
	if err != nil { 
		fmt.Println(err)
	}
	db.Close();
}

func TestObtenerResultado(t *testing.T){
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	transaction, err := db.Begin()  // el tipo de transaction es Tx
	if err != nil { 
		fmt.Println(err)
		return
	}
	statement, err := transaction.Prepare("INSERT INTO ResultadoTest(idTest, algoritmo, tamanhoEntrada, tiempo)  VALUES( ?, ?, ?, ?)") 
	if err != nil{
		fmt.Println(err)
	}
	var idResultado int
	var idTest int = 1
	var algoritmo string = "Radix sort"
	var tamanhoEntrada int = 1000000
	var tiempo int = 234234
	_, err = statement.Exec(idTest, algoritmo, tamanhoEntrada, tiempo)
	if err != nil{
		fmt.Println(err)
	}
	row := transaction.QueryRow("SELECT last_insert_rowid()")
	row.Scan(&idResultado)
	
	statement.Close();
	transaction.Commit();
	db.Close();
	
	dao := new(SqliteResultadoDAO)
	resultadoTest := &ResultadoTest{}
	fmt.Printf("Obtener %d\n", idResultado)
	resultadoTest = dao.ObtenerResultado(idResultado)
	
	if algoritmo != resultadoTest.Algoritmo{
		t.Errorf("El algoritmo recuperado es incorrecto. %s != %s\n", algoritmo, resultadoTest.Algoritmo)
	}
	if idTest != resultadoTest.IdTest{
		t.Errorf("El el idTest recuperado es incorrecto. %d != %d\n", idTest, resultadoTest.IdTest)
	}
	if tamanhoEntrada != resultadoTest.TamanhoEntrada{
		t.Error("El tamaño de entrada recuperado es incorrecto. %d != %d\n", tamanhoEntrada, resultadoTest.TamanhoEntrada )
	}
	if tiempo != resultadoTest.Tiempo{
		t.Error("El tiempo recuperado es incorrecto. %d != %d\n", tiempo, resultadoTest.Tiempo )
	}
}

func TestActualizarResultado(t *testing.T){
	resultadoTest := &ResultadoTest{1, 1, "mergesort", 1000000, 623442}
	
	dao := new(SqliteResultadoDAO)
	dao.ActualizarResultado(resultadoTest)
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>%d\n", resultadoTest.IdResultado)
	var idTest int 
	var algoritmo string
	var tamanhoEntrada int
	var tiempo int
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	//compruebo que se ha actualizado
	row := db.QueryRow("SELECT idTest, algoritmo, tamanhoEntrada, tiempo FROM ResultadoTest WHERE idResultado = ?", resultadoTest.IdResultado)
	row.Scan(&idTest, &algoritmo, &tamanhoEntrada, &tiempo)
	
	if algoritmo != resultadoTest.Algoritmo{
		t.Errorf("El algoritmo recuperado es incorrecto. %s != %s\n", algoritmo, resultadoTest.Algoritmo)
	}
	if idTest != resultadoTest.IdTest{
		t.Errorf("El el idTest recuperado es incorrecto. %d != %d\n", idTest, resultadoTest.IdTest)
	}
	if tamanhoEntrada != resultadoTest.TamanhoEntrada{
		t.Error("El tamaño de entrada recuperado es incorrecto. %d != %d\n", tamanhoEntrada, resultadoTest.TamanhoEntrada )
	}
	if tiempo != resultadoTest.Tiempo{
		t.Error("El tiempo recuperado es incorrecto. %d != %d\n", tiempo, resultadoTest.Tiempo )
	}
	db.Close();
}

func TestBorrarResultado(t *testing.T){
	db, err := sql.Open("sqlite3", nombreBd)
	if err != nil{
		fmt.Println(err)
	}
	transaction, err := db.Begin()  // el tipo de transaction es Tx
	if err != nil { 
		fmt.Println(err)
		return
	}
	statement, err := transaction.Prepare("INSERT INTO ResultadoTest(idTest, algoritmo, tamanhoEntrada, tiempo)  VALUES( ?, ?, ?, ?)") 
	if err != nil{
		fmt.Println(err)
	}
	var idResultadoTest int
	var idTest int = 1
	var algoritmo string = "Mergesort"
	var tamanhoEntrada int = 1000000
	var tiempo int = 124141
	_, err = statement.Exec(idTest, algoritmo, tamanhoEntrada, tiempo)
	if err != nil{
		fmt.Println(err)
	}
	row := db.QueryRow("SELECT last_insert_rowid()")
	row.Scan(&idResultadoTest)
	
	statement.Close();
	transaction.Commit();
	
	
	dao := new(SqliteTestDAO)
	dao.BorrarTest(idResultadoTest) 
	
	row = db.QueryRow("SELECT idTest, algoritmo, tamanhoEntrada, tiempo FROM ResultadoTest WHERE idResultado = ?", idResultadoTest)
	err = row.Scan(&idTest, &algoritmo, &tamanhoEntrada, &tiempo)
	if err == nil{
		t.Error("La fila no a sido borrada correctamente.\n")
	}
	db.Close();
}
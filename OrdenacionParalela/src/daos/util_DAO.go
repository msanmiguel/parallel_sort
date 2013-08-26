package daos

import (

)

func GuardarResultadosTest(test *Test, resultados []ResultadoTest){
	dao := new(SqliteTestDAO)
	dao.CrearTest(test)
	dao1 := new(SqliteResultadoDAO)
	for _,v := range resultados {
		v.IdTest = test.IdTest
		dao1.CrearResultado(&v)
	} 
}

func CargarResultadosTest(idTest int) (*Test, []*ResultadoTest) {
	dao := new(SqliteTestDAO)
	t := dao.ObtenerTest(idTest)
	dao1 := new(SqliteResultadoDAO)
	r := dao1.ObtenerResultadoPorIdtest(idTest)

	return t, r
}

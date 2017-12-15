/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package repo

import (
	// "database/sql"
	"log"
	// 	"fmt"
	"github.com/jeffotoni/mercuriuscrud/conf"
	// 	"github.com/jeffotoni/mercuriuscrud/model"
	// 	uuid "github.com/satori/go.uuid"
)

// checando se existe dado especifico na base de dados
func GetOne(table, Key string, Id int) (R int, err error) {

	err = nil
	db, err := conf.GetDB()

	if err != nil {
		log.Println("[GetOne] Erro Pegando o DB: " + err.Error())
		return
	}

	sql := `select 
				count(*)
			from
				` + table + `
			where
				` + Key + `  = ?`

	err = db.Get(&R, sql, Id)

	if err != nil {

		log.Println("[GetOne] Erro ao executar a query: ", sql, " com o parametro: ", Id, " - Erro: ", err.Error())
		return
	}

	return
}

// criando conexcao com
// banco de dados
// db, err := conf.GetDB()

// // checando se
// // tudo ocorreu
// // bem na conexao
// if err != nil {
// 	log.Println(err.Error())
// 	return
// }

// // fazeno um teste
// // com ping
// err = db.Ping()
// if err != nil {
// 	log.Println(err.Error())
// 	return
// }

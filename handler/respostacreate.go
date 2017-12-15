/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package handler

import (
	"github.com/jeffotoni/mercuriuscrud/conf"
	"github.com/jeffotoni/mercuriuscrud/lib/context"
	"github.com/jeffotoni/mercuriuscrud/repo"
	"net/http"
)

// inserindo perguntas na base de dados
func RepostaCreate(ctx *context.Context) {

	var msgJson string

	db, err := conf.GetDB()

	if err != nil {

		msgJson = `{"status":"error","msg":"Ocorreu algo errado ao abrir o banco postgres: ` + err.Error() + `"}`
		//user is the struct you want to return
		ctx.JSON(http.StatusUnauthorized, msgJson)

	} else {

		createTable, err := repo.SqlTmplString("trespostas.sql")

		if err != nil {

			msgJson = `{"status":"error","msg":"Ocorreu algo errado ao abrir template sql: ` + err.Error() + `"}`

			//user is the struct you want to return
			ctx.JSON(http.StatusUnauthorized, msgJson)

		} else {

			_, err := db.Exec(createTable)

			if err != nil {

				msgJson = `{"status":"error","msg":"Exec create folder deleted [` + err.Error() + `]!"}`
				ctx.JSON(http.StatusUnauthorized, msgJson)

			} else {

				msgJson = `{"status":"ok","msg":"Create Table Respota"}`
				//user is the struct you want to return
				ctx.JSON(http.StatusUnauthorized, msgJson)
			}
		}

	}
}

/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package repo

import (
	// 	"fmt"
	// "database/sql"
	"github.com/jeffotoni/mercuriuscrud/conf"
	"github.com/jeffotoni/mercuriuscrud/model"
	"log"
	// 	uuid "github.com/satori/go.uuid"
)

// Insere o dado nat tabela
func InsertResposta(resp model.TRespostas) (Id int, err error) {

	err = nil
	db, err := conf.GetDB()
	if err != nil {
		log.Println("[InsertResposta] Erro ao criar instancia do DB: " + err.Error())
		return
	}

	// Rsp_uuid       string `json:"rsp_uid"`
	// Rsp_cod        int    `json:"rsp_cod "`
	// Rsp_per_cod    int    `json:"rsp_per_cod "`
	// Rsp_titulo     int    `json:"rsp_titulo"`
	// Rsp_correta    int    `json:"rsp_correta"`
	// Rsp_dtcadastro string `json:"rsp_dtcadastro"`
	// Rsp_datetime   string `json:"rsp_datetime"`

	insert := `INSERT INTO trespostas (
		rsp_uid,
		rsp_cod,
		rsp_per_cod,
		rsp_titulo,
		rsp_correta,
		rsp_dtcadastro,
		rsp_datetime)
		VALUES (
		?,
		?,
		?,
		?,
		?,
		?,
		?)`

	result, err := db.Exec(insert, resp.Rsp_uuid, resp.Rsp_cod, resp.Rsp_per_cod, resp.Rsp_titulo, resp.Rsp_correta, resp.Rsp_dtcadastro, resp.Rsp_datetime)
	if err != nil {
		log.Printf("[InsertResposta] Erro ao inserir. SQL: [%s] - Objeto: [%+v] - Error: [%s]\n", insert, resp, err.Error())
		return
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		log.Printf("[InsertResposta] Erro ao obter o ID. SQL: [%s] - Objeto: [%+v] - Error: [%s]\n", insert, resp, err.Error())
		return
	}
	Id = int(lastId)
	return
}

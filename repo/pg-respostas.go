/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package repo

import (
	"fmt"
	// "database/sql"
	"github.com/jeffotoni/mercuriuscrud/conf"
	"github.com/jeffotoni/mercuriuscrud/model"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

// Insere o dado nat tabela
func InsertResposta(resp model.TRespostas) (Uuid string, err error) {

	err = nil

	var Id int

	db, err := conf.GetDB()

	if err != nil {
		log.Println("[InsertResposta] Erro ao criar instancia do DB: " + err.Error())
		return
	}

	// criando uuid
	// para salvar
	// na base
	u1 := uuid.NewV4()

	// convertendo uuid para string
	resp.Rsp_uuid = fmt.Sprintf("%s", u1)

	DataNow := time.Now().Format(conf.LayoutDate)
	HourNow := time.Now().Format(conf.LayoutHour)

	resp.Rsp_datetime = DataNow + " " + HourNow

	insert := `INSERT INTO trespostas (
		rsp_uid,
		rsp_cod,
		rsp_per_cod,
		rsp_titulo,
		rsp_correta,
		rsp_dtcadastro,
		rsp_datetime)
		VALUES (
		:rsp_uuid,
		:rsp_cod,
		:rsp_per_cod,
		:rsp_titulo,
		:rsp_correta,
		:rsp_dtcadastro,
		:rsp_datetime) RETURNING rsp_id`

	stmt, err := db.PrepareNamed(insert)

	if err != nil {
		log.Printf("[InsertResposta] Erro ao Preparar o inserir. SQL: [%s] - Objeto: [%+v] - Error: [%s]\n", insert, err.Error())
		return
	}

	err = stmt.Get(&Id, resp)

	if err != nil {
		log.Printf("[InsertResposta] Erro ao inserir. SQL: [%s] - Objeto: [%+v] - Error: [%s]\n", insert, resp, err.Error())
		return
	}

	Uuid = resp.Rsp_uuid
	return
}

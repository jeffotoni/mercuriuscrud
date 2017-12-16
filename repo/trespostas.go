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
	"errors"
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

	// criando instancia
	db, err := conf.GetDB()

	if err != nil {
		msgerror = "[InsertResposta] Erro ao criar instancia do DB: " + err.Error()
		err = errors.New(msgerror)
		log.Println(msgerror)
		return
	}

	// validar se existe o dado no banco
	// if exist o dado na base de dados
	exist := GetOne(Ttrespostas, "rsp_cod", resp.Rsp_cod)

	if exist > 0 {

		msgerror = "[InsertResposta] Erro registro já existe na base de dados!"
		err = errors.New(msgerror)
		log.Println(msgerror)
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
		msgerror = "[InsertResposta] Erro ao Preparar o inserir. SQL: [" + insert + "] - Error: [" + err.Error() + "]"
		err = errors.New(msgerror)
		log.Println(msgerror)
		return
	}

	err = stmt.Get(&Id, resp)

	if err != nil {

		msgerror = "[InsertResposta] Erro ao fazer insert SQL: [" + insert + "] - Error: [" + err.Error() + "]"
		err = errors.New(msgerror)
		log.Println(msgerror)
		return
	}

	Uuid = resp.Rsp_uuid
	return
}

// create table
func CreateTable() (tables string, err error) {

	db, err := conf.GetDB()

	if err != nil {
		msgerror = "[QuestionsCreate] Erro ao conectar no banco de dados: : " + err.Error()
		log.Println(msgerror)
		return
	}

	createTableTrespostas, err := SqlTmplString(trespostas)

	if err != nil {

		msgerror = "[CreateTable] Error ao retornar Schema Sql: " + err.Error()
		log.Println(msgerror)
		return
	}

	// criando a tabela dinamicamente
	// no banco de dados
	_, err = db.Exec(createTableTrespostas)

	if err != nil {
		msgerror = "[CreateTable] Error ao tentar criar table: " + err.Error()
		log.Println(msgerror)
		return
	}

	tables = trespostas
	return
}

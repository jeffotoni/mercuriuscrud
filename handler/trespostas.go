/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package handler

import (
	"encoding/json"
	"github.com/jeffotoni/mercuriuscrud/lib/context"
	"github.com/jeffotoni/mercuriuscrud/model"
	"github.com/jeffotoni/mercuriuscrud/repo"
	"log"
	"net/http"
	"strings"
)

// inserindo perguntas na base de dados
func AnswersCreateTable(ctx *context.Context) {

	// create tables
	tables, err := repo.CreateTable()

	if err != nil {
		msgerror = "[AnswersCreateTable] Error ao criar tables: " + err.Error()
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}

	// sucesso
	msgJson = `{"status":"ok","msg":"tabela criada com sucesso", "tables":"` + tables + `"}`
	ctx.JSON(http.StatusOK, msgJson)
}

// inserindo perguntas na base de dados
func AnswersInsert(ctx *context.Context) {

	// struct model data
	var Tr model.TRespostas

	cType := ctx.Req.Header.Get("Content-Type")

	if strings.ToLower(strings.TrimSpace(cType)) != CTYPE_ACCEPT {

		msgerror = "[AnswersInsert] error no Content-Type: " + cType + ", aceitamos somente [Content-Type: " + CTYPE_ACCEPT + "]"
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}

	// capturando json findo da requisicao
	// estamos pegando em bytes para ser
	// usado no Unmarshal que recebe bytes
	byteJson, err := ctx.Req.Body().Bytes()

	if err != nil {
		msgerror = "[AnswersInsert] error em ler seu json: " + err.Error()
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}

	// fechando Req.Body
	defer ctx.Req.Body().ReadCloser()

	if string(byteJson) == "" {
		msgerror = "[AnswersInsert] error seu json nao pode ser vazio!"
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}

	// convert json to struct
	err = json.Unmarshal(byteJson, &Tr)

	if err != nil {

		msgerror = "[AnswersInsert] error ao tentar fazer Unmarshal em seu json: " + err.Error()
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}

	Uuid, err := repo.InsertResposta(Tr)

	if err != nil {

		msgerror = "[AnswersInsert] error ao tentar fazer insert: " + err.Error()
		log.Println(msgerror)
		msgJson = `{"status":"error","msg":"` + msgerror + `}`
		ctx.JSON(http.StatusUnauthorized, msgJson)
		return
	}

	// chamar
	msgJson = `{"status":"ok","msg":"seus dados foram inseridos com sucesso na base de dados!", "uuid":"` + Uuid + `"}`
	ctx.JSON(http.StatusOK, msgJson)
}

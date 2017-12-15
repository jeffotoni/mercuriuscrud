/*
* Go Library (C) 2017 Inc.
*
* @project     Project Mercurius test
* @author      @jeffotoni
* @size        14/12/2017
 */

package repo

import (
	"encoding/json"
	"github.com/jeffotoni/mercuriuscrud/model"
	"github.com/jeffotoni/mercuriuscrud/repo"
	bson "gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"strings"
)

// criando question
func AddQuestion(byteJson []byte) (Uuid string, err error) {

	// collection
	collection := "tpesqperguntas"

	// struct model data
	var Tp model.TPesqPerguntas
	// var TpC model.TPCondition

	// convert json to struct
	err := json.Unmarshal(byteJson, &Tp)

	//checar o erro
	if err != nil {
		log.Println("[AddQuestion] Erro nao consegui converter sua string em json: " + err.Error())
		return
	} else {

		// campo obrigatorio string
		if Tp.Ppr_dtcadastro == "" {
			log.Println("[AddQuestion] Erro A coluna Ppr_dtcadastro obrigatoria!")
			return
		} else if Tp.Ppr_dtaltera == "" { // campo obrigatorio string
			log.Println("[AddQuestion] Erro A coluna Ppr_dtaltera obrigatoria!")
			return
		} else {

			// validar se existe o dado no banco
			// if exist o dado no mongo nao faca
			exist, err := repo.FindExist(collection, bson.M{"ppr_cod": Tp.Ppr_cod})

			// validar
			if err != nil {
				log.Println("[AddQuestion] A collection pode esta errada ou sua condicao error: !" + err.Error())
				return
			} else {

				// se existe
				// nao faca
				// o insert
				if exist {
					log.Println("[AddQuestion] Error estes dados ja existe na base de dados!")
					return
				} else {

					// faca o insert no banco de dados
					// retorne o uuid do salvar
					// caso tenha erro informe ao cliente
					insertId, err := repo.Insert(collection, Tp)
					if err != nil {
						log.Println("[AddQuestion] ocorreu algum erro ao salvar seus dados error: !" + err.Error())
						return
					} else {
						Uuid = insertId
						return
					}
				}
			}
		}
	}
}

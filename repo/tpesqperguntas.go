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
	"errors"
	"github.com/jeffotoni/mercuriuscrud/model"
	bson "gopkg.in/mgo.v2/bson"
	"log"
)

// criando question
func AddQuestion(byteJson []byte) (Uuid string, err error) {

	// collection
	collection := "tpesqperguntas"

	// conter os erros
	var msgerror string

	// struct model data
	var Tp model.TPesqPerguntas

	//
	var exist bool

	// convert json to struct
	err = json.Unmarshal(byteJson, &Tp)

	//checar o erro
	if err != nil {
		log.Println("[AddQuestion] Erro nao consegui converter sua string em json: " + err.Error())
		return
	} else {

		//
		// campos obrigatorios
		//

		if Tp.Ppr_ordem <= 0 {
			msgerror = "[AddQuestion] Erro A coluna ppr_ordem obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_per_cod <= 0 {
			msgerror = "[AddQuestion] Erro A coluna Ppr_per_cod obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_cod <= 0 {
			msgerror = "[AddQuestion] Erro A coluna Ppr_cod obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_ppq_cod <= 0 {
			msgerror = "[AddQuestion] Erro A coluna Ppr_ppq_cod obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_dtcadastro == "" {
			msgerror = "[AddQuestion] Erro A coluna Ppr_dtcadastro obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else if Tp.Ppr_dtaltera == "" { // campo obrigatorio string
			msgerror = "[AddQuestion] Erro A coluna Ppr_dtaltera obrigatoria!"
			err = errors.New(msgerror)
			log.Println(msgerror)
			return
		} else {

			// validar se existe o dado no banco
			// if exist o dado no mongo nao faca
			exist, err = FindExist(collection, bson.M{"ppr_cod": Tp.Ppr_cod})

			// validar
			if err != nil {
				log.Println("[AddQuestion] A collection pode esta errada ou sua condicao error: !" + err.Error())
				return
			} else {

				// se existe
				// nao faca
				// o insert
				if exist {
					msgerror = "[AddQuestion] Error estes dados ja existe na base de dados!"
					err = errors.New(msgerror)
					log.Println(msgerror)
					return
				} else {

					// faca o insert no banco de dados
					// retorne o uuid do salvar
					// caso tenha erro informe ao cliente
					Uuid, err = Insert(collection, Tp)
					if err != nil {
						log.Println("[AddQuestion] ocorreu algum erro ao salvar seus dados error: !" + err.Error())
						return
					} else {

						return
					}
				}
			}
		}
	}
}

// criando question
func DelQuestion(Uuid string) (err error) {

	// mensagem error
	var msgerror string

	if Uuid == "" {

		msgerror = "[DelQuestion] Uuid obrigatorio!"
		err = errors.New(msgerror)
		log.Println(msgerror)
		return
	}

	// collection
	collection := "tpesqperguntas"

	// chamando o metodo para remover o registro do banco
	err = Remove(collection, bson.M{"ppr_uuid": Uuid})

	if err != nil {
		msgerror = "[DelQuestion] Algo errado ocorreu ao remover registro: " + err.Error()
		err = errors.New(msgerror)
		log.Println(msgerror)
		return
	}

	return
}

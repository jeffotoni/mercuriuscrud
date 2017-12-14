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
	"github.com/jeffotoni/mercuriuscrud/conf"
	"github.com/jeffotoni/mercuriuscrud/model"
	uuid "github.com/satori/go.uuid"
	bson "gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

// fazendo um insert em um collection especifica
// o metodo retorna o Uuid caso tenha sucesso
// ou error caso ocorra algum imprevisto
func Insert(namecollection string, insert model.TPesqPerguntas) (Uuid string, err error) {

	// fazendo um select no collection
	col, err := conf.GetMongoCollection(namecollection)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// garantindo que sera fechado
	// a session do banco assim
	// que finalizar
	defer col.Database.Session.Close()

	// criando uuid
	// para salvar
	// na base
	u1 := uuid.NewV4()

	// convertendo uuid para string
	insert.Ppr_uuid = fmt.Sprintf("%s", u1)

	DataNow := time.Now().Format(conf.LayoutDate)
	HourNow := time.Now().Format(conf.LayoutHour)

	insert.Ppr_datetime = DataNow + " " + HourNow

	// salvando na base de
	// dados o insert do
	// collection
	err = col.Insert(insert)

	// checando
	// se tudo
	// ocorreu bem
	if err != nil {
		log.Println(err.Error())
		return
	}

	// se tudo ocorreu bem
	// retorna Uuid para o usuario
	Uuid = string(insert.Ppr_uuid)
	return
}

// find Exist
func FindExist(namecollection string, Jcondition bson.M) (exist bool, err error) {

	col, err := conf.GetMongoCollection(namecollection)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//busca na base se existir os mesmos dados nao deixa inserir
	// Find and Count
	tpesqcount, err := col.Find(Jcondition).Count()

	if err != nil {

		exist = false
		log.Println(err.Error())
		return
	}

	if tpesqcount > 0 {

		// log.Println(err.Error())
		exist = true
		return
	}

	exist = false
	return
}

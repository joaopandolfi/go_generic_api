package dao

import (
	"../models"
	"github.com/joaopandolfi/blackwhale/remotes/mongo"
	"golang.org/x/xerrors"
	"gopkg.in/mgo.v2/bson"
)

type LambdaDAO interface {
	Save(models.Lambda)
}

type Lambda struct {}

func (d *Lambda) Save(lambda models.Lambda) error{
	return mongo.GenericInsert("lambda_data",lambda)
}

func (d *Lambda) GetByUser(userID int) ([]models.Lambda,error){
	var results []models.Lambda
	session, err := mongo.NewSession()
	if err != nil {
		err = xerrors.Errorf("Unable to connect to mongo: %v",err)
		return nil, err
	}

	err = session.GetCollection("labda_data").Find(bson.M{"userid": userID}).All(&results)

	if err != nil{
		err = xerrors.Errorf("Get by user error -> %v",err)
	}

	return results,err
}
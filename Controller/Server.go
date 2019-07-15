package Controller

// https://github.com/mongodb/mongo-go-driver
// go get github.com/derekparker/delve/cmd/dlv
import (
	logger "GoLibs/logs"
	"BaxEnd/Controller/global"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	routes *mux.Router
	porta  int
)

type item struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Coluna1 string             `bson:"Coluna1" json:"Coluna1"`
	Coluna2 string             `bson:"Coluna2" json:"Coluna2"`
}

type STRetorno struct {
	Erro      error
	Conectado bool
}

func NewRouter() *mux.Router {
	dirPublic := global.DirPublic()
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix(dirPublic).Handler(http.StripPrefix(dirPublic, http.FileServer(http.Dir("."+dirPublic))))
	return router
}

func ListenServer(sPorta int) {

	logger.Sucesso("Inicializando configurações do servidor http")
	porta = sPorta

	setRoutes()

	go func() {

		err := global.DBConnect()
		if err != nil {
			logger.Erro("Database Desconectado", err)
			return
		}

		err = http.ListenAndServe(":"+strconv.Itoa(porta), routes)
		if err != nil {
			logger.Erro("ListenAndServe: ", err)
			return
		}

	}()

}

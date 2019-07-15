package main

import (
	"BaxEnd/Controller"
	logger "GoLibs/logs"
	"os"
	"os/signal"
)

func main() {
	logger.DebugErro = true
	logger.DebugSucesso = true
	// logger.DebugOrigem = true

	Controller.ListenServer(2000)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	logger.Sucesso("Finalizando servidor")
	os.Exit(0)
}

package global

import (
	"os"
)

const (
	PRIVATE_DIR = "/private/"
	PUBLIC_DIR  = "/public/"
	PORT        = "80"
	Versao      = "1.0.0"
)

func DBConnect() error {

	// err := database.DBConnect()
	// if err == nil {
	// 	logger.Atencao("Mongo online")
	// }
	// return err
	return nil
}

func Load() {
	LoadConfigs()
}

func DirPublic() string {
	return PUBLIC_DIR
}

func DirPrivate() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	// Logger.Atencao("Diretorio private", dir+PRIVATE_DIR)
	return dir + PRIVATE_DIR
}

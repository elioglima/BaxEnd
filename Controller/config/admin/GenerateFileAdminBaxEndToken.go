package admin

import (
	"GoLibs/logs"
	"io/ioutil"
	"os"
)

func LocalGenerateFileBaxEndToken(i string) error {

	// preparar KeyAPI para o admin local para dev byteValueJSON
	// byteValueJSON, err := json.Marshal(i)
	// if err != nil {
	// 	return err
	// }

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	dirAdmin := dir + "/admin"
	_, err = os.Stat(dirAdmin)
	if err != nil {
		logs.Erro(dirAdmin)
		return err
	}

	// dirAdminRoot := dirAdmin + "/root"
	// _, err = os.Stat(dirAdminRoot)
	// if err != nil {
	// 	logs.Erro(dirAdminRoot)
	// 	return err
	// }

	// dirAdminKeys := dirAdminRoot + "/keys"
	// _, err = os.Stat(dirAdminKeys)
	// if err != nil {
	// 	logs.Erro(dirAdminKeys)
	// 	return err
	// }

	fileName := dirAdmin + "/.env"
	logs.Atencao(fileName)
	// Por fim vamos salvar em um arquivo JSON alterado.
	err = ioutil.WriteFile(fileName, []byte(i), 0644)
	if err != nil {
		logs.Erro(err)
		return err
	}

	logs.Sucesso("Chave criada com sucesso:" + fileName)
	return nil
}

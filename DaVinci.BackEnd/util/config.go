package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Config : Archivo de configuracion
type Config struct {
	PublicHTML   string `json:"public_html"`
	DataBasePath string `json:"database_path"`
}

//LoadFile :Carga archivos!
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//ReadConfigValues : Lee archivo de configuracion de Davinci
func ReadConfigValues(pathJSON string) Config {
	file, _ := ioutil.ReadFile(pathJSON)
	config := Config{}
	err := json.Unmarshal([]byte(file), &config)
	if err != nil {
		fmt.Println("Error al intentar leer archivo de configuracion")
	} else {
		fmt.Println("Config [OK]")
	}
	return config
}

//StringifyJSON , convierte cualquier objeto en una representacion STRING en formato JSON
func StringifyJSON(object interface{}) string {
	b, err := json.Marshal(object)
	if err != nil {
		return "Error"
	}
	return string(b)
}

//IsEmpty es un string vacio?
func IsEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	} else {
		return false
	}
}

//ExistsPath : Retorna si existe archivo o directorio
func ExistsPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

package socket

import (
	"encoding/json"
	"log"

	controllers "../controllers"
	models "../models"
	davinci "../util"
	socketio "github.com/googollee/go-socket.io"
)

//SocketServer : socket del juego!
type SocketServer struct {
}

//GetServer Servidor Socket de GOGOGO
func (sg *SocketServer) GetServer() *socketio.Server {
	server, err := socketio.NewServer(nil)
	log.Println("GET SERVER")
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("GET OnConnect")
		return nil
	})

	server.OnEvent("/", "doLogin", func(s socketio.Conn, login string) bool {
		s.SetContext("")
		log.Println("GET doLogin")
		online := (login == "admin")
		log.Println(login+"> Conectado  :", online)
		return online
	})

	server.OnEvent("/", "davinciSay", func(s socketio.Conn) {
		s.SetContext("")
		//log.Println("GET davinciSay")
		davinci := davinci.DavinciTales{}
		frase := davinci.DiLoTuyo()
		s.Emit("DiloTuyo", frase)
	})

	server.OnEvent("/", "import", func(s socketio.Conn, dto string, projectCode string, inventionCode string) {
		s.SetContext("")
		log.Println("GET import")
		var inventionCtrl = controllers.InventionController{}
		var projectCtrl = controllers.ProjectController{}

		invention := inventionCtrl.GetInventionByCode(inventionCode)

		type MapArtifactType map[string][]string
		data := make(MapArtifactType)
		var result = data
		json.Unmarshal([]byte(dto), &result)

		keyData := result[invention.KeyValue][0]
		/*
			log.Println("result ", davinci.StringifyJSON(result))
			log.Println("keyData ", davinci.StringifyJSON(keyData))
		*/
		log.Println("invention.KeyValue ", davinci.StringifyJSON(invention.KeyValue))
		pid := models.ProjectInventionData{}
		pid.ProjectCode = projectCode
		pid.KeyRef = invention.KeyValue
		pid.InventionCode = inventionCode

		dataInvention := models.DataInvention{}
		dataInvention.KeyValue = keyData
		/*
			log.Println("dto ", dto)
			log.Println("dataInvention ", davinci.StringifyJSON(dataInvention))
		*/
		dataInvention.ContentJSON = dto
		pid.Container = append(pid.Container, dataInvention)

		projectCtrl.AddInvention(pid)
		//TODO: Determinar si se actualizo o se creo invento.
		s.Emit("importResult", davinci.StringifyJSON(pid))
	})
	/*
		server.OnError("/", func(s socketio.Conn, e error) {
			log.Println("GET error")
			//log.Println("[!] meet error:", e)
		})
	*/
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		log.Println("GET closed")
		log.Println("[!] closed", msg)
	})

	return server
}

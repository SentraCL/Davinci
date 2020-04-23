package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	//"os"
	//"os/signal"
	"path/filepath"

	Models "./models"
	Routes "./routes"
	socket "./socket"
	util "./util"
)

func main() {
	/*
		//creo canal para la señaol desde el S.O.
		sigChan := make(chan os.Signal)
		//asigno todas las notificaciones al canal
		signal.Notify(sigChan)
		select {
		case sig := <-sigChan:
			fmt.Println("Davinci a sido interrumpido por el usuario.", sig)
		}
	*/

	var port string
	var https string
	var configPath string

	flag.StringVar(&port, "p", "8080", "Application Port")
	flag.StringVar(&https, "h", "N", "Go to HTTPS Server (Y/N)")
	flag.StringVar(&configPath, "c", "davinci.cfg", "Config Path")
	flag.Parse()
	configDef, err := filepath.Abs(configPath)

	if err != nil {
		fmt.Println("Config file not found.")
	} else {
		config := util.ReadConfigValues(configDef)
		if Models.WakeUPDataBase(config.DataBasePath) {
			fmt.Println("DataBase [OK]")
		}
		socket := socket.SocketServer{}
		io := socket.GetServer()
		go io.Serve()
		fmt.Println("Socket [OK]")
		defer io.Close()
		http.Handle("/socket.io/", io)

		r := Routes.Router{}
		r.SetPublicHTML(config.PublicHTML)
		http.Handle("/api/", r.Routers())
		http.Handle("/davinci/", r.SubRouters())
		http.Handle("/", http.FileServer(http.Dir(config.PublicHTML)))
		fmt.Println("Davinci in port :", port)

		if strings.ToUpper(https) == "Y" {
			fmt.Println("HTTPS [OK]")
			log.Fatal(http.ListenAndServeTLS(":"+port, "https-server.crt", "https-server.key", nil))
		} else {
			log.Fatal(http.ListenAndServe(":"+port, nil))
		}
	}
}

package main

import (
	"log"
	"net/http"
	handlers "rpsweb/handlers"
)

func main() {
	// Crear enrutador
	router := http.NewServeMux()

	//Manejador para servir los archivos estaticos
	fs := http.FileServer(http.Dir("static"))

	//Ruta para acceder a los archivos estaticos
	router.Handle("/static/", http.StripPrefix("/static/", fs)) //http.StripPrefix(): se utiliza para eliminar un prefijo de la URL de la solicitud HTTP antes de enrutarla al manejador.

	//Configurar las rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

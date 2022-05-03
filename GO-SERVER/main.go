package main

// importamos los packages necesarios para utilizar las tools y metodos.
import (
	"fmt"
	"log"
	"net/http"
)

// Handlers: utilizados para manejar funciones
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	adress := r.FormValue("adress")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Adress = %s\n", adress)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound) // Si el URL path es distinto a "HELLO" me trae el error 404.
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound) // Por default el metodo cuando utilizamos una URL es GET, por lo tanto si no
		return                                                        // Hacen un GET me arroja que el metodo no es soportado.
	}

}

func main() {
	// := Es un atajo para declarar variables
	fileServer := http.FileServer(http.Dir("./static")) // Trae los archivos del proyecto en el directorio donde estan guardados
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler) // Declaracion de las rutas para el proyecto "/, /form, /hello" (Se utilizan en la URL)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { //Si ocurre algun erro con el puerto expuesto me arroja un error fatal.
		log.Fatal(err)
	}
}

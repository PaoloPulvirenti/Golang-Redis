package main

import (
	"fmt"
	"net/http"

	"github.com/ProjectTwo/controller"
	router "github.com/ProjectTwo/http"
)

var (
	ClienteController  controller.ClienteController  = controller.NewClienteController()
	LibroController    controller.LibroController    = controller.NewLibroController()
	PrestitoController controller.PrestitoController = controller.NewPrestitoController()
	httpRouter         router.Router                 = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Eseguito")
	})

	httpRouter.GET("/clienteID/{id}", ClienteController.GetClienteByID)
	httpRouter.POST("/cliente", ClienteController.AddCliente)
	httpRouter.GET("/clienti", ClienteController.GetClienti)
	httpRouter.POST("/libro", LibroController.AddLibro)
	httpRouter.GET("/libri", LibroController.GetLibri)
	httpRouter.GET("/libroID/{id}", LibroController.GetLibroByID)
	httpRouter.PUT("/prenota/{id}", PrestitoController.PrestaLibro)
	httpRouter.POST("/prenota", PrestitoController.PrestaLibroPOST)
	httpRouter.PUT("/restituisci/{id}", PrestitoController.RestituisciLibro)
	httpRouter.POST("/restituisci", PrestitoController.RestituisciLibroPOST)

	httpRouter.SERVE(port)
}

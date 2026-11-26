package routes

import (
    "net/http"
    "meteo/controllers"
)

func SetupRoutes() {
    http.HandleFunc("/", controllers.HomeController)
    http.HandleFunc("/ville", controllers.VillesController)
    http.HandleFunc("/meteo", controllers.MeteoController)
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}
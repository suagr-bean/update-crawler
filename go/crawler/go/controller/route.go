package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

// 路由
func Route() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /add", AddController)
	mux.HandleFunc("GET /home", ShowController)
	mux.HandleFunc("GET /showdetail",ShowDetail)
	log.Println("Starting server on :8080")
	c := cors.New(cors.Options{
				AllowedOrigins:   []string{"*"}, 
						AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
								AllowedHeaders:   []string{"Content-Type", "Authorization"},
										AllowCredentials: true,
											})
	
	handler:=c.Handler(mux)
   port :=os.Getenv("PORT")
   if port == "" {
       port = "8080" 
	   }
	err := http.ListenAndServe("0.0.0.0:"+port, handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

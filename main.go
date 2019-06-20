package main

import (
	"database/sql"
	"flag"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"

	"github.com/PayloadPro/svc.pro.payload.bins/configs"
	"github.com/PayloadPro/svc.pro.payload.bins/deps"
	"github.com/PayloadPro/svc.pro.payload.bins/rpc"
	"github.com/PayloadPro/svc.pro.payload.bins/services"
)

func main() {

	var err error

	sa := getFlagConfig()

	// Services
	services := &deps.Services{
		Bin: &services.BinService{},
	}

	// Config
	config := &deps.Config{
		DB: &configs.DBConfig{},
	}
	config.Setup()

	// Create a DB Connection
	db, err := sql.Open("postgres", config.DB.DSN())
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	defer db.Close()

	// Add the DB to the Service
	services.Bin.DB = db

	router := createRouter(services, config)

	log.Fatal(http.ListenAndServe(*sa, handler(router)))
}

func handler(router *mux.Router) http.Handler {
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})

	return handlers.CORS(origins, headers, methods)(router)
}

func createRouter(services *deps.Services, config *deps.Config) *mux.Router {

	// Context
	rand.Seed(time.Now().UnixNano())
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	defer cancel()

	router := mux.NewRouter()

	router.HandleFunc("/bins", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewCreateBin(services, config)(ctx, r)
		})
	}).Methods("POST")

	router.HandleFunc("/bins", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewGetBins(services, config)(ctx, r)
		})
	}).Methods("GET")

	router.HandleFunc("/bins/{id}", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewGetBin(services, config)(ctx, r)
		})
	}).Methods("GET")

	return router
}

// getFlagConfig sets the runtime variables
func getFlagConfig() *string {

	fs := flag.NewFlagSet("", flag.ExitOnError)
	server := fs.String("server", "0.0.0.0", "HTTP server")
	port := fs.String("port", "8081", "HTTP server port")
	flag.Usage = fs.Usage

	fs.Parse(os.Args[1:])

	si := make([]string, 0)
	si = append(si, *server, *port)

	sa := strings.Join(si, ":")

	return &sa
}

package app

import (
    "fmt"
    "github.com/Admiral-Piett/bootstrap-go-http-service/app/handlers"
    "github.com/Admiral-Piett/bootstrap-go-http-service/app/models"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
    "gitlab.com/avarf/getenvs"
    "net/http"
    "os"
)

type App struct {
    Router *mux.Router
    Logger *logrus.Logger
}

func New() *App {
    logger := logrus.New()
    if os.Getenv("LOG_LEVEL") == "DEBUG" {
        logger.SetLevel(logrus.DebugLevel)
    } else {
        logger.SetLevel(logrus.InfoLevel)
    }

    logger.WithFields(logrus.Fields{"it's a": "debug log!"}).Debug("We're de-buggin' now")

    logger.Info("Starting Sound Control App...")

    InitializeSettings(logger)

    a := &App{
        Router: mux.NewRouter(),
        Logger: logger,
    }
    a.initRoutes()
    return a
}

func (a *App) initRoutes() {
    a.Router.HandleFunc("/api/auth", handlers.AuthHandler).Methods("POST")
    a.Router.HandleFunc("/{rest:.*}", handlers.NotFoundHandler()).Methods("GET", "POST", "PUT", "PATCH", "DELETE")
}

// --- CORS Proxy ---
// FIXME - For PRODUCTION update to allow only certain origins
var CORS_ALLOW_HEADERS = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
var CORS_ALLOW_METHODS = "POST, GET, OPTIONS, PUT, PATCH, DELETE"
var CORS_ALLOW_ORIGINS = "*"

func setupCORS(w *http.ResponseWriter, req *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", CORS_ALLOW_ORIGINS)
    (*w).Header().Set("Access-Control-Allow-Methods", CORS_ALLOW_METHODS)
    (*w).Header().Set("Access-Control-Allow-Headers", CORS_ALLOW_HEADERS)
}

// Proxy Handler to deal with all incoming requests in main.go.  If the Method is OPTIONS, assume this is a pre-flight
//	CORS check and return CORS headers here.
func (a *App) ProxyHandler(w http.ResponseWriter, req *http.Request) {
    setupCORS(&w, req)
    if req.Method == "OPTIONS" {
        return
    }
    rww := models.NewResponseWriterWrapper(w)

    a.Router.ServeHTTP(rww, req)
    a.Logger.Info(fmt.Sprintf("%d: %s %s", rww.StatusCode, req.Method, req.RequestURI))
}

func InitializeSettings(logger *logrus.Logger) {
    models.SETTINGS.Port = getenvs.GetEnvString("MUSIZTICLE_PORT", "9000")
}

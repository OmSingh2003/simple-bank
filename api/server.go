package api 

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/OmSingh2003/simple-bank/db/sqlc"
	"github.com/OmSingh2003/simple-bank/token"
	"github.com/OmSingh2003/simple-bank/util"
)

// Server serves HTTP requests for our banking service  
type Server struct {
	config util.Config
	store  db.Store 
	tokenMaker token.Maker 
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing 
func NewServer(config util.Config,store db.Store) (*Server,error) {
	tokenMaker , err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err!=nil {
		return nil , fmt.Errorf("cannot create token master: %w",err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker:tokenMaker}

	
	if v , ok  := binding.Validator.Engine().(*validator.Validate) ; ok {
		v.RegisterValidation("currency",validCurrency)
	}
	
	server.setUpRouter()

	return server , nil 
}

func (server *Server) setUpRouter() {
	router := gin.Default()
	
	// User routes
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	
	// Account routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	
	// Transfer routes
	router.POST("/transfers", server.createTransfer)
	
	server.router = router
}
// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// ServeHTTP implements http.Handler interface for testing
func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

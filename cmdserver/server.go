package cmdserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CmdCallback func(interface{}) chan string

type Command struct {

	// Name of the command
	Name string

	// Handler for the command
	Callback CmdCallback
}

type ApiRequest struct {

	// command to execute
	command string `json: "command"`

	// command arguments
	args []string `json: "args"`
}

type CmdServer struct {

	// mux router
	router *mux.Router

	// List of registered commands
	Commands []*Command
}

func New() *CmdServer {

	// Create server and router
	srv := CmdServer{}
	srv.router = mux.NewRouter()

	// Mount routes
	srv.router.HandleFunc("/execute/{command}", srv.ApiHandler).Methods("POST")
	srv.Commands = make([]*Command, 0)
	return &srv
}

func (srv *CmdServer) ApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if rawdata, err := io.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformed request"))
		return
	}
	var req ApiRequest
	if req, err := ParseApiBody()
}

func (srv *CmdServer) RegisterFunction(name string, callback CmdCallback) {
	cmd := Command{name, callback}
	srv.Commands = append(srv.Commands, &cmd)
}

package api

import (
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	address string
}

func NewApiServer(listenAddress string) *ApiServer {
	return &ApiServer{
		address: listenAddress,
	}
}

type CustomData struct {
	Message string `json:"message"`
}

func (s *ApiServer) Run() {
	router := NewCustomRouter()

	router.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(s.address, router)
}

type CustomRouter struct {
	mux *http.ServeMux
}

func NewCustomRouter() *CustomRouter {
	return &CustomRouter{
		mux: http.NewServeMux(),
	}
}

func (c *CustomRouter) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	c.mux.HandleFunc(pattern, handler)
}

func (c *CustomRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mux.ServeHTTP(w, r)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := CustomData{
		Message: "Hello, this is the home page!",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Quanghh2233/MyApp/models"
	"github.com/Quanghh2233/MyApp/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{
		service: service,
	}
}

func (e *EngineHandler) GetEngineByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := e.service.GetEngineByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error: ", err)
		return
	}
	body, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println("Error Writing Response: ", err)
		return
	}
}

func (e *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Err: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var EngineReq models.EngineRequest
	err = json.Unmarshal(body, &EngineReq)
	if err != nil {
		log.Println("Error while Unmarshalling Request body ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdEngine, err := e.service.CreateEngine(ctx, &EngineReq)
	if err != nil {
		log.Println("Error Creating Engine ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(createdEngine)
	if err != nil {
		log.Println("Error while marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, _ = w.Write(responseBody)
}

func (e *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error Reading Request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var EngineReq models.EngineRequest
	err = json.Unmarshal(body, &EngineReq)
	if err != nil {
		log.Println("Error while Unmarshalling Request body: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedEngine, err := e.service.UpdateEngine(ctx, id, &EngineReq)
	if err != nil {
		log.Println("Error while Updating Engine:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody, err := json.Marshal(updatedEngine)
	if err != nil {
		log.Println("Error while Marshalling Request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, _ = w.Write(resBody)
}

func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	deletedEngine, err := e.service.DeleteEngine(ctx, id)
	if err != nil {
		log.Println("Error while Deleting Engine:", err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": "Invalid ID or Engine not found"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)
		return
	}

	//check if success
	if deletedEngine.EngineID == uuid.Nil {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "Engine not found"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)
		return
	}

	jsonResponse, err := json.Marshal(deletedEngine)
	if err != nil {
		log.Println("Error while Marshalling deleted engine response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": "Internal server error"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, _ = w.Write(jsonResponse)
}

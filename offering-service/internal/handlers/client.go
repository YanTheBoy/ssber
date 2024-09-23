package handlers

import (
	"encoding/json"
	"net/http"
	"offering-service/internal/generated"
	"offering-service/internal/logger"
)

type OfferingImpl struct {
	log          logger.Log
	flatService FlatsManagerService
}

var _ gen.ServerInterface = (*OfferingImpl)(nil)

func New(log logger.Log, s FlatsManagerService) *OfferingImpl {
	return &OfferingImpl{
		log:          log,
		flatService: s,
	}
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	message := err.Error()
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(gen.ErrorResponse{Error: message})
}

func writeAuthError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
}

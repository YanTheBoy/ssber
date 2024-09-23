package handlers

import (
	"encoding/json"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	gen "offering-service/internal/generated"
	"offering-service/internal/model"
	otel2 "offering-service/internal/otel"
	"strconv"
)

func (h *OfferingImpl) GetFlat(w http.ResponseWriter, r *http.Request, flatId string, params gen.GetFlatParams) {
	// TODO add validation
	if params.XUserID <= 0 {
		writeAuthError(w)
		return
	}
	flatData := gen.CreateFlatJSONBody{}
	if err := json.NewDecoder(r.Body).Decode(&flatData); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	_, span := otel2.GetTracer().Start(r.Context(), "getFlat", trace.WithAttributes(attribute.String("userID", strconv.Itoa(params.XUserID))))
	defer span.End()

	var flat *model.Flat
	flat, err := h.flatService.GetFlat(r.Context(), flatId)

	if err != nil {
		span.SetStatus(codes.Error, "failed create flat")
		span.RecordError(err)
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	responseFlat := gen.Flat{
		Id:          flat.ID,
		Name:        flat.Name,
		Description: flat.Description,
		Price:       flat.Price,
		CreatedAt:   flat.CreatedAt,
	}
	_ = json.NewEncoder(w).Encode(responseFlat)

}

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

func (h *OfferingImpl) CreateFlat(w http.ResponseWriter, r *http.Request, params gen.CreateFlatParams) {
	h.log.Info("Create Flat by user ", params.XUserID)
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

	ctx, span := otel2.GetTracer().Start(r.Context(), "createFlat", trace.WithAttributes(attribute.String("userID", strconv.Itoa(params.XUserID))))
	defer span.End()

	createdFlat, err := h.flatService.CreateFlat(ctx, model.Flat{
		Name:        *flatData.Name,
		Price:       *flatData.Price,
		Description: *flatData.Description,
	})

	if err != nil {
		span.SetStatus(codes.Error, "failed create flat")
		span.RecordError(err)
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	responseFlat := gen.Flat{
		Id:          createdFlat.ID,
		Name:        createdFlat.Name,
		Description: createdFlat.Description,
		Price:       createdFlat.Price,
		CreatedAt: createdFlat.CreatedAt,
	}
	h.log.Info("Flat was successfully created", "id", responseFlat.Id)


	_ = json.NewEncoder(w).Encode(responseFlat)

}

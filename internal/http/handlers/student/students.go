package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/hitdhami0438/rest-api-go/internal/types"
	"github.com/hitdhami0438/rest-api-go/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		slog.Info("create a student")

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}

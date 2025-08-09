package report

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Ошибка чтения JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Валидэйшн
	if dto.CVURL == "" {
		http.Error(w, "Поле cv_url обязательно", http.StatusBadRequest)
		return
	}
	if dto.Email == "" {
		http.Error(w, "Поле email обязательно", http.StatusBadRequest)
		return
	}
	if dto.SMTPLogin == "" {
		http.Error(w, "Поле smtp_login обязательно", http.StatusBadRequest)
		return
	}
	if dto.SMTPPassword == "" {
		http.Error(w, "Поле smtp_password обязательно", http.StatusBadRequest)
		return
	}
	if dto.SMTPServer == "" {
		dto.SMTPServer = "smtp.gmail.com"
	}
	if dto.SMTPPort == 0 {
		dto.SMTPPort = 587
	}

	userId, err := h.service.Create(dto)
	if err != nil {
		http.Error(w, "Ошибка обработки отчёта: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Отчёт отправлен",
		"userId":  userId,
	})
}
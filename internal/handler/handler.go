package handler

import (
	"dork-otomasyon/internal/service"
	"html/template"
	"net/http"
)

type Handler struct {
	service *service.DorkService
	tmpl    *template.Template
}

func NewHandler(s *service.DorkService, t *template.Template) *Handler {
	return &Handler{
		service: s,
		tmpl:    t,
	}
}
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	err := h.tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Sayfa açılırken hata oluştu", http.StatusInternalServerError)
		return
	}
}
func (h *Handler) Generate(w http.ResponseWriter, r *http.Request) {
	target := r.FormValue("target")
	kategori := r.FormValue("category")
	results, err := h.service.Generate(target, kategori)
	if err != nil {
		http.Error(w, "Dorklar üretilirken hata oluştu", http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"Target":  target,
		"Results": results,
	}
	err = h.tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Sayfa gösterilirken hata oluştu", http.StatusInternalServerError)
		return
	}

}

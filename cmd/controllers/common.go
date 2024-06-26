package controllers

import (
	"embed"
	"fmt"
	"github.com/ad/domru/pkg/auth"
	"github.com/ad/domru/pkg/domru"
	"github.com/ad/domru/pkg/domru/constants"
	"github.com/ad/domru/pkg/home_assistant"
	"html/template"
	"log/slog"
	"net/http"
)

type Handler struct {
	Logger           *slog.Logger
	domruApi         *domru.APIWrapper
	credentialsStore auth.CredentialsStore

	TemplateFs embed.FS
}

func NewHandlers(templateFs embed.FS, credentialsStore auth.CredentialsStore, domruApi *domru.APIWrapper) (h *Handler) {
	h = &Handler{
		TemplateFs:       templateFs,
		Logger:           slog.Default(),
		credentialsStore: credentialsStore,
		domruApi:         domruApi,
	}

	return h
}

func (h *Handler) renderTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")

	templateFile := fmt.Sprintf("templates/%s.html.tmpl", templateName)
	tmpl, err := h.TemplateFs.ReadFile(templateFile)
	if err != nil {
		return fmt.Errorf("readfile %s: %w", templateFile, err)
	}

	t, err := template.New(templateName).Funcs(getTemplateFunctions()).Parse(string(tmpl))
	if err != nil {
		return fmt.Errorf("parse %s error: %w", templateFile, err)
	}

	err = t.Execute(w, data)
	if err != nil {
		return fmt.Errorf("execute %s error: %w", templateFile, err)
	}

	return nil
}

func getTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"getSnapshotUrl":     constants.GetSnapshotUrl,
		"getOpenDoorUrl":     constants.GetOpenDoorUrl,
		"getCameraStreamUrl": constants.GetCameraStreamUrl,
	}
}
func (h *Handler) determineBaseUrl(r *http.Request) string {
	var scheme string
	var host string

	if scheme = r.URL.Scheme; scheme == "" {
		scheme = "http"
	}
	haHost, haNetworkErr := home_assistant.GetHomeAssistantNetworkAddress()
	if haNetworkErr != nil {
		host = r.Host
	}
	ingressPath := r.Header.Get("X-Ingress-Path")
	if ingressPath == "" && haHost != "" {
		h.Logger.With("ha_host", haHost).Warn("X-Ingress-Path header is empty, when using Home Assistant host")
	}

	return fmt.Sprintf("%s://%s%s", scheme, host, ingressPath)
}

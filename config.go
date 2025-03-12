package show

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Config struct {
	RasterLayers map[string]string `json:"raster_layers,omitempty"`
	VectorLayers map[string]string `json:"vector_layers,omitempty"`
}

func ConfigHandler(cfg *Config) http.Handler {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		rsp.Header().Set("Content-type", "application/json")

		enc := json.NewEncoder(rsp)
		err := enc.Encode(cfg)

		if err != nil {
			slog.Error("Failed to encode site config", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
		}

		return
	}

	return http.HandlerFunc(fn)
}

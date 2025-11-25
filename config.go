package show

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// Config is a struct with configuration data for displaying PMTiles layers.
type Config struct {
	// RasterLayers is a map of "label" and "path" URIs for raster-based PMTiles layers to display.
	RasterLayers map[string]string `json:"raster_layers,omitempty"`
	// VectorLayers is a map of "label" and "path" URIs for vector-based PMTiles layers to display.
	VectorLayers map[string]string `json:"vector_layers,omitempty"`
}

// ConfigHandler returns a `http.Handler` instance for serving 'cfg' as a JSON-encoded string.
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

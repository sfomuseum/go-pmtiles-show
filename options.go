package show

import (
	"context"
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
	www_show "github.com/sfomuseum/go-www-show/v2"
)

type RunOptions struct {
	MapProvider       string
	MapTileURI        string
	InitialView       string
	LeafletStyle      string
	LeafletPointStyle string
	ProtomapsTheme    string
	RasterLayers      map[string]string `json:"raster_layers,omitempty"`
	VectorLayers      map[string]string `json:"vector_layers,omitempty"`
	Port              int
	Browser           www_show.Browser
	Verbose           bool
}

func RunOptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	raster_layers := make(map[string]string)
	vector_layers := make(map[string]string)

	for _, kv := range raster_tiles {
		k := kv.Key()
		path := kv.Value().(string)
		raster_layers[k] = path
	}

	for _, kv := range vector_tiles {
		k := kv.Key()
		path := kv.Value().(string)
		vector_layers[k] = path
	}

	opts := &RunOptions{
		MapProvider:       map_provider,
		MapTileURI:        map_tile_uri,
		LeafletStyle:      leaflet_style,
		LeafletPointStyle: leaflet_point_style,
		ProtomapsTheme:    protomaps_theme,
		InitialView:       initial_view,
		RasterLayers:      raster_layers,
		VectorLayers:      vector_layers,
		Port:              port,
		Verbose:           verbose,
	}

	br, err := www_show.NewBrowser(ctx, browser_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new browser, %w", err)
	}

	opts.Browser = br

	return opts, nil
}

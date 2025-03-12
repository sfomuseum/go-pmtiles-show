package show

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aaronland/go-http-maps/v2"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	"github.com/sfomuseum/go-www-show/v2"
)

var port int
var verbose bool

var browser_uri string

var initial_view string
var map_provider string
var map_tile_uri string
var protomaps_theme string
var leaflet_style string
var leaflet_point_style string

var raster_tiles multi.KeyValueString
var vector_tiles multi.KeyValueString

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("show")

	browser_schemes := show.BrowserSchemes()
	str_schemes := strings.Join(browser_schemes, ",")

	browser_desc := fmt.Sprintf("A valid sfomuseum/go-www-show/v2.Browser URI. Valid options are: %s", str_schemes)
	fs.StringVar(&browser_uri, "browser-uri", "web://", browser_desc)

	fs.StringVar(&map_provider, "map-provider", "leaflet", "Valid options are: leaflet, protomaps")
	fs.StringVar(&map_tile_uri, "map-tile-uri", maps.LEAFLET_OSM_TILE_URL, "A valid Leaflet tile layer URI. See documentation for special-case (interpolated tile) URIs.")
	fs.StringVar(&protomaps_theme, "protomaps-theme", "white", "A valid Protomaps theme label (for the base map not individual PMTiles databases).")
	fs.StringVar(&leaflet_style, "leaflet_style", "", "A custom Leaflet style definition for geometries. This may either be a JSON-encoded string or a path on disk.")
	fs.StringVar(&leaflet_point_style, "leaflet_point_style", "", "A custom Leaflet style definition for points. This may either be a JSON-encoded string or a path on disk.")
	fs.StringVar(&initial_view, "initial-view", "", "A comma-separated string indicating the map's initial view. Valid options are: 'LON,LAT', 'LON,LAT,ZOOM' or 'MINX,MINY,MAXX,MAXY'.")

	fs.Var(&raster_tiles, "raster", "Zero or more {LAYER_NAME}={PATH} pairs referencing PMTiles databases containing raster data.")
	fs.Var(&vector_tiles, "vector", "Zero or more {LAYER_NAME}={PATH} pairs referencing PMTiles databases containing vector (MVT) data.")

	fs.IntVar(&port, "port", 0, "The port number to listen for requests on (on localhost). If 0 then a random port number will be chosen.")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Command-line tool for serving PMTiles tiles from an on-demand web server.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	return fs
}

package show

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aaronland/go-http-maps/v2"
	"github.com/sfomuseum/go-pmtiles-show/static/www"
	www_show "github.com/sfomuseum/go-www-show/v2"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	if opts.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	mux := http.NewServeMux()

	maps_opts := &maps.AssignMapConfigHandlerOptions{
		MapProvider:       opts.MapProvider,
		MapTileURI:        opts.MapTileURI,
		InitialView:       opts.InitialView,
		LeafletStyle:      opts.LeafletStyle,
		LeafletPointStyle: opts.LeafletPointStyle,
		ProtomapsTheme:    opts.ProtomapsTheme,
	}

	err := maps.AssignMapConfigHandler(maps_opts, mux, "/map.json")

	if err != nil {
		return fmt.Errorf("Failed to assign map config handler, %w", err)
	}

	www_fs := http.FS(www.FS)
	mux.Handle("/", http.FileServer(www_fs))

	www_show_opts := &www_show.RunOptions{
		Port:    opts.Port,
		Browser: opts.Browser,
		Mux:     mux,
	}

	return www_show.RunWithOptions(ctx, www_show_opts)
}

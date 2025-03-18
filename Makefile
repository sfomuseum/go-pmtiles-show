GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/show cmd/show/main.go

raster:
	go run cmd/show/main.go \
		-initial-view -122.408061,37.601617,-122.354907,37.640167 \
		-raster test=fixtures/1930-raster.pmtiles

vector:
	go run cmd/show/main.go \
		-initial-view -122.408061,37.601617,-122.354907,37.640167 \
		-vector sfo=fixtures/sfo.pmtiles


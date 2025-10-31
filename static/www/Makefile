LEAFLET_VERSION=1.9.4
PROTOMAPS_VERSION=4.3.0
PROTOMAPS_LEAFLET_VERSION4=4.0.1
PROTOMAPS_LEAFLET_VERSION5=5.1.0

leaflet:
	curl -s -o css/leaflet-$(LEAFLET_VERSION).css https://unpkg.com/leaflet@$(LEAFLET_VERSION)/dist/leaflet.css
	curl -s -o javascript/leaflet-$(LEAFLET_VERSION).js https://unpkg.com/leaflet@$(LEAFLET_VERSION)/dist/leaflet.js

protomaps:
	curl -s -o javascript/protomaps-leaflet-$(PROTOMAPS_LEAFLET_VERSION4).js https://unpkg.com/protomaps-leaflet@$(PROTOMAPS_LEAFLET_VERSION4)/dist/protomaps-leaflet.js
	curl -s -o javascript/protomaps-leaflet-$(PROTOMAPS_LEAFLET_VERSION5).js https://unpkg.com/protomaps-leaflet@$(PROTOMAPS_LEAFLET_VERSION5)/dist/protomaps-leaflet.js
	curl -s -o javascript/pmtiles-$(PROTOMAPS_VERSION).js https://unpkg.com/pmtiles@$(PROTOMAPS_VERSION)/dist/pmtiles.js

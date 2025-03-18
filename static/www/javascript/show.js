// import { PMTiles, leafletRasterLayer } from 'pmtiles';

window.addEventListener("load", function load(event){

    // Null Island    
    var map = L.map('map').setView([0.0, 0.0], 1);    

    var fetch_site_config = function(){

	return new Promise((resolve, reject) => {

	    fetch("/config.json")
		.then((rsp) => rsp.json())
		.then((cfg) => {
		    resolve(cfg);
		}).catch((err) => {
		    reject(err);
		    return;
		});
	});
    };
    
    var fetch_map_config = function(){

	return new Promise((resolve, reject) => {
	    
	    fetch("/map.json")
		.then((rsp) => rsp.json())
		.then((cfg) => {
		    
		    switch (cfg.provider) {
			case "leaflet":
			    
			    var tile_url = cfg.tile_url;
			    
			    var tile_layer = L.tileLayer(tile_url, {
				maxZoom: 19,
			    });
			    
			    tile_layer.addTo(map);
			    break;
			    
			case "protomaps":
			    
			    var tile_url = cfg.tile_url;
			    
			    var tile_layer = protomapsL.leafletLayer({
				url: tile_url,
				theme: cfg.protomaps.theme,
			    })
			    
			    tile_layer.addTo(map);
			    break;
			    
			default:
			    console.error("Uknown or unsupported map provider");
			    reject(err);
			    return;
		    }
		    
		    if (cfg.initial_view) {
			
			var zm = map.getZoom();
			
			if (cfg.initial_zoom){
			    zm = cfg.initial_zoom;
			}
			
			map.setView([cfg.initial_view[1], cfg.initial_view[0]], zm);
			
		    } else if (cfg.initial_bounds){
			
			var bounds = [
			    [ cfg.initial_bounds[1], cfg.initial_bounds[0] ],
			    [ cfg.initial_bounds[3], cfg.initial_bounds[2] ],
			];
			
			map.fitBounds(bounds);
		    }
		    
		    resolve(cfg);
		    
		}).catch((err) => {
		    console.error("Failed to derive map config", err);
		    reject(err);
		    return;
		});    
	});
    };
    
    fetch_map_config().then((map_cfg) => {

	console.log(map_cfg);
	
	fetch_site_config().then((site_cfg) => {

            var base_maps = {};
            var overlays = {};

	    for (label in site_cfg.vector_layers) {

		var tile_url = site_cfg.vector_layers[label];

		// Important: The "show" theme is NOT part of the default protomaps/protomaps-leaflet package.
		    
		var tile_layer = protomapsL.leafletLayer({
                    url: tile_url,
		    theme: 'show',
		})

		tile_layer.addTo(map);
		base_maps[label] = tile_layer;
	    }

	    for (label in site_cfg.raster_layers){

		var tile_url = site_cfg.raster_layers[label];
		
		const p = new pmtiles.PMTiles(tile_url);
		const tile_layer = pmtiles.leafletRasterLayer(p);
		
		tile_layer.addTo(map);
		base_maps[label] = tile_layer;		
	    }
	    
            var layerControl = L.control.layers(base_maps, overlays);
            layerControl.addTo(map);
	    
	}).catch((err) => {
	    console.error("Failed to fetch site config", err);
	});
	    
    }).catch((err) => {
	console.error("Failed to fetch map config", err);
    });
});

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
		    
		    // To do: Set bounding box from configs (if defined)
		    resolve(cfg);
		    
		}).catch((err) => {
		    console.error("Failed to derive map config", err);
		    reject(err);
		    return;
		});    
	});
    };
    
    fetch_map_config().then((map_cfg) => {

	fetch_site_config().then((site_cfg) => {

	    console.log("SITE", site_cfg);
	    
	}).catch((err) => {
	    console.error("Failed to fetch site config", err);
	});
	    
    }).catch((err) => {
	console.error("Failed to fetch map config", err);
    });
});

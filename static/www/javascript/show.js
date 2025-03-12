window.addEventListener("load", function load(event){

    // Null Island    
    var map = L.map('map').setView([0.0, 0.0], 1);    
    
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
                    return;
	    }

	    // To do: Set bounding box from configs (if defined)
	    
	    console.log("Okay protomaps");
	    
        }).catch((err) => {
	    console.error("Failed to derive map config", err);
	    return;
	});    
    
});

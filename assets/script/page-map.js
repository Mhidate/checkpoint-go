var mymap = L.map("mapid").setView([-8.3730354, 116.4608215], 12);
L.tileLayer(
    "https://server.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer/tile/{z}/{y}/{x}",
    {
        attribution:
            "Tiles &copy; Esri &mdash; Source: Esri, i-cubed, USDA, USGS, AEX, GeoEye, Getmapping, Aerogrid, IGN, IGP, UPR-EGP, and the GIS User Community",
    }
).addTo(mymap);

//tampilkan ori
fetch("/tampil-marker")
    .then(response => response.json())
    .then(data => {
        var locations = data.locations;

        locations.forEach(function (location) {
            var picon = L.icon({
                iconUrl: 'images/marker.png',
                iconSize: [36, 41],
                iconAnchor: [12, 41],
                popupAnchor: [1, -34],
                tooltipAnchor: [16, -28],
                shadowSize: [41, 41]
            });

            L.marker([parseFloat(location.lat), parseFloat(location.longt)], {
                icon: picon
            }).addTo(mymap).bindPopup(
                `${location.nama_pos}<br>dengan Latitude dan Longitude ${location.lat}, ${location.longt}<br>keterangan: ${location.keterangan}`
            );
        });
    })
    .catch(error => {
        console.error("Error:", error);
    });
//l
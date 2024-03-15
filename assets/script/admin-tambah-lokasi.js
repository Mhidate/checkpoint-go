var mymap = L.map("mapid").setView([-8.3730354, 116.4608215], 12);
L.tileLayer(
    "https://tile.openstreetmap.org/{z}/{x}/{y}.png",
    {
        attribution: "Tiles &copy; Esri &mdash; Source: Esri, i-cubed, USDA, USGS, AEX, GeoEye, Getmapping, Aerogrid, IGN, IGP, UPR-EGP, and the GIS User Community"
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

var popup = L.popup();

function onMapClick(e) {
    popup
        .setLatLng(e.latlng)
        .setContent("Koordinatnya adalah " + e.latlng.toString())
        .openOn(mymap);

    document.getElementById("latitude").value = e.latlng.lat;
    document.getElementById("longitude").value = e.latlng.lng;
}

mymap.on("click", onMapClick);



var form = document.getElementById("addMarkerForm");
form.addEventListener("submit", function (event) {
    event.preventDefault();

    var formData = new FormData(form);
    var data = {};
    formData.forEach(function (value, key) {
        data[key] = value;
    });

    fetch("/add-marker", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json"
        }
    })
        .then(function (response) {
            if (!response.ok) {
                alert("Gagal menghapus data!");
                throw new Error("Failed to add marker");
            } else {
                alert("Berhasil menambahkan pos");
            }
            return response.json();
        })
        .then(function (data) {
            console.log("Marker added successfully:", data);
            // Clear form fields after successful submission
            form.reset();
            document.getElementById("latitude").value = "";
            document.getElementById("longitude").value = "";
            // Optional: Add the new marker to the map
            var picon = L.icon({
                iconUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon.png',
                iconSize: [25, 41],
                iconAnchor: [12, 41],
                popupAnchor: [1, -34],
                tooltipAnchor: [16, -28],
                shadowSize: [41, 41]
            });
            L.marker([data.latitude, data.longitude], { icon: picon })
                .addTo(mymap)
                .bindPopup(`<b>${data.nama_pos}</b><br>Latitude: ${data.latitude}<br>Longitude: ${data.longitude}<br>Keterangan: ${data.keterangan}`)
                .openPopup();
        })
        .catch(function (error) {
            console.error("Error adding marker:", error);
        });
});
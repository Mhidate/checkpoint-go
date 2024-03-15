$(document).ready(function () {
    // Function untuk memuat data dari server
    function loadData(url) {
        $.get(url, function (data) {
            var groups = {}; // Objek untuk menyimpan grup-grup unik
            $.each(data, function (index, item) {
                // Menambahkan grup ke objek groups
                if (!groups.hasOwnProperty(item.grup)) {
                    groups[item.grup] = $("<div class='column' id='grup-" + item.grup + "'>" +
                        "<table>" +
                        "<thead>" +
                        "<tr>" +
                        "<th class='column-primary' data-header='Pendaki'><span>Nama :</span></th>" +
                        "<th>Grup :</th>" +
                        "<th>Alamat :</th>" +
                        "<th>POS :</th>" +
                        "<th>Tanggal/Waktu :</th>" +
                        "</tr>" +
                        "</thead>" +
                        "<tbody id='data-grup-" + item.grup + "'>" +
                        "</tbody>" +
                        "</table>" +
                        "</div>");
                    $("#columns-container").append(groups[item.grup]);
                }

                // Membuat baris data
                var row = "<tr>" +
                    "<td data-header='Nama:'>" + item.nama + "</td>" +
                    "<td data-header='Grup:'>" + item.grup + "</td>" +
                    "<td data-header='Alamat:'>" + item.alamat + "</td>" +
                    "<td data-header='Pos:'>" + item.pos + "</td>" +
                    "<td data-header='Tanggal/Waktu:'>" + item.waktu + "</td>" +
                    "</tr>";

                // Menambahkan baris ke tbody yang sesuai dengan grup
                $("#data-grup-" + item.grup).append(row);
            });
        });
    }

    // Mengambil data pertama kali saat halaman dimuat
    loadData("http://localhost:8080/data");
});
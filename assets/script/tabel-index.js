
$(document).ready(function () {
    // Function untuk memuat data dari server
    function loadData(url) {
        $.get(url, function (data) {
            $("#data").empty(); // Mengosongkan tabel sebelum menambahkan data baru
            $.each(data, function (index, item) {
                var row = "<tr>" +
                    "<td data-header='Nama:'>" + item.nama + "</td>" +
                    "<td data-header='Grup:'>" + item.grup + "</td>" +
                    "<td data-header='Alamat:'>" + item.alamat + "</td>" +
                    "<td data-header='Pos:'>" + item.pos + "</td>" +
                    "<td data-header='Tanggal/Waktu:'>" + item.waktu + "</td>" +
                    "<td data-header='Lokasi pos:'><a href='/peta'>Lihat</a></td>" +
                    "</tr>";

                $("#data").append(row);
            });
        });
    }

    // Mengambil data pertama kali saat halaman dimuat
    loadData("http://localhost:8080/data");

    // Event listener untuk form pencarian
    $("#searchForm").submit(function (event) {
        event.preventDefault(); // Mencegah form melakukan submit

        var searchQuery = $("#searchInput").val(); // Mendapatkan nilai dari input pencarian
        var url = "http://localhost:8080/cari-log?cari=" + searchQuery; // URL endpoint pencarian

        loadData(url); // Memuat data berdasarkan hasil pencarian
    });
});
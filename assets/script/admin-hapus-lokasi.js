$(document).ready(function () {
    // Function untuk memuat data dari server
    function loadData(url) {
        $.get(url, function (response) {
            $("#data").empty(); // Mengosongkan tabel sebelum menambahkan data baru
            $.each(response.locations, function (index, item) {
                var row = "<tr>" +
                    "<td>" + item.id + "</td>" +
                    "<td>" + item.lat + "</td>" +
                    "<td>" + item.longt + "</td>" +
                    "<td>" + item.nama_pos + "</td>" +
                    "<td>" + item.keterangan + "</td>" +
                    "<td><button class='b2' data-id='" + item.id + "'><img src='images/button-hapus.svg'  width='20'height='20'></button></td>" +
                    "</tr>";

                $("#data").append(row);
            });
        });
    }

    // Mengambil data pertama kali saat halaman dimuat
    loadData("http://localhost:8080/tampil-marker");

    // Event listener untuk tombol hapus
    $(document).on("click", ".b2", function () {
        var id = $(this).data("id");
        $.post("/hapus-pos", { id: id }, function (response) {
            // Tampilkan pesan yang diterima dari server
            alert(response.message);
            // Reload halaman saat berhasil menghapus data
            if (response.success) {
                loadData("http://localhost:8080/tampil-marker");
            }
        });
    });
});
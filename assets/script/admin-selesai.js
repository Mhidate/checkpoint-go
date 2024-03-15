$(document).ready(function () {
    // Function untuk memuat data dari server
    function loadData(url) {
        $.get(url, function (data) {
            $("#data").empty(); // Mengosongkan tabel sebelum menambahkan data baru
            $.each(data, function (index, item) {
                var row = "<tr>" +
                    "<td data-header='Nama:'>" + item.nama + "</td>" +
                    "<td data-header='Grup:'>" + item.grup + "</td>" +
                    "<td data-header='Grup:'>" + item.jenis + "</td>" +
                    "<td data-header='Grup:'>" + item.nik + "</td>" +
                    "<td data-header='Alamat:'>" + item.alamat + "</td>" +
                    "<td data-header='Tanggal selesai:'>" + item.tgl + "</td>" +
                    "<td data-header='Catatan:'>" + item.catatan + "</td>" +
                    "</tr>";

                $("#data").append(row);
            });
        });
    }

    // Mengambil data pertama kali saat halaman dimuat
    loadData("http://localhost:8080/data-selesai");


    // Ambil parameter dari URL
    const params = new URLSearchParams(window.location.search);
    const status = params.get('status');

    if (status === 'success') {
        alert('Berhasil');
    }

    else if (status === 'error') {
        alert('GAGAL!! ');
    }


    // Mendapatkan URL saat ini
    var currentUrl = window.location.href;

    // Menghapus parameter URL
    function removeUrlParameter(parameter) {
        var url = new URL(currentUrl);
        url.searchParams.delete(parameter);
        window.history.replaceState({}, '', url);
    }

    // Contoh penggunaan: Menghapus parameter 'page'
    removeUrlParameter('status');

});
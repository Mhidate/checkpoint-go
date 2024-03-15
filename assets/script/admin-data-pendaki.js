$(document).ready(function () {
    // Function untuk memuat data dari server
    function loadData(url) {
        $.get(url, function (pendaki) {
            $("#data2").empty(); // Mengosongkan tabel sebelum menambahkan data baru
            $.each(pendaki.pendaki, function (index, item) {
                var row = "<tr>" +
                    "<td data-header='ID Gelang'>" + item.id + "</td>" +
                    "<td data-header='Grup'>" + item.grup + "</td>" +
                    "<td data-header='Nama'>" + item.nama + "</td>" +
                    "<td data-header='Jenis kelamin'>" + item.jenis + "</td>" +
                    "<td data-header='NIK'>" + item.nik + "</td>" +
                    "<td data-header='Alamat'>" + item.alamat + "</td>" +
                    "<td data-header='Edit'><button class='b1' data-id='" + item.id + "'><img src=images/button-edit.svg width=20 height=20 /></button></td>" +
                    "<td data-header='Hapus'><button class='b2' data-id='" + item.id + "'><img src=images/button-hapus.svg width=20 height=20 /></button></td>" +
                    "<td data-header='Hapus'><button class='b3' data-id='" + item.id + "'><img src=images/button-selesai.svg width=20 height=20 /></button></td>" +
                    "</tr>";

                $("#data2").append(row);
            });
        });
    }

    // Mengambil data pertama kali saat halaman dimuat
    loadData("http://localhost:8080/jsun-pendaki");

    // Event listener untuk form pencarian
    $("#pencarian").submit(function (event) {
        event.preventDefault(); // Mencegah form melakukan submit

        var searchQuery = $("#searchInput").val(); // Mendapatkan nilai dari input pencarian
        var url = "http://localhost:8080/cari-pendaki?cari=" + searchQuery; // URL endpoint pencarian

        loadData(url); // Memuat data berdasarkan hasil pencarian
    });

    // Event listener untuk tombol Hapus
    $(document).on("click", ".b2", function () {
        var id = $(this).data("id");
        $.post("/hapus-pendaki", { id: id }, function (response) {
            // Reload halaman setelah penghapusan data
            if (response.success) {
                loadData("http://localhost:8080/jsun-pendaki");
                alert("Data berhasil dihapus!");
            } else {
                alert("Gagal menghapus data!");
            }
        });
    });



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

    //event untuk edit
    $(document).ready(function () {
        // Event listener untuk tombol Edit
        $(document).on("click", ".b1", function () {
            var id = $(this).data("id");

            // Lakukan apa yang perlu dilakukan saat tombol Edit ditekan
            // Contoh: Redirect ke halaman edit dengan mengirim ID
            window.location.href = "/edit/" + id;
        });
    });


    //event untuk selesai
    $(document).ready(function () {
        // Event listener untuk tombol Edit
        $(document).on("click", ".b3", function () {
            var id = $(this).data("id");

            // Mengatur nilai ID ke input di dalam form atau span
            // $("#idSelesai").val(id); val untuk memberikan value ke input
            // $("#namaSelesai").text(nama); text untuk span
            // $("#idSelesai").text(id);
            // $("#namaSelesai").text(nama);

            // Menampilkan form
            // $("#formSelesai").show();

            // Lakukan apa yang perlu dilakukan saat tombol Edit ditekan
            // Contoh: Redirect ke halaman edit dengan mengirim ID
            window.location.href = "/form-selesai/" + id;
        });
    });



});
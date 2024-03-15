// Ambil parameter dari URL
const params = new URLSearchParams(window.location.search);
const status = params.get('status');

if (status === 'success') {
    alert('Data pendaki berhasil disimpan!');
}

else if (status === 'error') {
    alert('GAGAL!! Id Gelang sudah digunakan');
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

//untuk refresh page
function refreshPage() {
    location.reload();
}
document.getElementById("tombolRefresh").addEventListener("click", function () {
    refreshPage();
});

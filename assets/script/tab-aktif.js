
  // Mendapatkan semua elemen anchor di dalam topnav
  var navLinks = document.querySelectorAll(".topnav a");

  // Loop melalui setiap elemen anchor
  navLinks.forEach(function(navLink) {
    // Tambahkan event listener untuk setiap anchor
    navLink.addEventListener("click", function() {
      // Hapus kelas 'active' dari semua elemen anchor
      navLinks.forEach(function(link) {
        link.classList.remove("active");
      });

      // Tambahkan kelas 'active' pada elemen anchor yang sedang di-klik
      this.classList.add("active");
    });
  });


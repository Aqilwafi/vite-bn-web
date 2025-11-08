CREATE TABLE artikel (
    id SERIAL PRIMARY KEY,
    judul VARCHAR(255) NOT NULL,                
    slug VARCHAR(255) UNIQUE NOT NULL,          
    konten_md TEXT NOT NULL,                    
    ringkasan TEXT,                             
    gambar TEXT,                                
    kategori VARCHAR(100),                      
    penulis VARCHAR(100),                       
    waktu_baca INTEGER DEFAULT 1,               
    jumlah_komentar INTEGER DEFAULT 0,          
    unggulan BOOLEAN DEFAULT FALSE,             
    tanggal_dibuat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tanggal_diperbarui TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE komentar (
    id SERIAL PRIMARY KEY,
    artikel_id INT NOT NULL REFERENCES artikel(id) ON DELETE CASCADE,
    nama VARCHAR(100) DEFAULT 'Anonim',
    isi TEXT NOT NULL,
    tanggal_dibuat TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 1️⃣ Fungsi untuk menambah jumlah komentar
CREATE OR REPLACE FUNCTION tambah_jumlah_komentar()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE artikel
  SET jumlah_komentar = jumlah_komentar + 1,
      tanggal_diperbarui = CURRENT_TIMESTAMP
  WHERE id = NEW.artikel_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 2️⃣ Fungsi untuk mengurangi jumlah komentar
CREATE OR REPLACE FUNCTION kurangi_jumlah_komentar()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE artikel
  SET jumlah_komentar = GREATEST(jumlah_komentar - 1, 0),
      tanggal_diperbarui = CURRENT_TIMESTAMP
  WHERE id = OLD.artikel_id;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- 3️⃣ Trigger setelah komentar ditambahkan
CREATE TRIGGER komentar_tambah_trigger
AFTER INSERT ON komentar
FOR EACH ROW
EXECUTE FUNCTION tambah_jumlah_komentar();

-- 4️⃣ Trigger setelah komentar dihapus
CREATE TRIGGER komentar_hapus_trigger
AFTER DELETE ON komentar
FOR EACH ROW
EXECUTE FUNCTION kurangi_jumlah_komentar();


INSERT INTO artikel (
    judul,
    slug,
    ringkasan,
    konten_md,
    gambar,
    kategori,
    penulis
) VALUES
-- 1
(
    'Panggung Kartini, Panggung Ekspresi Siswa MI Al Maarif Ngadri',
    'panggung-kartini-panggung-ekspresi-siswa-mi-al-maarif-ngadri',
    'MI Al Maarif Ngadri merayakan Hari Kartini dengan berbagai lomba yang menumbuhkan semangat emansipasi siswa. Lebih dari 60 peserta ikut menampilkan bakat mereka dengan penuh semangat.',
    'Ngadri, 21 April 2025 - MI Al Maarif Ngadri merayakan Hari Kartini dengan menggelar lomba puisi, menyanyi, dan story telling pada 21 April 2025.

Kegiatan ini bertujuan menumbuhkan semangat emansipasi dan kreativitas siswa. Lebih dari 60 siswa ikut berpartisipasi, menampilkan bakat terbaik mereka dalam berbagai kategori.

“Peringatan ini tidak hanya mengenang Kartini, tetapi juga melatih keberanian dan kreativitas anak-anak,” ujar Usth Ita Kumala, koordinator kegiatan. 

Peringatan Hari Kartini kali ini tidak hanya seru, tetapi juga memberi ruang bagi siswa untuk berkembang, mengasah bakat, dan mengenal tokoh inspiratif bangsa.

#HariKartini2025  
#MIAlMaarifNgadri  
#KartiniMasaKini  
#LombaKartini',
    '/images/FLYR-BAITUNNAIM-02.jpg',
    'Acara Sekolah',
    'Admin'
),

-- 2
(
    'Gigi Sehat, Anak Ceria: Edukasi Kesehatan Gigi untuk Anak PAUD oleh Puskesmas Binangun',
    'gigi-sehat-anak-ceria-edukasi-kesehatan-gigi-untuk-anak-paud-oleh-puskesmas-binangun',
    'Kegiatan edukasi kesehatan gigi oleh Puskesmas Binangun disambut antusias anak-anak PAUD Plus Baitun Naim. Mereka belajar menjaga kebersihan mulut dengan cara yang menyenangkan dan interaktif.',
    'Rabu, 16 April 2025 - PAUD Plus Baitun Naim kedatangan tamu istimewa dari Puskesmas Binangun dalam kegiatan edukasi kesehatan gigi untuk anak-anak.

Dengan penuh semangat dan tawa, anak-anak belajar menjaga kebersihan mulut sejak dini. Kegiatan ini dikemas secara menarik dan interaktif, lengkap dengan peragaan dan praktik langsung bersama para petugas kesehatan.

Terima kasih Puskesmas Binangun atas kunjungannya! Semoga ilmu yang dibagikan hari ini bermanfaat dan menjadi bekal bagi anak-anak kami untuk tumbuh sehat dan ceria dengan senyum yang menawan 😊✨

#PAUDPlusBaitunNaim  
#SenyumSehatAnakHebat  
#PenyuluhanGigi  
#PuskesmasBinangun  
#AnakCeriaAnakSehat',
    '/images/FLYR-BAITUNNAIM-01.jpg',
    'Edukasi Kesehatan',
    'Admin'
),

-- 3
(
    'Market Day Penuh Kreativitas, Siswa-Siswi MI Al Maarif Ngadri Rayakan Hardiknas 2025 dengan Semangat Kewirausahaan',
    'market-day-mi-al-maarif-ngadri-hardiknas-2025',
    'Dalam rangka memperingati Hardiknas 2025, MI Al Maarif Ngadri menggelar Market Day penuh warna yang mengasah kreativitas dan jiwa kewirausahaan siswa-siswi.',
    'Dalam rangka memperingati Hari Pendidikan Nasional 2025, MI Al Maarif Ngadri menggelar acara Market Day yang penuh warna pada 2 Mei.

Acara ini menjadi ajang bagi siswa-siswi untuk menampilkan kreativitas mereka melalui berbagai produk hasil karya tangan, mulai dari makanan ringan hingga kerajinan tangan.

Selain sebagai bentuk perayaan, Market Day juga memberikan pengalaman berharga dalam berwirausaha, menggali potensi dan keterampilan bisnis para siswa-siswi.

Seluruh siswa-siswi dengan antusias mengisi stand mereka dengan berbagai barang dagangan yang mereka buat sendiri. Semangat kewirausahaan yang ditunjukkan tidak hanya mencerminkan keterampilan teknis, tetapi juga nilai-nilai inovasi dan kreativitas yang diterapkan dalam kehidupan sehari-hari.

Kegiatan ini bukan hanya sekadar ajang jual beli, tetapi juga menjadi sarana bagi siswa-siswi untuk belajar mengenai pentingnya bekerja keras, berkolaborasi, dan mengembangkan jiwa kepemimpinan.

Follow kami di media sosial untuk update terbaru:  
Instagram: @mi_almaarif.ngadri @paud_baitunnaim  
TikTok: @baitunnaim_official  
YouTube: @baitunnaim_official  

Hubungi Kami untuk informasi lebih lanjut:  
📞 Nomor HP/WA: 081364966677',
    '/images/FLYR-BAITUNNAIM-03.jpg',
    'Kegiatan Sekolah',
    'Admin'
),

-- 4
(
    'Peringatan Hari Pendidikan Nasional 2025 di MI Al Maarif Ngadri: Kegiatan yang Menggugah Semangat Pendidikan',
    'peringatan-hari-pendidikan-nasional-2025-mi-al-maarif-ngadri',
    'MI Al Maarif Ngadri merayakan Hardiknas 2025 dengan kegiatan jalan sehat dan market day yang menumbuhkan semangat belajar dan kebersamaan.',
    'MI Al Maarif Ngadri merayakan Hari Pendidikan Nasional 2025 pada 2 Mei dengan serangkaian kegiatan yang mengedepankan semangat kebersamaan dan penguatan nilai pendidikan.

Acara ini melibatkan seluruh elemen siswa-siswi dan guru dengan kegiatan seperti jalan sehat dan market day, yang membawa pesan tentang kerja keras, sportivitas, dan pentingnya pendidikan dalam kehidupan.

Semangat yang terbangun dalam acara ini diharapkan dapat terus menggugah para siswa-siswi untuk terus belajar, berinovasi, dan menjaga semangat kebersamaan dalam menjalani proses pendidikan.

Follow kami di media sosial untuk update terbaru:  
Instagram: @mi_almaarif.ngadri @paud_baitunnaim  
TikTok: @baitunnaim_official  
YouTube: @baitunnaim_official  

Hubungi Kami untuk informasi lebih lanjut:  
📞 Nomor HP/WA: 081364966677',
    '',
    'Pendidikan',
    'Admin'
),

-- 5
(
    'Asesmen Sumatif Akhir Jenjang (ASAJ) MI Al Maarif Ngadri: Menyongsong Masa Depan Pendidikan dengan Evaluasi Menyeluruh',
    'asaj-mi-al-maarif-ngadri-2025',
    'MI Al Maarif Ngadri menggelar ASAJ bagi siswa kelas VI sebagai evaluasi capaian belajar menjelang kelulusan tahun 2025.',
    'MI Al Maarif Ngadri menggelar Asesmen Sumatif Akhir Jenjang (ASAJ) bagi siswa kelas VI mulai 21 hingga 28 April 2025.

Kegiatan ini menjadi bagian penting dalam mengevaluasi capaian belajar selama enam tahun terakhir. Para siswa mengikuti asesmen dengan semangat dan persiapan matang, didampingi guru-guru yang terus memberikan motivasi.

Semoga seluruh proses berjalan lancar dan hasilnya menjadi bekal terbaik menuju jenjang selanjutnya.

#ASAJ2025  
#MIAlMaarifNgadri  
#PendidikanBermakna  
#UjianAkhirMadrasah  
#GenerasiBerkarakter',
    '',
    'Akademik',
    'Admin'
),

-- 6
(
    'Keceriaan PAUD Baitunnaim: Anak-anak Sambut Kembali Sekolah dengan Halal Bihalal Ceria Setelah Libur Idul Fitri!',
    'halal-bihalal-paud-baitunnaim-2025',
    'Anak-anak PAUD Baitunnaim kembali ke sekolah dengan keceriaan dan kegiatan halal bihalal yang penuh makna setelah libur Idul Fitri.',
    'PAUD Baitunnaim menyambut hari pertama masuk sekolah setelah libur Idul Fitri dengan penuh keceriaan.

Anak-anak terlihat antusias dan senang bertemu teman-teman serta guru-guru mereka. Untuk mempererat persaudaraan setelah liburan, PAUD Baitunnaim mengadakan acara halal bihalal yang menghangatkan suasana.

Anak-anak saling bermaaf-maafan dan berbagi cerita seru tentang liburan mereka, sembari menikmati kebersamaan yang menyenangkan.

Melalui kegiatan ini, anak-anak diajarkan tentang nilai-nilai persaudaraan, saling menghormati, dan menjaga hubungan baik sejak dini.

#PAUDBaitunnaim  
#KeceriaanAnak  
#HalalBihalalPAUD  
#KembaliKeSekolah  
#SemangatBaru  
#PersaudaraanPAUD  
#LiburLebaran  
#AnakHebatPAUD',
    '',
    'Acara PAUD',
    'Admin'
),

-- 7
(
    'Suasana Hangat dan Penuh Keceriaan, MI Al Maarif Ngadri Adakan Halal Bihalal Setelah Libur Lebaran',
    'halal-bihalal-mi-al-maarif-ngadri-2025',
    'Setelah libur panjang Idul Fitri, MI Al Maarif Ngadri kembali dengan semangat baru melalui kegiatan halal bihalal di sekolah.',
    'Hari ini, suasana di MI Al Maarif Ngadri sangat hangat dan penuh keceriaan.

Setelah menikmati libur panjang Idul Fitri, para siswa dan guru kembali ke sekolah dengan semangat yang tinggi.

Pada hari pertama masuk sekolah, Senin (14 April 2025), mereka mengadakan acara halal bihalal sebagai bentuk silaturahmi setelah liburan.

Dengan kegiatan halal bihalal ini, MI Al Maarif Ngadri berharap agar para siswa bisa memulai kembali hari-hari mereka di sekolah dengan semangat baru.

#MIAlMaarifNgadri  
#HalalBihalal  
#KembaliKeSekolah  
#SemangatBaru  
#SilaturahmiIdulFitri  
#HariPertamaSekolah  
#PersaudaraanSekolah  
#KeceriaanSetelahLebaran  
#SiswaBerprestasi  
#MIAlMaarifNgadriHebat',
    '/images/pagi.png',
    'Kegiatan Sekolah',
    'Admin'
),

-- 8
(
    'MI Al Maarif Ngadri Raih Juara Umum 2 Porseni Kecamatan Binangun 2025',
    'mi-al-maarif-ngadri-juara-umum-2-porseni-binangun-2025',
    'MI Al Maarif Ngadri sukses meraih juara umum 2 di ajang Porseni Kecamatan Binangun 2025 dengan total 32 piala dari berbagai cabang lomba.',
    'MI Almaarif Ngadri berhasil meraih prestasi gemilang dengan menjadi juara umum 2 pada ajang Porseni Kecamatan Binangun, yang dilaksanakan pada 12 April 2025.

Dalam kompetisi yang diikuti oleh berbagai sekolah se-Kecamatan Binangun, MI Almaarif sukses mengumpulkan total 32 piala, mencakup berbagai kategori olahraga dan seni.

Keberhasilan ini tak lepas dari dukungan penuh pihak sekolah, guru, dan para siswa yang tampil maksimal.

Dengan prestasi ini, MI Almaarif Ngadri semakin mengukuhkan namanya sebagai sekolah yang berkompeten di bidang pengembangan potensi siswa. 

#MIAlmaarifNgadri  
#JuaraUmumPorseni  
#PorseniBinangun2025  
#PrestasiGemilang  
#32Piala  
#OlahragaDanSeni  
#PendidikanBerkualitas  
#SemangatPelajar  
#PendidikanHebat  
#KreativitasSiswa',
    '',
    'Prestasi Sekolah',
    'Admin'
);


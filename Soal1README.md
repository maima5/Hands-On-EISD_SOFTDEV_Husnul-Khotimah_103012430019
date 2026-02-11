1. Soal noncoding (WAJIB)
EISD Lab baru saja merilis website pendaftaran Study Group untuk batch terbaru. Karena ingin cepat dipakai, versi awalnya dibuat sederhana: peserta tinggal isi Nama, Email, Pilihan Kelas, lalu klik Daftar.
Di hari pertama pendaftaran, panitia menemukan beberapa kejadian :
- Ada peserta yang panik karena halaman sempat loading lama, lalu menekan tombol Daftar berkali-kali yang mengakibatkan data masuk lebih dari sekali.
- Ada yang mengisi asal: nama “.”, email “aaa”, atau bahkan mengosongkan form mengakibatkan data berantakan.
- Di akhir hari, beberapa kelas yang kuotanya 25 orang ternyata terisi 30+ karena sistem tetap menerima pendaftaran.
- Saat panitia mau rekap, muncul pertanyaan: “Kalau data dobel dan data tidak valid begini, mana yang dianggap resmi?”
Anggap kamu menjadi orang yang diminta membantu memikirkan perbaikan versi berikutnya, tapi waktu sangat terbatas dan panitia minta kamu menentukan prioritas. 
Jika kamu hanya boleh memperbaiki satu hal pada website tersebut, apa yang akan kamu perbaiki terlebih dahulu dan mengapa?

Jawaban:
Saya akan memperbaiki BACKEND VALIDATION
Saya akan memperbaiki server dengan menambahkan pengecekan logika di server sebelum data yang diinput user disimpan ke dalam database.
1. Klik berkali kali akan terjadi duplikat data, maka diterapkan pengecekan email, bila email sudah terdaftar maka respon ditolak.
2. Input asal-asalan maka backend akan diterapkan pengecekan validasi yaitu; nama tidak kosong, email valid, terdapat pilihan kelas.
3. Respon melebihi kuota yang terdaftar maka diterapkan logika pengecekan, spt ini:
if jumlahTerdaftar >= kuota {
    tolak pendaftaran
}
tolak pendaftaran dapat diisi dengan error message atau alternative lainnya terkait problem tsb
4. Jika ketiga hal diatas diterapkan, maka hal ke-4 tidak akan terjadi karena data yang berhasil diinput adalah dataresmi yakni data yang unik (tidak dobel/duplikat) dan datanya valid.

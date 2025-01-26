Penjelasan Singkat Kode

Fungsi ReverseString:
-Fungsi ini menerima input berupa string.
-String diubah menjadi slice of rune agar mendukung semua karakter, termasuk Unicode (misalnya: emoji, aksara asing).
-Slice ini dibalik dengan menukar posisi karakter dari awal dan akhir menggunakan teknik dua pointer.
-Hasil akhirnya adalah slice yang sudah terbalik, kemudian diubah kembali menjadi string.

Fungsi main:
-Fungsi main adalah titik awal program.
-Program memberikan contoh string "Halo Dunia!".
-String ini dibalik menggunakan fungsi ReverseString, kemudian hasilnya dicetak ke layar.

Contoh Penggunaan
Misalnya, input string-nya adalah "Halo Dunia!", hasil yang akan ditampilkan adalah:

```yaml
String asli: Halo Dunia!
String terbalik: !ainuD olaH
```

Cara Menjalankan Program
Siapkan lingkungan Go:

Pastikan Go sudah terpasang di komputer. Kamu bisa mengunduhnya dari situs resmi Golang.
Periksa instalasi dengan mengetikkan go version di terminal/command prompt.
Buat file baru:

Simpan kode di atas dalam file bernama main.go.
Jalankan program:

Buka terminal/command prompt, navigasikan ke folder tempat file main.go berada.
Ketik perintah berikut untuk menjalankan program

```go
go run main.go
```

Lihat output:
-Output dari program akan terlihat seperti ini:

```yaml
String asli: Halo Dunia!
String terbalik: !ainuD olaH
```

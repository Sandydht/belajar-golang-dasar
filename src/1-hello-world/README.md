# Main Function
- Go-Lang, itu mirip seperti bahasa pemrograman C/C++, dimana perlu ada yang namanya main function.
- Main function adalah sebuah fungsi yang akan dijalankan ketika program berjalan.
- Untuk membuat function, kita bisa menggunakan kata kunci ```func```.
- Main function harus terdapat didalam main package.
- Titik koma di Golang, tidaklah wajib, artinya kita bisa menambahkan titik koma atau tidak diakhir kode program kita.

## Println
- Untuk menulis tulisan, kita perlu melakukan import module ```fmt``` terlebih dahulu.
- Mirip ketika kita belajar Java.
- Materi tentang import, akan kita bahas di bagian tersendiri.

# Multiple Main Function
- Di Golang, function dalam module / project adalah unik, artinya kita tidak boleh membuat nama function yang sama.
- Oleh karena itu, jika kita membuat file baru, misal sample.go, lalu membuat nama function yang sama yaitu main.
- Maka kita tidak bisa melakukan build module, karena main function tersebut duplikat dengan yang ada di main function hello-world.go.

## Solusinya ?
- Karena sekarang kita masih dalam fase belajar, oleh karena itu kita tidak akan melakukan build project module terlebih dahulu.
- Sekarang kita akan fokus menjalankan file Golang satu persatu, sehingga tidak akan terjadi error.
- Tapi INGAT, pada kenyataannya nanti, saat kita membuat project, kita hanya akan membuat satu main function saja.

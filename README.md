# Indonesian Numeral Spellers

## Latar Belakang
Mengeja angka merupakan salah satu kegiatan dasar yang dilakukan setiap harinya. Contoh kegiatan tersebut yaitu mengeja harga barang, nilai data, tanggal dan tahun, serta masih banyak lagi. Meskipun terkesan hal sepele, berdasarkan penelitian dari para dokter di Indonesia, seorang anak baru bisa membaca dan mengeja angka pada umur 4-6 tahun. Rentang usia tersebut tentunya terasa kurang cepat. Padahal, semakin cepat seorang bisa membaca dan mengeja angka, maka semakin cepat pula anak tersebut dapat belajar berhitung dan mempelajari hal-hal lainnya, bahkan termasuk belajar pemrograman.

Dari permasalah di atas, maka diperlukanlah suatu sarana pembelajaran yang dapat membantu anak-anak balita di Indonesia untuk membaca dan mengeja angka. Dengan adanya solusi tersebut, diharapkan anak-anak dapat membaca dan mengeja angka lebih cepat sehingga mampu segera mempelajari hal-hal lebih besar lainnya dan tentunya meningkatkan tingkat pendidikan di Indonesia.

## Spesifikasi
Buatlah dalam bahasa pemrograman **_Go_**, sebuah web service berupa **_REST API_**, yang dapat mengeja (dalam bahasa Indonesia) dari angka yang diberikan serta menuliskan angka yang tepat dari masukkan ejaan angka (dalam bahasa Indonesia juga), dengan ketentuan-ketentuan sebagai berikut :

1. Terdapat 2 buah endpoint API yang perlu dibuat, yaitu '**GET** /spell' yang menerima parameter angka, serta '**POST** /read' yang menerima body/payload berupa text/ejaan. Jika input parameter atau body/payload tidak valid, maka berikan response keterangan error/gagal dengan format dibebaskan.

2. Sebagai REST API, maka response harus berupa JSON. Struktur data response JSON dibebaskan.

3. Program dibuat dengan mengikuti standar development resmi Go (lihat referensi #3), yaitu environment kode program berada pada ```$GOPATH/src/```, misalkan ```$GOPATH/src/github.com/Indonesian-Numeral-Spellers```.

4. Arsitektur program dibebaskan (boleh mengikuti referensi-referensi _REST API with Go_ dari internet), namun harus tetap tersusun dengan rapi dan mengerti apa kegunaan setiap fungsi, file, serta package.

5. Batasan kasus uji : 2000000000 (dua milyar)

## Cara Penggunaan
1. Lakukan command run pada file go
```
go run spell.go read.go main.go
```
2. Uji coba API dengan menggunakan perangkat lunak [POSTMAN](https://www.getpostman.com/)

## Contoh Kasus Uji
### Contoh Kasus Uji 1 : Pengejaan
Request :
```
GET '/spell?number=123456'
```
Response :
![1](https://user-images.githubusercontent.com/38171936/57979914-06b3df80-7a4e-11e9-9087-9302117600f5.JPG)

### Contoh Kasus Uji 2 : Pembacaan
Request:
```JSON
POST '/read'
{
    "text" : "seribu sembilan ratus sembilan puluh tujuh"
}
```
Response :
![2](https://user-images.githubusercontent.com/38171936/57979915-0f0c1a80-7a4e-11e9-9899-038a87f43709.JPG)

## Referensi Pengerjaan
1. https://medium.freecodecamp.org/portfolio-app-using-react-618814e35843
2. https://hakaselogs.me/2018-04-20/building-a-web-app-with-go-gin-and-react/

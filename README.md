# Indonesian Numeral Spellers

## Latar Belakang
Mengeja angka merupakan salah satu kegiatan dasar yang dilakukan setiap harinya. Contoh kegiatan tersebut yaitu mengeja harga barang, nilai data, tanggal dan tahun, serta masih banyak lagi. Meskipun terkesan hal sepele, berdasarkan penelitian dari para dokter di Indonesia, seorang anak baru bisa membaca dan mengeja angka pada umur 4-6 tahun. Rentang usia tersebut tentunya terasa kurang cepat. Padahal, semakin cepat seorang bisa membaca dan mengeja angka, maka semakin cepat pula anak tersebut dapat belajar berhitung dan mempelajari hal-hal lainnya, bahkan termasuk belajar pemrograman.

Dari permasalah di atas, maka diperlukanlah suatu sarana pembelajaran yang dapat membantu anak-anak balita di Indonesia untuk membaca dan mengeja angka. Dengan adanya solusi tersebut, diharapkan anak-anak dapat membaca dan mengeja angka lebih cepat sehingga mampu segera mempelajari hal-hal lebih besar lainnya dan tentunya meningkatkan tingkat pendidikan di Indonesia.

## Tech Stack yang Digunakan
1. Go, dengan penggunaan library **gin** dan **cors**
2. React.JS dengan menggunakan library **axios** 

## Cara Penggunaan
1. Masuk ke $GOPATH/src, kemudian clone repository ini, dan jalankan :
```
go get github.com/Indonesian-Numeral-Spellers
```

2. Untuk menguji coba API, menggunakan perangkat lunak [POSTMAN](https://www.getpostman.com/) dengan melakukan request pada :
```
http://localhost:8080/
```

3. Jika ingin menggunakan UI, masuk ke direktori **app**, kemudian ketik perintah :

```
npm install
npm start
```

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

### Contoh Kasus Uji 3 : Pengejaan dan Pembacaan Angka Negatif
Request :
```
GET '/spell?number=-3'
```
```JSON
POST '/read'
{
    "text" : "negatif dua ribu"
}
```
Response :
![Capture](https://user-images.githubusercontent.com/38171936/57989049-acebfd80-7abf-11e9-9aeb-cc448c577dee.JPG)

### Contoh Kasus Uji 4 : Pengejaan dan Pembacaan dengan Masukan Tidak Valid
Request :
```
GET '/spell?number=3-'
```
```JSON
POST '/read'
{
    "text" : "dua ribu puluh"
}
```
Response :
![2](https://user-images.githubusercontent.com/38171936/57989062-e45aaa00-7abf-11e9-84d7-b9712f44c1ba.JPG)

## Referensi Pengerjaan
1. https://medium.freecodecamp.org/portfolio-app-using-react-618814e35843
2. https://hakaselogs.me/2018-04-20/building-a-web-app-with-go-gin-and-react/
3. https://reactjs.org/docs/forms.html
4. https://alligator.io/react/axios-react/
5. https://dasarpemrogramangolang.novalagung.com/C-12-cors-preflight-request.html
6. https://github.com/itsjamie/gin-cors

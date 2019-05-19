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

## Algoritma dan Proses Pengerjaan
Pada pembuatan program ini, proses pengerjaan dilakukan dengan dua tahap, yaitu tahap pembuatan _API/backend_ dengan menggunakan bahasa **Go** dan tahap pembuatan _User Interface/frontend_ dengan menggunakan **React.JS**

### API
1. Membuat masing-masing fungsi untuk _endpoint_, yaitu fungsi **spell** dan fungsi **read** beserta fungsi-fungsi utilitasnya, dengan cara membaca setiap **tiga** angka dikarenakan bahasa Indonesia yang baku hanya mengeja maksimal tiga angka untuk menghasilkan angka yang valid
2. Membuat fungsi untuk **validasi** masukan, untuk validasi angka hanya membutuhkan validasi apakah angka kurang dari sama dengan _dua milyar(2000000000)_, sedangkan untuk validasi teks harus dilakukan validasi _beruntun_ dikarenakan banyak kemungkinan masukan.

### User Interface
1. Membuat _desain_ dari user interface, dengan cara membuat terlebih dahulu template dalam bentuk **HTML** 
2. Mengubah template dari yang sebelumnya dalam bentuk HTML, menjadi bentuk **JSX** dengan menggunakan [HTMLtoJSXconverter](https://magic.reactjs.net/htmltojsx.htm)
3. Membuat sebuah **kelas form** sesuai aturan _React.js_ , kemudian membuat fungsi-fungsi yang akan dipanggil ketika mendapatkan masukan dari **form**
4. Dengan menggunakan **axios**, akan dihubungkan API dengan User Interface, dengan User Interface akan mengirimkan nilai dalam bentuk **teks/angka**, dan mendapatkan hasil dalam bentuk **JSON**

## Tech Stack yang Digunakan
1. Go, dengan penggunaan library **gin** dan **cors**
2. React.JS dengan menggunakan library **axios** 

## Cara Penggunaan
1. Masuk ke direktori **backend**, kemudian ketik perintah :

```
go run spell.go read.go main.go
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

## Referensi Pengerjaan
1. https://medium.freecodecamp.org/portfolio-app-using-react-618814e35843
2. https://hakaselogs.me/2018-04-20/building-a-web-app-with-go-gin-and-react/
3. https://reactjs.org/docs/forms.html
4. https://alligator.io/react/axios-react/

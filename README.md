## Tubes3_13520074
# **Tubes Stima 3 by Rucika Wavin**

Tugas Besar III IF2211 Strategi Algoritma Semester II Tahun 2020

**Penerapan _String Matching_ dan _Regular Expression_ dalam DNA _Pattern Matching_**

## Table of Contents
* [Introduction](#introduction)
* [General Information](#general-information)
* [Technologies Used](#technologies-used)
* [Installation and Requirement](#installation-and-requirement)
* [How to Run](#how-to-run)
* [How to Use](#how-to-use)
* [Project Status](#project-status)
* [Room for Improvement](#room-for-improvement)
* [Acknowledgements](#acknowledgements)
* [Contact](#contact)


## Introduction
Hai, Selamat datang di Repository Github kami!

Proyek ini dibuat oleh:
| No. | Nama | NIM |
| :---: | :---: | :---: |
| 1. | Eiffel Aqila Amarendra | 13520074 |
| 2. | Firizky Ardiansyah | 13520095 |
| 3. | Ilham Bintang Nurmansyah | 13520102 |


## General Information

Dalam tugas besar ini, kelompok kami membangun sebuah aplikasi web interaktif
sederhana untuk mendeteksi dan memprediksi keberadaan penyakit genetik tertentu pada seorang pengguna berdasarkan _sequence_ DNA-nya.

Pendeteksian tersebut diimplementasikan dengan menggunakan algoritma Knuth Morris Pratt (KMP),Boyer-Moore (BM), serta Regular Expression (Regex). Selain pendeteksian penyakit, website kami juga memiliki fitur untuk menginput data penyakit baru serta pencarian hasil-hasil prediksi.


## Technologies Used
Implementasi program ditulis dalam bahasa HTML, CSS, JS, Go dengan framework Vue JS.


## Installation and Requirement
- Unduh seluruh folder dan file pada repository ini atau clone repository
- Unduh dan pasang [Go](https://go.dev/doc/install)
- Unduh dan pasang [Node JS](https://nodejs.org/en/download/)


## How to Run
- Pastikan semua requirement di atas sudah terpasang (install) pada perangkat keras yang akan digunakan
- Buka direktori Frontend (rucikawavin) menggunakan perintah berikut `cd .\src\Frontend\rucikawavin\`
- Kemudian, jalankan perintah `npm install`
- Setelah proses instalasi selesai, jalankan perintah `npm run build`
- Kemudian, setelah proses build selesai, buka direktori dist hasil build pada langkah sebelumnya. Folder dist akan berisi `index.html` yang perlu di-_edit_ agar bisa di-_load_ oleh Backend.
- Cari tag bernama `script` dengan parameter `type = "module"`. Ubah module pada kedua sisi tag menjadi `type = "text/javascript"`. Simpan hasil perubahan.
- Buka direktori backend pada terminal dengan menggunakan perintah berikut `cd ..\..\Backend\`
- Lakukan build pada backend menggunakan perintah `go build`
- Setelah proses build selesai, jalankan `go run main.go`
- Kemudian, buka browser dan buka halaman berikut `localhost:8080`


## Project Status
Project status: _complete_

Seluruh spesifikasi telah dipenuhi. Selain itu, kami juga mengerjakan bonus, yakni tingkat kemiripan DNA pengguna.


## Room for Improvement
- Sebuah algoritma program yang dapat membuat program berjalan lebih cepat dan efisien


## Acknowledgements
- Project ini dibuat berdasarkan [Spesifikasi Tugas Besar 3 Stima](https://informatika.stei.itb.ac.id/~rinaldi.munir/Stmik/2021-2022/Tugas-Besar-3-IF2211-Strategi-Algoritma-2022.pdf);
- Terima kasih kepada Tuhan Yang Maha Esa;
- Terima kasih kepada Ibu Masayu Leylia Khodra, Ibu Nur Ulfa Maulidevi, dan Pak Rinaldi sebagai dosen kami;
- Terima kasih kepada asisten;
- Project ini dibuat untuk memenuhi Tugas Besar 3 IF2211 Strategi Algoritma.


## Contact
Created by Rucika Wavin. 2022 All Rights Reserved.
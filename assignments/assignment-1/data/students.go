package data

import "assignment-1/domain"

var Students = []domain.Student{
	{
		Id:        0,
		Nama:      "Taufik",
		Alamat:    "Bandung",
		Pekerjaan: "Mahasiswa",
		Alasan:    "Karena saya ingin meningkatkan pengetahuan saya mengenai back-end, khususnya golang",
	},
	{
		Id:        1,
		Nama:      "Agus",
		Alamat:    "Jakarta Selatan",
		Pekerjaan: "Software Engineer",
		Alasan:    "Karena ingin mempelajari bahasa pemrograman baru",
	},
	{
		Id:        2,
		Nama:      "Santoso",
		Alamat:    "Surabaya",
		Pekerjaan: "Web Developer",
		Alasan:    "Karena ingin memperdalam pengetahuan dan keterampilan dalam pengembangan aplikasi web yang berbasis Go",
	},
	{
		Id:        3,
		Nama:      "Budi",
		Alamat:    "Tangerang Selatan",
		Pekerjaan: "Data Analyst",
		Alasan:    "Karena ingin mempelajari bahasa pemrograman baru yang lebih efisien dalam menangani analisis data besar dan kompleks",
	},
	{
		Id:        4,
		Nama:      "Siti",
		Alamat:    "Bandung",
		Pekerjaan: "Backend Developer",
		Alasan:    "Merasa tertarik dengan fitur-fitur concurrency yang dimiliki oleh Golang dan ingin mempelajarinya lebih dalam",
	},
}

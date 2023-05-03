package biodata

type DataDiri struct {
	Nama   string
	NIM    int
	Alamat string
}

func ShowDataDiri(nama string, NIM int, Alamat string) DataDiri {
	Profil := DataDiri{
		Nama:   nama,
		NIM:    NIM,
		Alamat: Alamat,
	}
	return Profil
}

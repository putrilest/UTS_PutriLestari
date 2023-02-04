package models

type Kategori struct {
	Id_ktg   int    `gorm:"primaryKey;autoIncrement" json:"id_ktg"`
	Nama_ktg string `json:"nama_ktg" binding:"required"`
	Menu     []Menu `gorm:"foreignKey:Id_ktg;" json:"menu"`
}

type Menu struct {
	Id_menu   int    `gorm:"primaryKey;autoIncrement" json:"id_menu"`
	Id_ktg    int    `json:"id_ktg" binding:"required"`
	Nama_menu string `json:"nama_menu" binding:"required"`
	Harga     int    `json:"harga" binding:"required"`
}

type Meja struct {
	Id_meja int    `gorm:"primaryKey;autoIncrement" json:"id_meja"`
	No_meja int    `json:"no_meja"`
	Status  string `json:"status"`
}

type Customer struct {
	Id_cust   int         `gorm:"primaryKey;autoIncrement" json:"id_cust"`
	Nama_cust string      `json:"nama_cust" binding:"required"`
	Alamat    string      `json:"alamat" binding:"required"`
	No_hp     string      `json:"no_hp" binding:"required"`
	Reservasi []Reservasi `gorm:"foreignKey:Id_cust;" json:"reservasi"`
}

type Reservasi struct {
	Id_reservasi int    `gorm:"primaryKey;autoIncrement" json:"id_reservasi"`
	Id_cust      int    `json:"id_cust"`
	Id_meja      int    `json:"id_meja"`
	Tanggal      string `json:"tanggal"`
	Meja         Meja   `gorm:"foreignKey:Id_meja;references:Id_meja;" json:"meja"`
}

type Kasir struct {
	Id_kasir     int    `gorm:"primaryKey;autoIncrement" json:"id_kasir"`
	Id_reservasi int    `json:"id_reservasi" binding:"required"`
	Id_menu      int    `json:"id_menu" binding:"required"`
	Jumlah       int    `json:"jumlah" binding:"required"`
	Total        int    `json:"total"`
	Reservasi Reservasi `gorm:"foreignKey:Id_reservasi;references:Id_reservasi" json:"reservasi"`
	Menu         []Menu `gorm:"foreignKey:Id_menu;" json:"menu"`
}

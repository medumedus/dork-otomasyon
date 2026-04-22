package models

type Dork struct {
	ID       int    `json:"id"`
	Category string `json:"kategori"`
	Title    string `json:"baslik"`
	Desc     string `json:"aciklama"`
	Template string `json:"sorgu"`
}

type GeneratedDork struct {
	ID        int
	Category  string
	Title     string
	Desc      string
	Query     string
	GoogleURL string
}

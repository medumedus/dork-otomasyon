#  Otomatik Google Dork Aracı

## Proje Hakkında
Bu proje, Siber Tehdit İstihbaratı (CTI) süreçlerinde bilgi toplama aşamasını hızlandırmak için geliştirilmiş, Go (Golang) tabanlı otomatik Google Dork aracıdır. Sistem, kullanıcı tarafından girilen hedef alan adını veya kelimeyi alarak, sisteme önceden belirtilmiş 40 farklı siber güvenlik dork sorgusunu dinamik olarak oluşturur ve tarayıcı üzerinden çalıştırılmasını sağlar.

## Özellikler
* Girilen hedefe göre `{target}` değişkenini kullanarak otomatik dork oluşturma.
* Üretilen sorguları panoya kopyalama ve doğrudan Google üzerinde aratma.
* Servis, repository ve handler katmanlarına ayrılmış temiz mimari yapı.

## Kullanılan Teknolojiler
* **Backend:** Go (Golang 1.25)
* **Frontend:** HTML5, CSS3, Vanilla JavaScript
* **Veri Depolama:** JSON (`data/dorks.json`)

## Kurulum ve Kullanım
Projeyi kendi yerel bilgisayarınızda (localhost) çalıştırmak için aşağıdaki adımları izleyebilirsiniz.

**1. Repoyu bilgisayarınıza klonlayın:**
````bash
git clone https://github.com/medumedus/dork-otomasyon.git
````
**2.Proje dizinine giderek, uygulamayı derleyip çalıştırın:**
````bash
cd dork-otomasyon
go run main.go
````
**3.Tarayıcınıza aşağıdaki URL'yi girin:**
````bash
http://localhost:8080/
````
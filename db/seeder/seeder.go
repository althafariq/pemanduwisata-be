package seeder

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func Seed(db *sql.DB) {
	adminhashedPassword, _ := bcrypt.GenerateFromPassword([]byte("tgadmin123"), bcrypt.DefaultCost)
	userhashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	user, err := db.Exec(`INSERT INTO user (firstname, lastname, email, password, role) 
	VALUES 
	('Admin', 'Tour Guide', 'admin@email.com', ?, 'admin'),
	('Althaf', 'Ariq', 'althafariq@gmail.com', ?, 'user');`, adminhashedPassword, userhashedPassword)
	if err != nil {
		panic(err)
	}
	_, err = user.LastInsertId()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO destinations (id, name, location, description, budaya_name, budaya_description, photo_path)
	VALUES
	('1001', 'Gunung Bromo', 'Kab. Pasuruan, Jawa Timur', 'Gunung Bromo adalah salah satu gunung api yang masih aktif di Indonesia. Gunung yang memiliki ketinggian 2.392 meter di atas permukaan laut ini merupakan destinasi andalan Jawa Timur. Gunung Bromo berdiri gagah dikelilingi kaldera atau lautan pasir seluas 10 kilometer persegi. 

	Wisatawan yang berkunjung ke Gunung Bromo akan disambut oleh pemandangan yang indah. Salah satu hal yang terkenal dari Gunung Bromo adalah golden sunrise-nya, pasalnya, Gunung Bromo dinobatkan sebagai tempat yang menawarkan pemandangan matahari terbit terbaik di Pulau Jawa.Sesaat setelah momen matahari terbit berakhir, wisatawan akan kembali disuguhkan pemandangan yang tak kalah indanya, yaitu pemandangan negeri di atas awan.
	
	Sebagai daerah vulkanik terbesar di provinsi Jawa Timur, Taman Nasional Bromo Tengger Semeru memiliki wilayah seluas 800 km persegi. Destinasi pariwisata yang satu ini jangan sampai terlewatkan, terutama untuk Sobat Pesona yang tertarik pada aktivitas vulkanik. Sobat Pesona dapat menyaksikan asap dan abu yang berasal dari Gunung Semeru, sebuah gunung berapi yang masih aktif dengan ketinggian 3.676 meter di atas permukaan laut.
	
	Taman Nasional Bromo Tengger Semeru merupakan satu-satunya kawasan konservasi di Indonesia yang memiliki lautan pasir seluas 10 km yang disebut Tengger, tempat dimana empat anak gunung berapi baru berada. Anak gunung berapi tersebut adalah Gunung Batok (2.470 m), Gunung Kursi (2.581 m), Gunung Watangan (2.661 m), dan Gunung Widodaren (2.650 m). Namun, dari deretan gunung tersebut, hanya Gunung Bromo lah satu-satunya yang masih aktif. Temperatur di puncak Gunung Bromo berkisar 5-18 derajat Celcius. Bila Sobat Pesona menuju ke arah Selatan Taman, Sobat akan menemukan dataran terjal yang terbelah oleh lembah dan dihiasi dengan danau-danau yang indah hingga mencapai kaki Gunung Semeru.', 'Suku Tengger', 'Suku Tengger merupakan masyarakat yang berasal dari dataran tinggi Bromo-Tengger-Semeru. Mereka juga biasa disebut orang Tengger. Penduduknya menempati wilayah Kabupaten Pasuruan, Lumajang, Probolinggo, dan Malang.
	Bagi orang Tengger, Gunung Bromo atau yang juga disebut Gunung Brahma diyakini sebagai gunung suci. Setiap tahun, masyarakat Tengger mengadakan sebuah upacara Yadnya Kasada di bawah kaki Gunung Bromo, tepatnya di Pura Luhur Poten Bromo dan dilanjutkan ke puncak gunung.', 'media/destination/gunungbromo.png'),

	('1002', 'Desa Wisata Adat Osing Kemiren', 'Kab. Banyuwangi, Jawa Timur', 'Desa Wisata Adat Osing Kemiren terletak di Kecamatan Glagah, Kabupaten Banyuwangi, memiliki luas 177.052 Ha dengan penduduk 2.569 jiwa. Desa Adat Osing Kemiren berasal dari nama kemirian, atau banyak pohon kemiri. dan mayoritas masyarakat adalah suku osing yang merupakan suku asli kabupaten Banyuwangi. Desa Kemiren juga menjadi bagian dari kawasan Ijen Geopark sebagai culture site.

	Kemiren memiliki budaya yang beraneka ragam. mulai dari adat istiadat, bahasa, manuskrip, kesenian, tradisi lisan, ritus, pengetahuan, teknologi dan permainan Tradisional.
	
	Wisatawan tak perlu hawatir ketika berkunjung ke kemiren. Karena Ada homestay yang siap sebagai tempat singgah. Homestay dengan arsitektur osing dan keramahan warganya membuat nyaman terasa di kampung sendiri', 'Suku Osing', 'Suku Osing merupakan suku asli yang berasal dari Kabupaten Banyuwangi. Suku Osing memiliki budaya yang beraneka ragam. Mulai dari adat istiadat, bahasa, manuskrip, kesenian, tradisi lisan, ritus, pengetahuan, teknologi dan permainan Tradisional.', 'media/destination/kemiren.png')
	;`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO reviews (user_id, destination_id, rating, review)
	VALUES
	('2', '1001', '5', 'Gunung Bromo, pemandangan indah, dan suasana sejuk. Saya sangat suka dengan tempat ini. Saya sangat merekomendasikan tempat ini untuk anda yang ingin berlibur.'),
	('2', '1001', '4', 'Gunung Bromo 2, pemandangan indah, dan suasana sejuk. Saya sangat suka dengan tempat ini. Saya sangat merekomendasikan tempat ini untuk anda yang ingin berlibur.'),
	('2', '1002', '4', 'Desa Wisata, pemandangan indah, dan suasana sejuk. Saya sangat suka dengan tempat ini. Saya sangat merekomendasikan tempat ini untuk anda yang ingin berlibur.')
	;`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO TelpDarurat (name, number)
	VALUES
	('Police', '110'),
	('Ambulance', '112'),
	('Fire', '112')
	;`)

	if err != nil {
		panic(err)
	}
}

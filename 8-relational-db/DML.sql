-- memasukkan data ke sebuah table
INSERT INTO jurusan (nama)
VALUES ("Agroteknologi");

-- menampilkan data dari table
SELECT jur.id, jur.nama FROM jurusan jur;

-- no 1 a 
UPDATE jurusan SET 
jurusan.nama = "kedokteran"
WHERE jurusan.id = 7;

-- no 1 b
DELETE FROM jurusan WHERE id = 5;


INSERT INTO mahasiswa (id, nama, jenis_kelamin, alamat, jurusan_id, telp)
values (1001, "Adi", "laki-laki", "Jakarta", 1, "0812345");

INSERT INTO mahasiswa (id, nama, jenis_kelamin, alamat, jurusan_id, telp)
values (1002, "Budi", "laki-laki", "Surabaya", 7, "0812345");

INSERT INTO mahasiswa (id, nama, jenis_kelamin, alamat, jurusan_id, telp)
values (1003, "Cindy", "perempuan", "Medan", 1, "0812346");

SELECT * FROM mahasiswa m WHERE m.jurusan_id =7;

SELECT * FROM jurusan j WHERE j.nama LIKE "%Tek%";
SELECT * FROM jurusan j WHERE j.nama not LIKE "%Tek%";
SELECT * FROM jurusan j WHERE j.nama LIKE "%Tek%" OR j.nama LIKE "%Computer%";

select * from mahasiswa where nama like "%d%";

SELECT * FROM jurusan j WHERE j.id BETWEEN 1 and 5;

SELECT * FROM jurusan j WHERE j.nama LIKE "%Tek%" OR j.id BETWEEN 1 and 3;

SELECT * FROM jurusan j WHERE j.nama LIKE "%Tek%" AND j.id BETWEEN 1 and 3;


SELECT * FROM mahasiswa m;
SELECT * FROM mahasiswa m ORDER BY m.jurusan_id ASC ;
SELECT * FROM mahasiswa m ORDER BY m.jurusan_id DESC ;

SELECT * FROM mahasiswa m ORDER BY m.jurusan_id ASC LIMIT 2;




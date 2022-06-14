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
values (1005, "Budi 2", "laki-laki", "Surabaya", 4, "0812345");

INSERT INTO mahasiswa (id, nama, jenis_kelamin, alamat, jurusan_id, telp)
values (1003, "Cindy", "perempuan", "Medan", 1, "0812346");

INSERT INTO mahasiswa (id, nama, jenis_kelamin, alamat, jurusan_id, telp)
values (1013, "Budi 6", "laki-laki", "Surabaya", 9, "0812345");

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

DELETE FROM jurusan WHERE jurusan.id = 1;
DELETE FROM mahasiswa WHERE mahasiswa.id = 1002;

DELETE FROM jurusan WHERE jurusan.id = 4;


INSERT into mata_kuliah (id, nama, sks)
values (1, "pemrograman dasar", 4);
INSERT into mata_kuliah (id, nama, sks)
values (2, "pemrograman lanjut", 4);
INSERT into mata_kuliah (id, nama, sks)
values (3, "database", 3);
INSERT into mata_kuliah (id, nama, sks)
values (4, "API", 5);

SELECT * FROM mata_kuliah mk ;

INSERT into data_semester (id, mahasiswa_id, matakuliah_id, semester, tahun_ajaran)
values (1, 1001, 1, 1, "2022/2023");
INSERT into data_semester (id, mahasiswa_id, matakuliah_id, semester, tahun_ajaran)
values (2, 1001, 3, 1, "2022/2023");

INSERT into data_semester (id, mahasiswa_id, matakuliah_id, semester, tahun_ajaran)
values (3, 1003, 1, 1, "2022/2023");
INSERT into data_semester (id, mahasiswa_id, matakuliah_id, semester, tahun_ajaran)
values (4, 1003, 2, 2, "2023/2024");

INSERT into data_semester (id, mahasiswa_id, matakuliah_id, semester, tahun_ajaran)
values (5, 1006, 1, 1, "2022/2023");

SELECT * FROM data_semester ds ;


-- id # mahasiswa_id # nama_mahasiswa # matakuliah_id # nama_matkul # semester # tahun_ajaran
-- data_semester, mata_kuliah, mahasiswa
SELECT ds.id, ds.mahasiswa_id, mhs.nama as nama_mahasiswa, ds.matakuliah_id, mk.nama as nama_matkul, ds.semester, ds.tahun_ajaran 
from data_semester ds
INNER JOIN mata_kuliah mk ON ds.matakuliah_id = mk.id 
INNER JOIN mahasiswa mhs ON ds.mahasiswa_id = mhs.id;

-- +jurusan
SELECT ds.id, ds.mahasiswa_id, mhs.nama as nama_mahasiswa, jrs.nama as nama_jurusan, ds.matakuliah_id, mk.nama as nama_matkul, ds.semester, ds.tahun_ajaran 
from data_semester ds
INNER JOIN mata_kuliah mk ON ds.matakuliah_id = mk.id 
INNER JOIN mahasiswa mhs ON ds.mahasiswa_id = mhs.id
INNER JOIN jurusan jrs ON mhs.jurusan_id = jrs.id ;


--  union
select mk.nama FROM mata_kuliah mk
UNION
SELECT j.nama FROM jurusan j ;


select m.nama, mk.sks as jumlah_sks from mahasiswa m 
inner join data_semester ds ON ds.mahasiswa_id = m.id 
inner join mata_kuliah mk ON ds.matakuliah_id = mk.id 
WHERE m.id =1003 ;
 -- agregasi SUM
select m.nama, SUM(mk.sks) as jumlah_sks from mahasiswa m 
inner join data_semester ds ON ds.mahasiswa_id = m.id 
inner join mata_kuliah mk ON ds.matakuliah_id = mk.id 
WHERE m.id =1003
GROUP BY m.id ;

select m.nama, SUM(mk.sks) as jumlah_sks from mahasiswa m 
inner join data_semester ds ON ds.mahasiswa_id = m.id 
inner join mata_kuliah mk ON ds.matakuliah_id = mk.id 
GROUP BY m.id 
ORDER BY SUM(mk.sks) DESC;

-- count
SELECT m.nama, COUNT(ds.mahasiswa_id) from mahasiswa m 
INNER JOIN data_semester ds on ds.mahasiswa_id = m.id 
GROUP BY ds.mahasiswa_id 
;

SELECT count(*) from mahasiswa m ;

-- having
select m.nama, SUM(mk.sks) as jumlah_sks from mahasiswa m 
inner join data_semester ds ON ds.mahasiswa_id = m.id 
inner join mata_kuliah mk ON ds.matakuliah_id = mk.id 
GROUP BY m.id 
HAVING SUM(mk.sks) > 5
ORDER BY SUM(mk.sks) DESC;

-- SUBQUERY
-- tampilkan data mhs yang pernah mengambil matkul
SELECT * FROM mahasiswa m 
WHERE m.id IN (select ds.mahasiswa_id FROM data_semester ds);

select DISTINCT ds.mahasiswa_id FROM data_semester ds;

SELECT * from mahasiswa m ;

-- subquery di perintah update
UPDATE mahasiswa m set 
m.status = "aktif"
WHERE m.id IN (select ds.mahasiswa_id FROM data_semester ds);

UPDATE mahasiswa m set 
m.status = "tidak aktif"
WHERE m.id NOT IN (select ds.mahasiswa_id FROM data_semester ds);


UPDATE mahasiswa m set 
m.status = "aktif"
WHERE m.id = 1010;

SELECT id, nama, jenis_kelamin, alamat, jurusan_id, telp, status FROM mahasiswa;



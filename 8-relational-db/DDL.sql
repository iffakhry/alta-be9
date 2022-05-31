CREATE DATABASE db_be9;

USE db_be9;

-- DDL
-- no 1
CREATE TABLE jurusan(
id INT primary key not null auto_increment,
nama varchar(200)
);

-- no 2
CREATE TABlE mahasiswa (
id INT PRIMARY KEY NOT NULL,
nama varchar(100),
jenis_kelamin varchar(20),
alamat longtext,
jurusan_id int,
telp varchar(15),
CONSTRAINT FK_MahasiswaJurusan FOREIGN KEY (jurusan_id) REFERENCES jurusan(id)
);

-- no 3
CREATE TABLE mata_kuliah(
id INT NOT NULL,
nama VARCHAR(200) NOT NULL,
sks SMALLINT NOT NULL,
PRIMARY KEY (id)
);

CREATE TABLE data_semester(
id INT PRIMARY KEY NOT NULL,
mahasiswa_id INT NOT NULL,
matakuliah_id INT NOT NULL,
semester INT NOT NULL,
tahun_ajaran varchar(50) NOT NULL,
CONSTRAINT FK_DataSemesterMahasiswa FOREIGN KEY (mahasiswa_id) REFERENCES mahasiswa(id),
CONSTRAINT FK_DataSemesterMatkul FOREIGN KEY (matakuliah_id) REFERENCES mata_kuliah(id)
);


ALTER TABLE ruangan 
ADD COLUMN created_at DATETIME default CURRENT_TIMESTAMP;

show columns from ruangan;

create table dummy(
id int
);

DROP table dummy;
show tables;






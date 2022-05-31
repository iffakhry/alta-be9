SELECT id, nama, alamat, jurusan_id FROM mahasiswa m ;

SELECT id, nama FROM jurusan j ;

SELECT m.id, m.nama as nama_mahasiswa, m.alamat, m.jurusan_id, j.nama as nama_jurusan FROM mahasiswa m
INNER JOIN jurusan j ON m.jurusan_id = j.id;


SELECT m.id, m.nama, m.jurusan_id, j.nama FROM mahasiswa m
LEFT JOIN jurusan j ON m.jurusan_id = j.id;

SELECT m.id, m.nama, m.jurusan_id, j.nama FROM jurusan j 
LEFT JOIN mahasiswa m ON m.jurusan_id = j.id;

SELECT m.id, m.nama, m.jurusan_id, j.nama FROM jurusan j 
RIGHT JOIN mahasiswa m ON m.jurusan_id = j.id;
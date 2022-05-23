package bilangan

/*
1. habis dibagi 1
2. habis dibagi dirinya sendiri
3. maksimal habis dibagi 2 bilangan bulat positif selain 1 dan bilangan itu sendiri.

contoh:
10 --> YA --> 1,2,5,10
9 --> YA --> 1,3,9
5 --> YA --> 1,5
20 --> TIDAK  --> 1,2,4,5,10,20
15 --> 	YA	--> 1,3,5,15
*/

func BilanganKeren(angka uint64) string {
	temp := []int{}
	for i := 1; i <= int(angka); i++ {
		if int(angka)%i == 0 {
			temp = append(temp, i)
		}

	}
	if len(temp) <= 4 {
		return "YA"
	}
	return "TIDAK"
}

//tanpa slice
func BilanganKeren2(angka uint64) string {
	var temp int
	for i := 1; i <= int(angka); i++ {
		if int(angka)%i == 0 {
			temp++
		}

	}
	if temp <= 4 {
		return "YA"
	}
	return "TIDAK"
}

//lebih cepat
func BilanganKeren3(angka uint64) string {
	var temp int
	for i := 2; i < int(angka); i++ {
		if int(angka)%i == 0 {
			temp++
		}
		if temp > 2 {
			break
		}

	}
	if temp <= 2 {
		return "YA"
	}
	return "TIDAK"
}

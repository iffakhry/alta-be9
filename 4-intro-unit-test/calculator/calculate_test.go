package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
go test ./calculator

go test ./calculator -cover

===============
download testify : go get -u github.com/stretchr/testify
*/
func TestTambah(t *testing.T) {
	t.Run("Cek Tambah bilangan positif", func(t *testing.T) {
		bil1 := 20
		bil2 := 30
		expected := 50
		actual := Tambah(bil1, bil2)
		if actual != expected {
			t.Error("result error func tambah")
		}
	})

	t.Run("Cek tambah bilangan negatif", func(t *testing.T) {
		bil1 := -10
		bil2 := -20
		expected := -30
		actual := Tambah(bil1, bil2)
		if actual != expected {
			t.Error("result error func tambah bilangan negatif")
		}
	})
}

func TestKurang(t *testing.T) {
	t.Run("Case kurang bilangan positif", func(t *testing.T) {
		expected := 5
		actual := Kurang(10, 5)
		assert.Equal(t, expected, actual, "Hasil pengurangan tidak sesuai")
		// assert.Equal(t, 5, Kurang(10, 5), "Hasil pengurangan tidak sesuai")
	})
	t.Run("Case kurang bilangan negatif", func(t *testing.T) {
		expected := -10
		actual := Kurang(-15, -5)
		assert.Equal(t, expected, actual, "Hasil pengurangan tidak sesuai")
		// assert.Equal(t, 5, Kurang(10, 5), "Hasil pengurangan tidak sesuai")
	})
}

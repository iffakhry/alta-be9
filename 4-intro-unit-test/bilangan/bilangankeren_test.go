package bilangan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
go test ./bilangan -cover

go test ./bilangan -coverprofile=cover.out && go tool cover -html=cover.out

*/

func TestBilanganKeren(t *testing.T) {
	t.Run("Case check angka 10", func(t *testing.T) {
		actual := BilanganKeren(10)
		expected := "YA"
		assert.Equal(t, expected, actual, "Hasil tidak sesuai")
	})

	t.Run("Case check angka 100", func(t *testing.T) {
		actual := BilanganKeren(100)
		expected := "TIDAK"
		assert.Equal(t, expected, actual, "Hasil tidak sesuai")
	})
}

func TestBilanganKeren2(t *testing.T) {
	t.Run("Case check angka 10", func(t *testing.T) {
		actual := BilanganKeren2(10)
		expected := "YA"
		assert.Equal(t, expected, actual, "Hasil tidak sesuai")
	})

	t.Run("Case check angka 100", func(t *testing.T) {
		actual := BilanganKeren2(100)
		expected := "TIDAK"
		assert.Equal(t, expected, actual, "Hasil tidak sesuai")
	})
}

func TestBilanganKeren3(t *testing.T) {
	t.Run("Case check angka 10 BilanganKeren3", func(t *testing.T) {
		actual := BilanganKeren3(10) //2,5 --> temp=2
		expected := "YA"
		assert.Equal(t, expected, actual, "Hasil tidak sesuai")
	})

	t.Run("Case check angka 100", func(t *testing.T) {
		actual := BilanganKeren3(100) //2,5,10,20....
		expected := "TIDAK"
		assert.Equal(t, expected, actual, "Hasil tidak sesuai")
	})
}

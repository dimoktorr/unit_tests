package task1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProducts_TotalPrice_Ok(t *testing.T) {
	//Arrange
	products := Products{
		{
			ID:    1,
			Price: 4545.34,
		},
		{
			ID:    2,
			Price: 343.32,
		},
		{
			ID:    3,
			Price: 9948,
		},
		{
			ID:    4,
			Price: 134.43,
		},
	}

	want := 14971.09

	//Act
	got := products.TotalPrice()

	//Assert
	assert.Equal(t, want, got)
}
func TestProducts_TotalPrice_IsEmpty(t *testing.T) {
	//Arrange
	products := Products{}

	var want float64 = 0

	//Act
	got := products.TotalPrice()

	//Assert
	assert.Equal(t, want, got)
}

func TestProducts_TotalPrice(t *testing.T) {
	//Arrange
	testData := []struct {
		description string
		products    Products
		want        float64
	}{
		{
			description: "ok",
			products: Products{
				{
					ID:    1,
					Price: 4545.34,
				},
				{
					ID:    2,
					Price: 343.32,
				},
				{
					ID:    3,
					Price: 9948,
				},
				{
					ID:    4,
					Price: 134.43,
				},
			},
			want: 14971.09,
		},
		{
			description: "one elem",
			products: Products{
				{
					ID:    1,
					Price: 4545.34,
				},
			},
			want: 4545.34,
		},
		{
			description: "empty",
			products:    Products{},
			want:        0,
		},
	}

	for _, tt := range testData {
		t.Run(tt.description, func(t *testing.T) {
			//Act
			got := tt.products.TotalPrice()

			//Assert
			if got != tt.want {
				t.Errorf("TotalPrice() = %f; want %f", got, tt.want)
			}
		})
	}
}

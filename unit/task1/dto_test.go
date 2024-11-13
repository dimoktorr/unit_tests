package task1

import (
	"testing"
)

func TestProducts_TotalPrice(t *testing.T) {
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
	if got != want {
		t.Errorf("TotalPrice() = %f; want %f", got, want)
	}
}

func TestProducts_TotalPrice_tableTest(t *testing.T) {
	//Arrange
	testData := []struct {
		description string
		products    Products
		want        float64
	}{
		{
			description: "test_1",
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
			description: "test_2",
			products: Products{
				{
					ID:    1,
					Price: 4545.34,
				},
			},
			want: 4545.34,
		},
		{
			description: "test_3",
			products:    Products{},
			want:        0,
		},
	}

	for _, tt := range testData {
		//Act
		got := tt.products.TotalPrice()

		//Assert
		if got != tt.want {
			t.Errorf("TotalPrice() = %f; want %f", got, tt.want)
		}
	}
}

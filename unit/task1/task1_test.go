package task1

import (
	"github.com/dimoktorr/unit_tests/unit/task1/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_sum(t *testing.T) {
	//Arrange
	a := 5
	b := 5
	want := 10

	//Act
	got := sum(a, b)

	//Assert
	if got != want {
		t.Errorf("sum(%d, %d) = %d; want %d", a, b, got, want)
	}
}

func Test_toModelsProduct(t *testing.T) {
	//Arrange

	createdAt := time.Date(2021, 11, 17, 16, 31, 12, 0, time.UTC)
	updatedAt := time.Date(2021, 11, 17, 16, 31, 12, 0, time.UTC)
	testData := []struct {
		description string
		input       *models.Product
		expected    *Product
	}{
		{
			description: "product",
			input: &models.Product{
				ID:          100,
				Name:        "product",
				Description: "description",
				Price:       55554.0,
				InStock:     true,
				Category: &models.Category{
					ID:          5,
					Name:        "category_name",
					Description: "description",
				},
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				Rating: []*models.Rating{
					{
						ID:      2,
						Score:   5,
						Comment: "comment",
						User: &models.User{
							ID:        1,
							Username:  "user",
							Email:     "test@mail.ru",
							Password:  "45356hbjhvjwer23424",
							CreatedAt: createdAt,
							UpdatedAt: updatedAt,
						},
						ProductID: 100,
						CreatedAt: createdAt,
					},
				},
				Reviews: []*models.Review{
					{
						ID:      5,
						Content: "content",
						Rating:  7,
						User: &models.User{
							ID:        1,
							Username:  "user",
							Email:     "test@mail.ru",
							Password:  "45356hbjhvjwer23424",
							CreatedAt: createdAt,
							UpdatedAt: updatedAt,
						},
						ProductID: 0,
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
				},
				SKU:          "sku",
				Weight:       1.0,
				Dimensions:   "dimensions",
				Manufacturer: "manufacturer",
				Barcode:      "barcode",
			},
			expected: &Product{
				ID:          100,
				Name:        "product",
				Description: "description",
				Price:       55555.0,
				InStock:     true,
				Category: &Category{
					ID:          5,
					Name:        "category_name",
					Description: "description",
				},
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				Rating: []*Rating{
					{
						ID:      2,
						Score:   5,
						Comment: "comment",
						User: &User{
							ID:        1,
							Username:  "user",
							Email:     "test@mail.ru",
							Password:  "45356hbjhvjwer23424",
							CreatedAt: createdAt,
							UpdatedAt: updatedAt,
						},
						ProductID: 100,
						CreatedAt: createdAt,
					},
				},
				Reviews: []*Review{
					{
						ID:      5,
						Content: "content",
						Rating:  7,
						User: &User{
							ID:        1,
							Username:  "user",
							Email:     "test@mail.ru",
							Password:  "45356hbjhvjwer23424",
							CreatedAt: createdAt,
							UpdatedAt: updatedAt,
						},
						ProductID: 0,
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
				},
				SKU:          "sku",
				Weight:       1.0,
				Dimensions:   "dimensions",
				Manufacturer: "manufacturer",
				Barcode:      "barcode",
			},
		},
	}

	for tt := range testData {
		t.Run(testData[tt].description, func(t *testing.T) {
			//Act
			got := toModelsProduct(testData[tt].input)

			//Assert
			assert.Equal(t, testData[tt].expected, got)
		})
	}
}

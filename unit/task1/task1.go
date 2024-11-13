package task1

import "github.com/dimoktorr/unit_tests/unit/task1/models"

// тестирование функции с помощью пакета Т
// табличные юнит тесты
// рекомендую к прочтению https://runebook.dev/ru/docs/go/testing/index#B.ResetTimer

// go test ./... -cover Тулчеин Go имеет встроенные возможности для создания показателей покрытия кода.

func sum(a, b int) int {
	return a + b
}

func toModelsProduct(product *models.Product) *Product {
	return &Product{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		InStock:      product.InStock,
		Category:     toModelsCategory(product.Category),
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
		Rating:       toModelsRating(product.Rating),
		Reviews:      toModelsReview(product.Reviews),
		SKU:          product.SKU,
		Weight:       product.Weight,
		Dimensions:   product.Dimensions,
		Manufacturer: product.Manufacturer,
		Barcode:      product.Barcode,
	}
}

func toModelsRating(rating []*models.Rating) []*Rating {
	var ratings []*Rating
	for _, r := range rating {
		ratings = append(ratings, &Rating{
			ID:        r.ID,
			Score:     r.Score,
			Comment:   r.Comment,
			User:      toModelsUser(r.User),
			ProductID: r.ProductID,
			CreatedAt: r.CreatedAt,
		})
	}
	return ratings
}

func toModelsReview(review []*models.Review) []*Review {
	var reviews []*Review
	for _, r := range review {
		reviews = append(reviews, &Review{
			ID:        r.ID,
			Content:   r.Content,
			Rating:    r.Rating,
			User:      toModelsUser(r.User),
			ProductID: r.ProductID,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}
	return reviews
}

func toModelsUser(user *models.User) *User {
	return &User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toModelsCategory(category *models.Category) *Category {
	return &Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

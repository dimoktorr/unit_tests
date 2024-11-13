package task6

import "go.mongodb.org/mongo-driver/bson/primitive"

//bench тестирование
// По умолчанию go test -bench=. тестирует только скорость вашего кода, однако вы можете добавить флаг -benchmem,
//который позволит тестировать потребление памяти и количество аллокаций памяти.

type Chat struct {
	ID          string    `bson:"id"`
	OrderID     string    `bson:"order_id"`
	OrderNumber string    `bson:"order_number"`
	Source      string    `bson:"source"`
	Topic       string    `bson:"topic"`
	Messages    []Message `bson:"messages"`
}

type ChatMessagesPointer struct {
	ID          string     `bson:"id"`
	OrderID     string     `bson:"order_id"`
	OrderNumber string     `bson:"order_number"`
	MtsID       string     `bson:"mts_id"`
	UserEmail   string     `bson:"user_email"`
	Source      string     `bson:"source"`
	Topic       string     `bson:"topic"`
	Messages    []*Message `bson:"messages"`
}

type Message struct {
	Sender         string             `bson:"sender"`
	SystemName     string             `bson:"system_name"`
	SystemVersion  string             `bson:"system_version"`
	AppVersion     string             `bson:"app_version"`
	Device         string             `bson:"device"`
	MtsID          string             `bson:"mts_id"`
	UserEmail      string             `bson:"user_email"`
	Text           string             `bson:"text"`
	ClientDatetime primitive.DateTime `bson:"client_datetime"`
	CreatedAt      primitive.DateTime `bson:"created_at"`
	DeliveredAt    primitive.DateTime `bson:"delivered_at"`
}

type ChatInfo struct {
	ID      string `bson:"id"`
	OrderID string `bson:"order_id"`
	MtsID   string `bson:"mts_id"`
}

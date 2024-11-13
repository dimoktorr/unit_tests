package task4

//тестирование с помощью генератора мока
//suite test
// //go:generate mockgen -destination=./mock/pricing_mock_generated.go -package=mock  -source pricing.go
// go get go.uber.org/mock/gomock

//Mock - это более сложный объект, который может проверять взаимодействие с ним.
//Используется для проверки того, как тестируемый код взаимодействует с внешними зависимостями.
//Mock может проверять количество вызовов методов, порядок вызовов и переданные параметры.

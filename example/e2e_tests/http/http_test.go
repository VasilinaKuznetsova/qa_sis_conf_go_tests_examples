package http

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetExchangeRates(t *testing.T) {
	//Arrange
	c := NewHttpClient("", 10*time.Second)
	ctx := context.Background()
	url := "http://www.cbr-xml-daily.ru/daily_json.js"

	//Act
	resp, err := c.GetExchangeRatesToday(ctx, url)

	//Assert
	require.NoError(t, err, "Ошибка при получении курса")
	require.NotNil(t, resp, "Ответ не должен быть nil")
	require.NotEmpty(t, resp.Valute, "Список валют должен быть не пустым")
}

package fortune

import (
	"testing"
	"time"
)

func TestForecast(t *testing.T) {
	ti := time.Date(2021, 1, 1, 8, 4, 18, 0, time.UTC)
	res := Draw(ti)
	if res.String() != "大吉" {
		t.Error("Forecast is illigal. Expected 大吉, Result ", res.String())
	}
}

func TestForecastNormal(t *testing.T) {
	ti := time.Date(2021, 1, 1, 8, 4, 18, 0, time.UTC)
	res := Draw(ti)
	if res >= 7 {
		t.Error("Forecast is illigal. Expected <= 7, Result ", res)
	}
}
func TestString(t *testing.T) {
	c := Luck(2)
	res := c.String()
	if res != "小吉" {
		t.Error("Forecast is illigal. Expected 中吉, Result ", res)
	}
}

package eventSource

import (
	"log/slog"
	"math/rand"
	"reflect"
	"source/models"
	"strconv"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
)

func NewUserInteractionData() (*models.UserInteractionData, error) {
	var u models.UserInteractionData
	err := generateData(&u)
	if err != nil {
		return nil, err
	}
	u.TimeStamp = time.Now().Format(time.RFC3339)
	return &u, nil
}

// create a custom faker tag name events
func events(v reflect.Value) (interface{}, error) {
	userInteractionEvent := []string{
		"paigeView",
		"buttonClick",
		"adClick",
		"videoView",
		"addedToCart",
	}

	idx := rand.Intn(len(userInteractionEvent) - 1)
	return userInteractionEvent[idx], nil
}

func userId(v reflect.Value) (interface{}, error) {
	idNum := rand.Intn(10000)
	usrId := strings.Builder{}
	usrId.WriteString("usr")
	usrId.WriteString(strconv.Itoa(idNum))
	return usrId.String(), nil
}

func productId(v reflect.Value) (interface{}, error) {
	idNum := rand.Intn(100000)
	usrId := strings.Builder{}
	usrId.WriteString("prod")
	usrId.WriteString(strconv.Itoa(idNum))
	return usrId.String(), nil
}

func sessionDuration(v reflect.Value) (interface{}, error) {
	return strconv.Itoa(rand.Intn(700)), nil
}

func generateData(u *models.UserInteractionData) error {
	faker.AddProvider("events", events)
	faker.AddProvider("user_id", userId)
	faker.AddProvider("prod_id", productId)
	faker.AddProvider("sd", sessionDuration)

	if err := faker.FakeData(u); err != nil {
		slog.Error("Unable to generate the data from faker", "Details", err.Error())
		return err
	}
	return nil
}

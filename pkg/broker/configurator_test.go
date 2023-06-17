package broker_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Goboolean/shared-packages/pkg/broker"
	"github.com/Goboolean/shared-packages/pkg/resolver"
)

var (
	topic = "test"
	conf *broker.Configurator
)







func SetupConfigurator() {
	conf = broker.NewConfigurator(&resolver.Config{
		Host: os.Getenv("KAFKA_HOST"),
		Port: os.Getenv("KAFKA_PORT"),
	})
}

func TeardownConfigurator() {
	conf.Close()
}



func TestConfigurator(t *testing.T) {

	//defer func() {
	//	if r := recover(); r != nil {
	//		t.Errorf("Configurator Failed: %v", r)
	//	}
	//}()

	SetupConfigurator()
	TeardownConfigurator()
}



func TestCreateDeleteTopic(t *testing.T) {

	SetupConfigurator()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	if err := conf.CreateTopic(ctx, topic); err != nil {
		t.Errorf("CreateTopic() = %v", err)
	}
	
	if err := conf.DeleteTopic(ctx, topic); err != nil {
		t.Errorf("DeleteTopic() = %v", err)
	}

	TeardownConfigurator()
}
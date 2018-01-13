/*
Package mqttsimple is a straightforward wrapper for the "paho.mqtt.golang" library,
providing a bare-bones API to publish and subscribe to a mosquitto broker.
*/
package mqttsimple

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Mosquitto is a client object containing the broker URI as a string
//	and the wrapped-around paho.mqtt.golang Client
type Mosquitto struct {
	mqttClient mqtt.Client
	Broker     string
}

// Pub takes a topic and payload (both strings) and publishes the payload
//	to the broker on the given topic
func (mosq Mosquitto) Pub(topic string, payload string) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mosq.Broker)

	mosq.mqttClient = mqtt.NewClient(opts)

	if token := mosq.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := mosq.mqttClient.Publish(topic, 0, false, payload)
	token.Wait()

	mosq.mqttClient.Disconnect(250)
}

type callback func(string)

// Sub takes a topic and a callback function, subscribes to the broker on the
//	given topic, and calls the callback (passing the message as a string param)
//	for every published message.
func (mosq *Mosquitto) Sub(topic string, callbackParam callback) {
	receiveCount := 0
	choke := make(chan [2]string)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mosq.Broker)

	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	mosq.mqttClient = mqtt.NewClient(opts)

	if token := mosq.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := mosq.mqttClient.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		incoming := <-choke
		callbackParam(incoming[1])
		receiveCount++
	}
}

// UnSub disconnects the client from subscribing, after waiting the specified number
//	of milliseconds (for existing work to complete).
func (mosq *Mosquitto) UnSub(msToWait uint) {
	mosq.mqttClient.Disconnect(msToWait)
}

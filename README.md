# go-mqttsimple

PACKAGE DOCUMENTATION

package mqttsimple
```
    import "github.com/jasongwartz/go-mqttsimple"
```
    
Package mqttsimple is a straightforward wrapper for the "paho.mqtt.golang" library, providing a bare-bones API to publish and subscribe to a mosquitto broker.


## TYPES
```
type Mosquitto struct {
    Broker string
    // contains filtered or unexported fields
}
```
Mosquitto is a client object containing the broker URI as a string and the wrapped-around paho.mqtt.golang Client.

```
func (mosq Mosquitto) Pub(topic string, payload string)
```
Pub takes a topic and payload (both strings) and publishes the payload to the broker on the given topic.

```
func (mosq *Mosquitto) Sub(topic string, callbackParam callback)
```
Sub takes a topic and a callback function, subscribes to the broker on the given topic, and calls the callback (passing the message as a string param) for every published message.

```
func (mosq *Mosquitto) UnSub(msToWait uint)
```
UnSub disconnects the client from subscribing, after waiting the specified number of milliseconds (for existing work to complete).



# Simple Lorawan example

This demo code will connect Lorawan network and send sample uplink message

You may change your Lorawan keys (AppEUI, DevEUI, AppKEY) in key-default.go


```
$ tinygo monitor
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
Lorawan Simple Demo
Start Lorawan Join sequence
loraConnect: Connected !
```

# Building

## Simulator

```
tinygo flash -target pico ./examples/lora/lorawan/basic-demo
```

## PyBadge with LoRa Featherwing 

```
tinygo flash -target pybadge -tags featherwing ./examples/lora/lorawan/basic-demo
```

## LoRa-E5 

```
tinygo flash -target lorae5 ./examples/lora/lorawan/basic-demo
```


## Enable debugging

You can also enable some debug logs with ldflags :

```
$ tinygo build -ldflags="-X 'main.debug=true'" -target=lorae5
```


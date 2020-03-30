# Coral
A websocket based server for building real time apps.


This is current in Development



## Running the server

### Using command line

1. clone the repository.
2. From the Coral directory run the command
    go build -o server pkg/coral.go
    
3.     

### Using docker

### Demo
In this demo we are showing on one side if we user is placing an order and there are two devices installed in restaurant. Both of those devices are getting updates.

![coral_demo](https://user-images.githubusercontent.com/60743403/77870912-add9d980-7210-11ea-8694-63e5155d9f6b.gif)

For full demo you can see it on youtube: https://www.youtube.com/watch?v=wfMiPu1FEJk

## Use-case
If you have use case for a server which is maintaining lot of persitent connection to lot of devices. 
For example Tablets and Mobile phone app which need real time updates over websocket. 

Currently the coral server is used by OrderMaster.ca which order notifiations to retaurants.

![CoralServer](https://user-images.githubusercontent.com/60743403/77486107-36b5d700-6e05-11ea-80eb-cc20502824d8.png)


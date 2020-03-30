# Coral
A websocket based server for building real time apps.


This is current in Development


## Requirements
1. golang > 1.12
2. Google cloud account


## Running the server

### Using command line

1. clone the repository.
2. From the Coral directory run the command
     
     
       go build -o coral pkg/coral.go
    
3. Set the Google cloud credential file with command

         export GOOGLE_APPLICATION_CREDENTIALS=<PATH_TO_CREDENTIAL_FILE>
         
4. Run the build binary file. By default this would pick dev properties file. which you have to change.
       
         ./coral

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


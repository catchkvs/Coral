# Coral
A websocket based server for building real time apps. This is specifically focusing on Apps which require the Fact Table notifications to be sent to Dimension side consumer.
One example is:
  Order update notifications to restaurants.


## Features

1. Simple to operate. Built in docker image.
2. Completely open source so easy to customize for your own Fact and Dimensions Table.
3. Complete support for cloud environment.
4. Session Management (expiry and multiple concurrent session handling)
5. Achieves low latency through go channels and routine support.


This is current in Development


## Requirements
1. golang > 1.12
2. Google cloud account


## Running the server
Below instructions have been tested on Ubuntu.
MAC and windows operating system users might have to changes things according to their Operating System guidelines.

### Using command line

1. clone the repository.
2. From the Coral directory run the command
     
       go build -o coral pkg/coral.go
    
3. Set the Google cloud credential file with command

       export GOOGLE_APPLICATION_CREDENTIALS=<PATH_TO_CREDENTIAL_FILE>
         
4. Run the build binary file. By default this would pick dev properties file. which you have to change.
       
       ./coral
         
5. Once the server is running you can open the client.html to test the server as showing in demo
### Using docker
1. Clone the repository.

2. Docker file is included just run the command
    
           docker build -t coral:1.0 .

3. Once the docker container is build you can run through docker command. Ensure that Google Application credential
   environment variable is set appropriately.
   
4. Now again use the client.html to connect and test it yourself.           
           

### Demo
In this demo we are showing on one side if we user is placing an order and there are two devices installed in restaurant. Both of those devices are getting updates.

![coral_demo](https://user-images.githubusercontent.com/60743403/77870912-add9d980-7210-11ea-8694-63e5155d9f6b.gif)

For full demo you can see it on youtube: https://www.youtube.com/watch?v=wfMiPu1FEJk

## Use-case
If you have use case for a server which is maintaining lot of persitent connection to lot of devices. 
For example Tablets and Mobile phone app which need real time updates over websocket. 

Currently the coral server is used by OrderMaster.ca which order notifiations to retaurants.

![CoralServer](https://user-images.githubusercontent.com/60743403/77486107-36b5d700-6e05-11ea-80eb-cc20502824d8.png)


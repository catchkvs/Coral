$( document ).ready(function() {
    console.log("Document is ready...Starting connections.");
    var factWS = new WebSocket("ws://localhost:4040/session");
    var dimensionWS1 = new WebSocket('ws://localhost:4040/session');
    var dimensionWS2 = new WebSocket('ws://localhost:4040/session');
    var authToken = "6bae761605833e9f9c5494fab4ff2975";
    var factSessionId;
    var dimensionSessionId1;
    var dimensionSessionId2;

    function placeOrder() {
        console.log("Placing order:")
    }

    factWS.onmessage = function(event) {
        console.log(event.data);
        var serverMsg = JSON.parse(event.data);
        console.log(serverMsg.Command);
        console.log(serverMsg.Data);
        switch(serverMsg.Command) {
            case "ReceiveSessionId":
                factSessionId = serverMsg.Data;
                break;
        }
    }

    factWS.onopen = function() {
        console.log("clientWebSocket.onopen", factWS);
        console.log("clientWebSocket.readyState", "websocketstatus");
        factWS.send("ESTABLISHED");
    }

    factWS.onclose = function(error) {
        console.log("clientWebSocket.onclose", factWS, error);
        //events("Closing connection");
    }

    factWS.onerror = function(error) {
        console.log("clientWebSocket.onerror", factWS, error);
        //events("An error occured");
    }


    function placeNewOrder() {
        var fact = {
            Name : "Order",
            DimensionId : "123",
            Attributes: {
                "CustomerName": "Dummy",
                "PhoneNumber": "Dummy",
            }
        }
        console.log(JSON.stringify(fact));
        var encodedOrder = btoa(JSON.stringify(fact));
        sendMessageWithCommand(encodedOrder, "CreateFactEntity")
    }



    function sendMessageWithCommand(message, command) {
        var msg = {
            data: message,
            SessionId : factSessionId,
            AuthToken : authToken,
            command: command,
        };
        factWS.send(JSON.stringify(msg));
    }

    dimensionWS1.onmessage = function(event) {
        console.log(event.data);
        console.log(event.data);
        var serverMsg = JSON.parse(event.data);
        console.log(serverMsg.Command);
        console.log(serverMsg.Data);
        switch(serverMsg.Command) {
            case "ReceiveSessionId":
                dimensionSessionId1 = serverMsg.Data;
                var dimensionConnInput = {
                    Id: "123",
                    Name: "restaurant"
                }
                var encodedData = btoa(JSON.stringify(dimensionConnInput));
                var msg = {
                    data: encodedData,
                    SessionId : dimensionSessionId1,
                    AuthToken : authToken,
                    command: "GetLiveUpdates",
                };
                dimensionWS1.send(JSON.stringify(msg));
                break;
        }
    }

    dimensionWS1.onopen = function() {
        console.log("clientWebSocket.onopen", dimensionWS1);
        console.log("clientWebSocket.readyState", "websocketstatus");

    }

    dimensionWS1.onclose = function(error) {
        console.log("clientWebSocket.onclose", dimensionWS1, error);
        //events("Closing connection");
    }

    dimensionWS1.onerror = function(error) {
        console.log("clientWebSocket.onerror", dimensionWS1, error);
        //events("An error occured");
    }

    dimensionWS2.onmessage = function(event) {
        console.log(event.data);
        console.log(event.data);
        var serverMsg = JSON.parse(event.data);
        console.log(serverMsg.Command);
        console.log(serverMsg.Data);
        switch(serverMsg.Command) {
            case "ReceiveSessionId":
                dimensionSessionId2 = serverMsg.Data;
                var dimensionConnInput = {
                    Id: "123",
                    Name: "restaurant"
                }
                var encodedData = btoa(JSON.stringify(dimensionConnInput));
                var msg = {
                    data: encodedData,
                    SessionId : dimensionSessionId2,
                    AuthToken : authToken,
                    command: "GetLiveUpdates",
                };
                dimensionSessionId2.send(JSON.stringify(msg));
                break;
        }
    }

    dimensionWS2.onopen = function() {
        console.log("clientWebSocket.onopen", dimensionWS2);
        console.log("clientWebSocket.readyState", "websocketstatus");
        sendMessageWithCommand("123", "GetLiveUpdates")
    }

    dimensionWS2.onclose = function(error) {
        console.log("clientWebSocket.onclose", dimensionWS2, error);
        //events("Closing connection");
    }

    dimensionWS2.onerror = function(error) {
        console.log("clientWebSocket.onerror", dimensionWS2, error);
        //events("An error occured");
    }

});


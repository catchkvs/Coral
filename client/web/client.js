factWS = new WebSocket('ws://localhost:4040/session');
dimensionWS1 = new WebSocket('ws://localhost:4040/session');
dimensionWS2 = new WebSocket('ws://localhost:4040/session');

factWS.onmessage = function(event) {
    console.log(event.data);
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

dimensionWS1.onmessage = function(event) {
    console.log(event.data);
}

dimensionWS1.onopen = function() {
    console.log("clientWebSocket.onopen", dimensionWS1);
    console.log("clientWebSocket.readyState", "websocketstatus");
    dimensionWS1.send("ESTABLISHED");
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
    if (!sessionId) {
        console.log("Setting sessionId")
    }
}

dimensionWS2.onopen = function() {
    console.log("clientWebSocket.onopen", dimensionWS2);
    console.log("clientWebSocket.readyState", "websocketstatus");
    dimensionWS2.send("ESTABLISHED");
}

dimensionWS2.onclose = function(error) {
    console.log("clientWebSocket.onclose", dimensionWS2, error);
    //events("Closing connection");
}

dimensionWS2.onerror = function(error) {
    console.log("clientWebSocket.onerror", dimensionWS2, error);
    //events("An error occured");
}
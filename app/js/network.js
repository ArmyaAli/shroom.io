const URL = "ws://localhost:8090/websocket";

// Create a web socket - Async btw
socket = new WebSocket(URL);

// When the socket is open, we then send the message
socket.onopen = ($event) => { socket.send("Test message to for server"); }

// Recieve messages
socket.onmessage = ($event) => { console.log($event.data); };




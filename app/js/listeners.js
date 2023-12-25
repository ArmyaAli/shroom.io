import { Websocket, Player, MOVE_SPEED } from "./constants.js"

window.addEventListener("keydown", (event) => {
  if (event.code === "ArrowDown"){
    Player.y += MOVE_SPEED;
  } else if (event.code === "ArrowUp"){
    Player.y -= MOVE_SPEED;
  } else if (event.code === "ArrowLeft"){
    Player.x -= MOVE_SPEED;
  } else if (event.code === "ArrowRight"){
    Player.x += MOVE_SPEED;
  }

  if(Websocket.websocket.readyState === 1) { 
  console.log("Pressed");
    Websocket.websocket.send(JSON.stringify(Player)); 
  }

});

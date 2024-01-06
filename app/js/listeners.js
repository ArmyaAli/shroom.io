import { Websocket, Player, MOVE_SPEED, FPS } from "./constants.js"

window.addEventListener("keydown", (event) => {
  if (event.code === "ArrowDown"){
    Player.vel.y = 0.1;
  } else if (event.code === "ArrowUp"){
    Player.vel.y = -0.1;
  } else if (event.code === "ArrowLeft"){
    Player.vel.x = -0.1;
  } else if (event.code === "ArrowRight"){
    Player.vel.x = 0.1;
  }
});

window.addEventListener("keyup", (event) => {
  if (event.code === "ArrowDown"){
    Player.vel.y = 0;
  } else if (event.code === "ArrowUp"){
    Player.vel.y = 0;
  } else if (event.code === "ArrowLeft"){
    Player.vel.x = 0;
  } else if (event.code === "ArrowRight"){
    Player.vel.x = 0;
  }
});

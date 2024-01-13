import { Websocket, KEYCODES, GAMEFLAGS, Mouse, Player, MOVE_SPEED, FPS, Canvas as canvas} from "./constants.js"

window.addEventListener("keydown", (event) => {
  if (event.code === KEYCODES.W){ GAMEFLAGS.UPKEYDOWN = true } 
  if (event.code === KEYCODES.A){ GAMEFLAGS.LEFTKEYDOWN = true } 
  if (event.code === KEYCODES.S){ GAMEFLAGS.DOWNKEYDOWN = true }
  if (event.code === KEYCODES.D){ GAMEFLAGS.RIGHTKEYDOWN = true }
});

window.addEventListener("keyup", (event) => {
  if (event.code === KEYCODES.W){ GAMEFLAGS.UPKEYDOWN = false } 
  if (event.code === KEYCODES.A){ GAMEFLAGS.LEFTKEYDOWN = false } 
  if (event.code === KEYCODES.S){ GAMEFLAGS.DOWNKEYDOWN = false }
  if (event.code === KEYCODES.D){ GAMEFLAGS.RIGHTKEYDOWN = false }
});

canvas.addEventListener('mouseenter', (event) => { GAMEFLAGS.MOUSEONSCREEN = true });
canvas.addEventListener('mouseleave', (event) => { GAMEFLAGS.MOUSEONSCREEN = false });

canvas.addEventListener('mousemove', (event) => { 
  console.log(event)
  Mouse.pos.x = event.clientX;
  Mouse.pos.y = event.clientY;
});

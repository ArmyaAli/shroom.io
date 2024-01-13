import { FPS_LIMIT, FPS_INTERVAL, Mouse, PLAYERSHIELD, GAMEFLAGS, VIEW_MAP, FPS, Player, PLAYER_SIZE, Websocket, Canvas as canvas } from "./constants.js"
import { init_websocket} from "./network.js"
import { update_player_positions } from "./bus.js"

const ctx = canvas.getContext("2d");

const draw_player = (x, y, name) => {
  // Draw name
  ctx.font = "16px serif";
  ctx.fillText(name, x - 16, y - 32);

  // Draw body
  ctx.beginPath();
  ctx.arc(x, y, Player.radius , 0, 2 * Math.PI);
  ctx.fill();
  ctx.lineWidth = 2 
  ctx.strokeStyle = "green";
  ctx.stroke();
  ctx.closePath();
  
  // DRAW PLAYER OUTLINE HERE IF HOVER FLAG IS TRUE
  //////////////////////////////////////////////////
  //
  if(GAMEFLAGS.MOUSEONSCREEN) {
    console.log("Drawing Arc")

    const mouseX = Mouse.pos.x - canvas.getBoundingClientRect().left;
    const mouseY = Mouse.pos.y - canvas.getBoundingClientRect().top;

    // Calculate the angle between the mouse and the circle center
    PLAYERSHIELD.angle = Math.atan2(mouseY - Player.pos.y, mouseX - Player.pos.x);

    // Calculate the position of the arc on the circle
    PLAYERSHIELD.pos.x = Player.pos.x + Player.radius * Math.cos(PLAYERSHIELD.angle);
    PLAYERSHIELD.pos.y = Player.pos.y + Player.radius * Math.sin(PLAYERSHIELD.angle);

    ctx.beginPath();    // Calculate the new position of the arc based on the mouse
    ctx.arc(PLAYERSHIELD.pos.x, PLAYERSHIELD.pos.y, PLAYERSHIELD.radius, PLAYERSHIELD.angle + Math.PI, PLAYERSHIELD.angle);
    
    ctx.stroke();
    ctx.closePath();
  }
}

const draw_existing_players = () => {
  let i = 0

  for(const [key, val] of VIEW_MAP.entries()) {
    if(Player.id !== key) {
      draw_player(val.pos.x, val.pos.y, `Ali-${i}`) 
      ++i
    }
  }
}

const prep_player_drop = () => {
  Player.pos.x = Math.random() * 764 
  Player.pos.y = Math.random() * 364 
  Player.active = true;
}

const update = () => {
  const friction = 0.98;

  if(GAMEFLAGS.LEFTKEYDOWN) { Player.vel.x -= Player.acc.x } 
  else { Player.vel.x *= friction }

  if(GAMEFLAGS.RIGHTKEYDOWN) { Player.vel.x += Player.acc.x } 
  else { Player.vel.x *= friction }

  if(GAMEFLAGS.UPKEYDOWN) { Player.vel.y -= Player.acc.y } 
  else { Player.vel.y *= friction }

  if(GAMEFLAGS.DOWNKEYDOWN) { Player.vel.y += Player.acc.x } 
  else { Player.vel.y *= friction }

  Player.pos.x += Player.vel.x
  Player.pos.y += Player.vel.y

  update_player_positions();
}

const draw = () => {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  draw_existing_players();
  draw_player(Player.pos.x, Player.pos.y, Player.nick);
}

const bootstrap_game = () => {
  prep_player_drop();
}


const game_loop = () => {
  FPS.now = performance.now();
  FPS.dt = FPS.now - FPS.then;

  if(FPS.dt > FPS_INTERVAL) {
    FPS.then = FPS.now - (FPS.dt % FPS_INTERVAL)
    update();
    draw();
  }

  requestAnimationFrame(game_loop);
}

const game = () => {
  bootstrap_game();
  // INF LOOP
  game_loop();
}

game()

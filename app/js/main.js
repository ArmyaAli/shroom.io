import { FPS_LIMIT, FPS_INTERVAL, VIEW_MAP, FPS, Player, PLAYER_SIZE, Websocket, Canvas as canvas } from "./constants.js"
import { init_websocket} from "./network.js"
import { update_player_positions } from "./bus.js"

const ctx = canvas.getContext("2d");

const draw_player = (x, y, name) => {
  // Draw name
  ctx.font = "16px serif";
  ctx.fillText(name, x - 16, y - 32);
  
  // Draw body
  ctx.beginPath();
  ctx.arc(x, y, PLAYER_SIZE, 0, 2 * Math.PI);
  ctx.fill();
  ctx.stroke();
  ctx.closePath();
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
  Player.id = crypto.randomUUID(); 
  Player.pos.x = Math.random() * 764 
  Player.pos.y = Math.random() * 364 
  Player.active = true;
}

const update = () => {
  Player.pos.x += Player.vel.x * FPS.dt;
  Player.pos.y += Player.vel.y * FPS.dt;
  update_player_positions();
}

const draw = () => {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  draw_existing_players();
  draw_player(Player.pos.x, Player.pos.y, Player.nick);
}

const bootstrap_game = () => {
  init_websocket();
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
  game_loop();
}

game();

import { FPS_LIMIT, FPS_INTERVAL, VIEW_BUFFER, FPS, Player, PLAYER_SIZE, Websocket, Canvas as canvas } from "./constants.js"
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
  for(let i = 0; i < VIEW_BUFFER.length; ++i) {
    draw_player(VIEW_BUFFER[i].pos.x, VIEW_BUFFER[i].pos.y, "player 2");
  }
}

const prep_player_drop = () => {
  Player.id = crypto.randomUUID(); 
  Player.name = `ali-${Math.floor(Math.random() * 10)}`;
  Player.x = Math.random() * 764 
  Player.y = Math.random() * 364 
  Player.active = true;
}

const update = () => {
  update_player_positions();
  Player.x += Player.vel.x * FPS.dt;
  Player.y += Player.vel.y * FPS.dt;
}

const draw = () => {
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  draw_existing_players();
  draw_player(Player.x, Player.y, Player.name);
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

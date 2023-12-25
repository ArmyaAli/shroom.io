import { Player, PLAYER_SIZE, Websocket, PLAYERS_IN_VIEW, Canvas as canvas } from "./constants.js"
import { init_websocket} from "./network.js"
import { update_player_positions } from "./message_bus.js"

const ctx = canvas.getContext("2d");


const draw_player = (x, y, name) => {
  // Draw name
  ctx.font = "16px serif";
  console.log("Hello", name, x, y)
  ctx.fillText(name, x - 16, y - 32);
  
  // Draw body
  ctx.beginPath();
  ctx.arc(x, y, PLAYER_SIZE, 0, 2 * Math.PI);
  ctx.fill();
  ctx.stroke();
  ctx.closePath();
}

const draw_existing_players = () => {
  for(let i = 0; i < 4; ++i) {
    // if player is active
    const player = PLAYERS_IN_VIEW[i];
    if(player != null && player.active === true) {
      draw_player(player.x, player.y, player.name);
    }
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
  // Requests
  draw();
  requestAnimationFrame(game_loop);
}

const game = () => {
  bootstrap_game();
  game_loop();
}

game();

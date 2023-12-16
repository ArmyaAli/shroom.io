import { Player, WEBSOCKET, PLAYERS_IN_VIEW, Game } from "./constants.js"

export const set_handshake_data = (data) => {
  const player_data = JSON.parse(data);
  
  for(let i = 0; i < 64; ++i) {
    PLAYERS_IN_VIEW.push({ x: player_data[i].X, y: player_data[i].Y, name: player_data[i].Name });
  }

  Game.GAME_STATE = Game.GAME_STATE << 1;
  console.log(PLAYERS_IN_VIEW, Game.GAME_STATE)
}

export const register_player = () => {
  WEBSOCKET.websocket.send(JSON.stringify(Player))
}

export const update_player_positions = () => {
  console.log("in game state 2");
}

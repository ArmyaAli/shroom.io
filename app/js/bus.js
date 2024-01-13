import { Player, Websocket, Game } from "./constants.js"

export const set_handshake_data = (data) => {
  const player_data = JSON.parse(data);
  console.log(player_data)
  
  for(let i = 0; i < 64; ++i) {
    PLAYERS_IN_VIEW.push({ x: player_data[i].X, y: player_data[i].Y, name: player_data[i].Name, active: player_data[i].Active });
  }

  Game.GAME_STATE = Game.GAME_STATE << 1;
  console.log(Game.GAME_STATE)
}

export const register_player = () => {
  const to_send = { channel: "player", sessionid: Player.id, timestamp: Date.now(),  payload: Player };
  Websocket.websocket.send(JSON.stringify(to_send))
}

export const update_player_positions = () => {
  if(Websocket?.websocket?.readyState === 1) {
    const to_send = { channel: "position", sessionid: Player.id, timestamp: Date.now(), payload: Player };
    Websocket.websocket.send(JSON.stringify(to_send))
  }
}


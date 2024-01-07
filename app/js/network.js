import { Websocket, Player, Game, GAME_STATE_DISPATCH_MAP, VIEW_MAP } from "./constants.js"
import { register_player } from './bus.js'

export const init_websocket = () => {
  Websocket.websocket = new WebSocket(Websocket.url);

  Websocket.websocket.onopen = ($event) => { 
    register_player();
  }

  Websocket.websocket.onmessage = socketRouter;

  Websocket.websocket.onclose = ($event) => { 
    Player.active = false;
  };
}


// Channel the data appropriatly
const socketRouter = ($event) => {
  // Need to know if we need to update the player information or add a new player
  const data = JSON.parse($event.data);
  for(const key of Object.keys(data)) {
  }
  //console.log(data.PLAYER_MAP)
  for(const key of Object.keys(data.PLAYER_MAP)) {
    const entry = data.PLAYER_MAP[key]
    VIEW_MAP.set(key, entry)
  }
}


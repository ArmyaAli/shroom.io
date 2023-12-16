import { WEBSOCKET, Player, Game, GAME_STATE_DISPATCH_MAP } from "./constants.js"
import { register_player } from './message_bus.js'

export const init_websocket = () => {
  WEBSOCKET.websocket = new WebSocket(WEBSOCKET.url);

  WEBSOCKET.websocket.onopen = ($event) => { 
    // Send information about us to the server
    WEBSOCKET.websocket.send(JSON.stringify(Player));
    register_player();
  }

  WEBSOCKET.websocket.onmessage = ($event) => { 
    GAME_STATE_DISPATCH_MAP[Game.GAME_STATE]($event.data);
  };

  WEBSOCKET.websocket.onclose = ($event) => { 
    Player.active = false;
  };
}





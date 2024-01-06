import { Websocket, Player, Game, GAME_STATE_DISPATCH_MAP, VIEW_BUFFER } from "./constants.js"
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
  const data = JSON.parse($event.data);
  //
  // Add our item on the View Queue
  VIEW_BUFFER.push(data);
}


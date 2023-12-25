import { Websocket, Player, Game, GAME_STATE_DISPATCH_MAP } from "./constants.js"
import { register_player } from './message_bus.js'

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


const socketRouter = ($data) => {
  console.log("Router Channeler")
  console.log($data.channel)
}


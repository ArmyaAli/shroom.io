import { Websocket, Player, Game, GAME_STATE_DISPATCH_MAP, VIEW_MAP } from "./constants.js"
import { register_player } from './bus.js'

export const init_websocket = async () => {
  // before we open a websocket let's register our client
  //
  //
  await registerClient();
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

  //console.log(data)
  for(let i = 0; i < 1024; ++i) {
    if(Player.id !== data[i].id && data[i].id !== '') {
      VIEW_MAP.set(data[i].id, data[i]) 
    }
  }

  //console.log(VIEW_MAP)
}

const registerClient = async () => {
  const response = await fetch(`http://localhost:8090/register?id=${Player.id}&email=${Player.email}`, {
    credentials: "include"
  });
}

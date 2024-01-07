import { set_handshake_data, update_player_positions } from './bus.js'

export const Websocket = {
  websocket: null,
  url: "ws://localhost:8090/websocket"
}

export const Canvas = document.querySelector("canvas");
export const Pocketbase = new PocketBase('http://localhost:8090');

export const Player = {
  id: "",
  name: "Ali Umar",
  vel: { x: 0, y: 0 },
  pos: { x: 0, y: 0 },
};

export const FPS_LIMIT = 60;
export const FPS_INTERVAL = 1000 / FPS_LIMIT;

export const FPS = {
  acc: -1,
  now: -1,
  then: performance.now(),
  dt: -1
};


// View Id
// Player Count
// Player[]
export const VIEW_MAP = new Map();

export const PLAYER_SIZE = 25;

export const SESSION_ID = -1;

export const Game = {
  GAME_MAP: new Array(800 * 5)
                 .fill(0)
                 .map(() => new Array(400 * 20).fill(0)),

  GAME_STATE: 1, 
  NET_STATE: 1,
  IS_AUTH: false
};

export const MOVE_SPEED = 10;
// Always start in BOOTSTRAP state

// Game States - 8 bit wide bitfield
export const SYS_BOOTSTRAP      = 0b00000001;
export const GAME_MENU          = 0b00000010;
export const GAME_PLAYING       = 0b00000100;
//export const GAME_PLAYING     = 0b00000100;
//export const GAME_PLAYING     = 0b00000100;
//export const GAME_PLAYING     = 0b00000100;
//export const GAME_PLAYING     = 0b00000100;

// Manage communication pipe state
export const NET_HANDSHAKE    = 0x1;
export const NET_CONNECTED    = 0x2;
export const NET_DISCONNECTED = 0x3;

export const GAME_STATE_DISPATCH_MAP = {
  1: set_handshake_data,
  2: update_player_positions
}

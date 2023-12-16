import { set_handshake_data, update_player_positions } from './message_bus.js'
export const WEBSOCKET = {
  websocket: null,
  url: "ws://localhost:8090/websocket"
}

export const Player = {
  id: 1,
  name: "Ali",
  x: 100,
  y: 100,
  active: true
}

export const PLAYERS_IN_VIEW = [];

export const PLAYER_SIZE = 25;

export const Game = {
  GAME_MAP: new Array(800 * 5)
                 .fill(0)
                 .map(() => new Array(400 * 20).fill(0)),

  GAME_STATE: 1, 
  NET_STATE: 1
};

export const MOVE_SPEED = 10;
// Always start in BOOTSTRAP state

// Game States - 8 bit wide bitfield
export const SYS_BOOTSTRAP    = 0b00000001;
export const GAME_PRE_PLAYING = 0b00000010;
export const GAME_PLAYING     = 0b00000100;
//export const GAME_PLAYING     = 0b00000100;
//export const GAME_PLAYING     = 0b00000100;
//export const GAME_PLAYING     = 0b00000100;
//export const GAME_PLAYING     = 0b00000100;

// Manage communication pipe state
export const NET_HANDSHAKE = 0b00000001;
export const NET_GAME      = 0b00000010;

export const GAME_STATE_DISPATCH_MAP = {
  1: set_handshake_data,
  2: update_player_positions
}

/* 
//TODO(Ali) 
 *
 * I want to implement a login flow using Discord. We will register an application with discord
 * Then we will register redirect URLs with discord
 * We will keep state hash and session hash
 * Call /authorize with discord
 * Redirect back to our page once authenticated
 * Read the URL for the code value that discord will send over
 * Use the code value to obtain the token and clear in the URL bar
 * Then we are now authenticated
 * We will also implement a mechanism to logout
 * The logout mechanism may be proxied from our backend due to discord /revoke token endpoint
 * requiring client_id and client_secret to be passed to make a successful call (unsecure in the browser)
 *
*/


import { Game, Player, Pocketbase as pb } from './constants.js'
import { UI_setLoggedInView } from './uicontrols.js';
import { init_websocket } from './network.js';
// This method initializes a one-off realtime subscription and will
// open a popup window with the OAuth2 vendor page to authenticate.
//
// Once the external OAuth2 sign-in/sign-up flow is completed, the popup
// window will be automatically closed and the OAuth2 data sent back
// to the user through the previously established realtime connection.
//
// If the popup is being blocked on Safari, you can try the suggestion from:
// https://github.com/pocketbase/pocketbase/discussions/2429#discussioncomment-5943061.
// 

// Check Auth
const checkRegistrationCookie = () => {
  const cookie = {} 
  console.log(document.cookie)
  return document.cookie;
}

(async () => {
  if(pb.authStore.isValid) {
    Game.IS_AUTH = true;
    Player.id = crypto.randomUUID(); 
    Player.email = pb.authStore.model.email
    UI_setLoggedInView();

    checkRegistrationCookie()
    if(Player.nick === null || Player.nick === undefined || Player.nick === "") {
      Player.nick = pb.authStore.model.username
    }

    await init_websocket();
    return
  }
})()




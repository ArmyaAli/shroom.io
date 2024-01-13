import { validateLoginInput } from './validate.js'
import { Pocketbase as pb, Canvas as canvas, Game, Player } from './constants.js'

const LoginGroup = document.querySelector("#__login")
const LoginValidationLabel = document.querySelector("#__login > #__validateNickname")
const LoginTextInput = document.querySelector("#__login > input");
const LoginButton = document.querySelector("#__login > button")
const LogoutButton = document.querySelector("#__game-view > button")
const GameViewGroup = document.querySelector("#__game-view")
const DisplayTextContainer = document.querySelector("#__game-view-text")

LoginButton.addEventListener("click", async ($event) => {
    // Initiate the discord login process
    const nickname = LoginTextInput.value;

    if (validateLoginInput(nickname)) {
        LoginValidationLabel.innerHTML = "Logging in...";
        if (!Game.IS_AUTH) await pb.collection('users').authWithOAuth2({ provider: 'discord' });

        if (pb.authStore.isValid) {
            UI_setLoggedInView();
            Player.email = pb.authStore.model.email
            Player.id = crypto.randomUUID(); 
            if(Player.nick === null || Player.nick === undefined || Player.nick === "") {
              Player.nick = pb.authStore.model.username
            }
        } else {
            console.log("unable to login")
        }
    } else {
        console.log("Unsuccessful Login")
        LoginValidationLabel.innerHTML = "Cannot Login";
    }
})

LogoutButton.addEventListener("click", async ($event) => {
    // logout of the client
    pb.authStore.clear();
    UI_setLoggedOutView();
})

export const UI_setLoggedInView = () => {
    LoginGroup.style.display = "none";
    GameViewGroup.style.display = "block";

//    for (const key of Object.keys(pb.authStore.model)) {
//      const span = document.createElement("span");
//      span.innerHTML = `${key}: ${pb.authStore.model[key]}` 
//      DisplayTextContainer.appendChild(span);
//    }
    
}

export const UI_setLoggedOutView = () => {
    for (const c of Array.from(DisplayTextContainer.children)) c.remove();
    LoginGroup.style.display = "flex";
    GameViewGroup.style.display = "none";
}

export const UI_clearAllLabels = () => {
    LoginValidationLabel.innerHTML = "";
}

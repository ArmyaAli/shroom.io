//TODO(Ali) 
/* 
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



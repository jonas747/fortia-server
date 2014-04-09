var Fortia = Fortia || {};
Fortia.Net = Fortia.Net || {};

/**
(SERVER) Sends a message to the player "player"
@function sendUsrMessage
@param {String} name - The name of the message, You can attach listeners to names
@param {Object} data - The data, will be serialized with json
@param {Player} player - The player that receives the message
@memberof Net
**/
Fortia.Net.sendUsrMessage = function(name, data, player){
	// Need native stuff
	_fortiaSendUsrMessage(name, data, player);
}

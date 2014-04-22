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
Fortia.Net.sendUsrMessage = function(player, name, data, compress){
	if(arguments.length < 3){
		console.error("Not enough arguments to call Fortia.Net.sendUsrMessage()")
		return
	}
	// Need native stuff
	Fortia._sendUsrMessage(player, name, data, compress);
}

/**
Emits a network message event
@function emit
@param {String} name - The message name we will listen for
@param {object} data - The message data.
@memberof Net
**/ 
Fortia.Net.emit = function(name, data, sender){
	if(!Fortia.Net._eventListeners){
		return
	}
	if(!Fortia.Net._eventListeners[name]){
		return
	}
	var dataDecoded = JSON.parse(data);
	for (var i = 0; i < Fortia.Net._eventListeners[name].length; i++) {
		Fortia.Net._eventListeners[name][i](dataDecoded, Fortia.getPlayer(sender))			
	};
}
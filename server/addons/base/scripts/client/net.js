var Fortia = Fortia ||  {};

if(!Fortia.Net){
	FOrtia.Net = {}
}
/**
(CLIENT) Sends a message to the server
@function sendMessage
@param {String} name - The name of the message, You can attach listeners to names
@param {Object} data - The data, will be serialized with json
@memberof Net
**/
Fortia.Net.sendMessage = function(name, data){
	_fortiaSendMessage(name, data);
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
	for (var i = 0; i < Fortia.Net._eventListeners[name].length; i++) {
		if(_fortiaclient){
			Fortia.Net._eventListeners[name][i](data)
		}else{
			Fortia.Net._eventListeners[name][i](data, sender)			
		}
	};
}
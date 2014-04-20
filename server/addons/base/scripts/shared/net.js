var Fortia = Fortia || {};

/**
Networking namespace
@namespace Net
**/
if(!Fortia.Net){
	Fortia.Net = {};
}

/**
Listens for networked messages
@function on
@param {String} name - The message name we will listen for
@param {Function} callback - The callback. Will be called with data. and sender if were on the server
@memberof Net
**/ 
Fortia.Net.on = function(name, callback){
	if(!Fortia.Net._eventListeners){
		Fortia.Net._eventListeners = {};
	}

	if(!Fortia.Net._eventListeners[name]){
		Fortia.Net._eventListeners[name] = [];
	}

	Fortia.Net._eventListeners[name].push(callback)
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
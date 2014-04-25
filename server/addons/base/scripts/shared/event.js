var Fortia = Fortia || {};
Fortia._eventListeners = {};
Fortia.on = function(evt, callback){
	Fortia._eventListeners = Fortia._eventListeners || {};
	if(!Fortia._eventListeners[evt]){
		Fortia._eventListeners[evt] = new Array();;
	}
	Fortia._eventListeners[evt].push(callback)
}

Fortia.emit = function(evt){
	if(!Fortia._eventListeners[evt]){
		return;
	}

	// Build an array with the rest of the arguments
	var rest = [];
	if(arguments.length > 1){
		for (var i = 1; i < arguments.length; i++) {
			rest.push(arguments[i]);
		};
	}

	for (var i = 0; i < Fortia._eventListeners[evt].length; i++) {
		Fortia._eventListeners[evt][i].apply(Fortia._eventListeners[evt][i], rest);
	};
}
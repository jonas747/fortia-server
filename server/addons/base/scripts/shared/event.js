var Fortia = Fortia || {};

Fortia.on = function(evt, callback){
	Fortia._eventListenvers = Fortia._eventListenvers || {};
	if(!Fortia._eventListenvers[evt]){
		Fortia._eventListenvers[evt] = new Array();;
	}
	Fortia._eventListenvers[evt].push(callback)
}

Fortia.emit = function(evt){
	if(!Fortia._eventListenvers[evt]){
		return;
	}

	// Build an array with the rest of the arguments
	var rest = [];
	if(arguments.length > 1){
		for (var i = 1; i < arguments.length; i++) {
			rest.push(arguments[i]);
		};
	}

	for (var i = 0; i < Fortia._eventListenvers[evt].length; i++) {
		Fortia._eventListenvers[evt][i].apply(Fortia._eventListenvers[evt][i], rest);
	};
}
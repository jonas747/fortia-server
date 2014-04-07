var Fortia = Fortia || {};

Fortia.on = function(evt, callback){
	Fortia._eventListenvers = Fortia._eventListenvers || {};
	if(!Fortia._eventListenvers[evt]){
		Fortia._eventListenvers[evt] = new Array();;
	}
	Fortia._eventListenvers[evt].push(callback)
}

Fortia.emit = function(evt, data){
	if(!Fortia._eventListenvers[evt]){
		return;
	}

	for (var i = 0; i < Fortia._eventListenvers[evt].length; i++) {
		Fortia._eventListenvers[evt][i](data);
	};
}
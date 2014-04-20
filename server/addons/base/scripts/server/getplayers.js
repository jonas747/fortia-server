var Fortia = Fortia || {}

Fortia.getPlayers = function(){
	var arr =  _fortiaGetPlayers();
	var players = [] ;
	var length = arr[0].length;
	for (var i = 0; i < length; i++) {
		var id = arr[0][i];
		var x = arr[1][i];
		var y = arr[2][i];
		var z = arr[3][i];
		players.push({id: id, x: x, y: y, z: z});
	};
	return players
}
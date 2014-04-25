var Fortia = Fortia || {}

Fortia.getPlayers = function(){
	var arr =  Fortia._getPlayers();
	var players = [] ;
	for (var i = 0; i < arr.length; i++) {
		var ply = Fortia.getPlayer(arr[i]);
		players.push(ply);
	};
	return players
}
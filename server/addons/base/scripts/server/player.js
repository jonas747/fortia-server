var Fortia = Fortia || {};

Fortia.initPlayer = function(id){
	if(!Fortia.playerList)
		Fortia.playerList = [];

	Fortia.playerList[id] = new Player(id);
}

Fortia.removePlayer = function(id){
	if(!Fortia.playerList){
		Fortia.playerList = [];
		return;
	}
	Fortia.playerList[id] = null;
} 

Fortia.getPlayer = function(id){
	return Fortia.playerList[id];
}
// Update the position of the player on playermove
Fortia.on("playermove", function(playerId, x, y, z){
	var player = Fortia.getPlayer(playerId);
	player.x = x;
	player.y = y;
	player.z = z;
});
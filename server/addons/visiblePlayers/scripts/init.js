addClientJsFile("client.js")

Fortia.on("playerjoin", function(pid){
	var player = Fortia.getPlayer(pid);
	var players = Fortia.getPlayers();
	if(!players)
		console.log("No players returned")
		console.log(players)
		return 

	console.log(players.length)

	var newArray = [];
	// Send the message that this player joined to everyone except himself
	for (var i = 0; i < players.length; i++) {
		var ply = players[i]
		if(ply.id !== player.id){ 
			Fortia.Net.sendUsrMessage(ply, "playerjoin", player)
		}
		newArray.push(ply);
	};
	// Send the player that joined all the other players
	Fortia.Net.sendUsrMessage(player, "otherPlayers", newArray)
})

Fortia.on("playerleave", function(player){
	var players = Fortia.getPlayers()
	if(!players)
		return 

	for (var i = 0; i < players.length; i++) {
		var ply = players[i]
		if(ply.id !== player.id){ 
			Fortia.Net.sendUsrMessage(ply, "playerleave", player)
		}
	};
})

Fortia.on("playermove", function(player){
	var players = Fortia.getPlayers()
	if(!players)
		return 

	// Send the message that this player moved to everyone except himself
	for (var i = 0; i < players.length; i++) {
		var ply = players[i]
		if(ply.id !== player.id){ 
			Fortia.Net.sendUsrMessage(ply, "playermove", player)
		}
	}
})
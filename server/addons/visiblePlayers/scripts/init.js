addClientJsFile("client.js")

Fortia.on("playerjoin", function(player){
	var players = Fortia.getPlayers()
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
			Fortia.Net.sendUsrMessage("playerjoin", player, ply)
		}
		newArray.push(ply);
	};
	// Send the player that joined all the other players
	Fortia.Net.sendUsrMessage("otherPlayers", newArray, player)
})

Fortia.on("playerleave", function(player){
	var players = Fortia.getPlayers()
	if(!players)
		return 

	for (var i = 0; i < players.length; i++) {
		var ply = players[i]
		if(ply.id !== player.id){ 
			Fortia.Net.sendUsrMessage("playerleave", player, ply)
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
			Fortia.Net.sendUsrMessage("playermove", player, ply)
		}
	}
})
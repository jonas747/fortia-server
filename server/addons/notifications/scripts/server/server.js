Fortia.on("playerjoin", function(player){
	console.log("A player joined! ["+player.id+"]")
})

Fortia.on("playerleave", function(player){
	console.log("A player left! ["+player.id+"]")
})
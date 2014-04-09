Fortia.on("playerjoin", function(player){
	var msgData = {name:"frank", poop:"ing"};
	Fortia.Net.sendUsrMessage("something", msgData, player);
})

Fortia.Net.on("poop", function(){
	console.log("Received poop message")
})
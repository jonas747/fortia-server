Fortia.Net.sendMessage("poop", "ass");

Fortia.Net.on("something", function(data){
	console.log("something message! " + data.chatMsg);
})
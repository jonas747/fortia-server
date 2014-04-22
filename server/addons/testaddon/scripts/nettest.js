Fortia.on("playerjoin", function(pid){
	var player = Fortia.getPlayer(pid);
	Fortia.Net.sendUsrMessage(pid, "testmessage", {something: "Heyooo", other: "Nathing"}); 
})

Fortia.Net.on("testreply", function(data){
	console.log(data.rap)
})
Fortia.Net.on("testmessage", function(data){
	console.log(data.something);
	Fortia.Net.sendMessage("testreply", {rap: "my man!"})
})
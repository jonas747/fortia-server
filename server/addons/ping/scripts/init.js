addClientJsFile("client.js", true);

Fortia.Net.on("echo", function(data, sender){
	Fortia.Net.sendUsrMessage(sender, "echoresponse", sender);
})
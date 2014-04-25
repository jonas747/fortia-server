Fortia.on("loaded", function(){
	var pnode = document.createElement("p");
	pnode.innerHTML = "Ping: (no response yet, sure you're connected?)";

	pnode.style.top = "10px";
	pnode.style.left = "150px";
	pnode.style.position = "fixed";

	document.body.appendChild(pnode);

	var waiting = false;
	var timeSent = new Date();

	setInterval(function(){
		if(!waiting){
			//console.log("Send ping request")
			Fortia.Net.sendMessage("echo", "hi");
			waiting = true;
			timeSent = new Date();
		};
	}, 1000);

	Fortia.Net.on("echoresponse", function(data){
		waiting = false;
		var now = new Date();
		var ping = now.getTime() - timeSent.getTime();
		//console.log("Ping: ", ping);
		pnode.innerHTML = "Ping: " + ping + "ms";
	})
 /*
*/
});


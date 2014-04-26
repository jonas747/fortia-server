//addClientJsFile("nettestclient.js", true);

//include("nettest.js");
//include("filetest.js");

Fortia.on("loaded", function(){
	console.log("poop")
	var worker = new Worker("addons/testaddon/scripts/testworker.js");
	worker.postMessage("Hmmmm");

	worker.onmessage = function(data){
		console.log("Message from worker: ", JSON.stringify(data));
		console.log(data[0]);
	}

	for (var i = 0; i < 100000000; i++) {
		var frankenstain = i * Math.random();
	};
	console.log("done")
});

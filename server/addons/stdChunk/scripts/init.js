addClientJsFile("client.js", true);
addClientJsFile("shared.js", true);
addClientJsFile("base64.js", true);
addClientJsFile("lz-string.js", true);

include("shared.js");
include("base64.js");
include("lz-string.js");

Fortia.on("loaded", function(){

	var oldPlayerPositions = [];

	var chunkSize = new Vector3(50, 50, 50);
	var blockScale = 1;

	var chunkWorker = new Worker("addons/stdChunk/scripts/chunkworker.js");
	chunkWorker.onmessage = function(data){
		if(data.name === "getResponse"){
			var player = Fortia.getPlayer(data.pid);
			sendChunk(player, data.chunk, chunkSize, data.position, blockScale);
		}
	}

	chunkWorker.postMessage({
		name: "init",
		size: chunkSize,
		seed: 322,
		worldHeight: 100,
		scale: blockScale,
		blockIds: Fortia.blockIds
	});

	Fortia.on("playerjoin", function(player){
		//Fortia.Net.sendUsrMessage("chunk", {chunk: chunk, size: wgen.size}, player)
		//console.log("Sent a chunk")
	})

	Fortia.on("playermove", function(pid){
		var player = Fortia.getPlayer(pid);

		var oldPos = oldPlayerPositions[player.id];

		if(oldPos == undefined)
			oldPos = new Vector3();

		oldPos = oldPos.clone();
		oldPlayerPositions[player.id] = new Vector3(player.x, player.y, player.z);

		var oldChunkPos = worldToChunk(oldPos, chunkSize, blockScale);
		var newChunkPos = worldToChunk(new Vector3(player.x, player.y, player.z), chunkSize, blockScale);

		if(newChunkPos.x !== oldChunkPos.x || newChunkPos.y !== oldChunkPos.y || newChunkPos.z !== oldChunkPos.z){
			chunkWorker.postMessage({
				name: "getNearby",
				shouldGen: true,
				pos: newChunkPos,
				radius: new Vector3(2, 2, 2),
				pid: pid
			});
		}
	});

	Fortia.on("loaded", function(){
		// Quick spawn gen when server is starting
		//chunkCache.getNearbyChunks(new Vector3(0, 0, 0), new Vector3(2, 2, 2), true)
	});

	function sendChunk(player, chunk, size, position, scale){
		var chunkObj =  {
			chunk: chunk, 
			size: size,
			position: position,
			blockScale: scale
		};
		Fortia.Net.sendUsrMessage(player, "chunk", chunkObj);
	}	
})

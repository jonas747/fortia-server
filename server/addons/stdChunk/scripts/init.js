addClientJsFile("client.js");
addClientJsFile("shared.js");
//addClientJsFile("base64.js");
//addClientJsFile("lz-string.js");

include("shared.js");
//include("base64.js");
//include("lz-string.js");

(function(){

	var oldPlayerPositions = [];

	var wgen = new WorldGen(new Vector3(50, 50, 50), 100, 100, 0.1);
	//var chunk = wgen.genChunk(0, 0, 0);
	var chunks = {};

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

		var oldChunkPos = wgen.worldToChunk(oldPos);
		var newChunkPos = wgen.worldToChunk(new Vector3(player.x, player.y, player.z));

		if(newChunkPos.x !== oldChunkPos.x || newChunkPos.y !== oldChunkPos.y || newChunkPos.z !== oldChunkPos.z){

			// New chunk
			var blockId = newChunkPos.x+"|"+newChunkPos.y+"|"+newChunkPos.z;
			if(!chunks[blockId]){
				// Generate a chunk
				console.log("Generating a chunk")
				var startTime = new Date().getTime();
				var chunk = wgen.genChunk(newChunkPos.x, newChunkPos.y, newChunkPos.z);
				var diff = new Date().getTime() - startTime;
				var blocksPerSecond = Math.round(((wgen.size.x*wgen.size.y*wgen.size.z)/diff)*1000)
				console.log("done generating a chunk, size: "+chunk.length+ " Time spent: "+diff+" Blocks per second: "+blocksPerSecond);
				if(chunk.length < 1){
					//Empty chunk..
					chunks[blockId] = [];
				}
				chunks[blockId] = chunk;
			}
			var chunk = chunks[blockId];
			if(chunk.length < 1){
				return;
			}
			sendChunk(chunk, newChunkPos, player);

		}
	});

	function sendChunk(chunk, position, player){
		//var compressed = compressChunk(chunk)
		var chunkObj =  {
			chunk: chunk, 
			size: wgen.size,
			position: position,
			blockScale: wgen.blockScale
		};
		Fortia.Net.sendUsrMessage(player, "chunk", chunkObj, true)
	}
})();

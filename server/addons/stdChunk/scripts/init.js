addClientJsFile("client.js", true);
addClientJsFile("shared.js", true);
addClientJsFile("base64.js", true);
addClientJsFile("lz-string.js", true);

include("shared.js");
include("base64.js");
include("lz-string.js");
include("chunkcache.js");

(function(){

	var oldPlayerPositions = [];

	//						 size 				seed worldheight scale
	var wgen = new WorldGen(new Vector3(50, 50, 50), 123123, 100, 1);
	var chunkCache = new ChunkCache(wgen);

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
			nearbyChunks = chunkCache.getNearbyChunks(newChunkPos, new Vector3(2, 2, 2), true);
			for (var i = 0; i < nearbyChunks.length; i++) {
				if(nearbyChunks[i] !== "air" && nearbyChunks[i] !== undefined){
					sendChunk(nearbyChunks[i], player)
				}
			};
		}
	});

	Fortia.on("loaded", function(){
		// Quick spawn gen when server is starting
		chunkCache.getNearbyChunks(new Vector3(0, 0, 0), new Vector3(2, 2, 2), true)
	});

	function sendChunk(chunk, player){
		//var compressed = compressChunk(chunk.getCompressedB64(false));
		var compressed = chunk.getCompressedB64(true, chunkCache);
		var chunkObj =  {
			chunk: compressed, 
			size: wgen.size,
			position: chunk.pos,
			blockScale: wgen.blockScale
		};
		console.log("Sending normal compressed chun");
		Fortia.Net.sendUsrMessage(player, "chunk", chunkObj);
	}
})();

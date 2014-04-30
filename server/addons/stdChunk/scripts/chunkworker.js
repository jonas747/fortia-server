include("addons/base/scripts/shared/vector.js");
include("addons/stdWgen/scripts/simplexnoise.js");
include("addons/stdWgen/scripts/init.js");
include("addons/stdChunk/scripts/chunkcache.js");
include("addons/stdChunk/scripts/base64.js");
include("addons/stdChunk/scripts/lz-string.js");


var Fortia = Fortia || {}

var wgen;
var chunkCache;

self.onmessage = function(data){
	switch (data.name){
		case "init":
			Fortia.blockIds = data.blockIds;
			wgen = new WorldGen(new Vector3(data.size.x, data.size.y, data.size.z), data.seed, data.worldHeight, data.scale);
			chunkCache = new ChunkCache(wgen);
			console.log("Chunkworker is now ready!!")
			break;
		case "get":
			var pos = new Vector3(data.pos.x, data.pos.y, data.pos.z);
			var gen = data.shouldGen;
			var pid = data.pid;
			var chunk = chunkCache.getChunk(pos, gen);
			
			self.postMessage({
				name: "getResponse",
				pid: pid,
				chunk: chunk.getCompressedB64(true, chunkCache),
				position: chunk.pos
			});
			
			break;
		case "getNearby":
			var pos = new Vector3(data.pos.x, data.pos.y, data.pos.z);
			var gen = data.shouldGen;
			var radius = data.radius || new Vector3(1, 1, 1);
			var pid = data.pid;	

			var chunks = chunkCache.getNearbyChunks(pos, radius, gen);
			for (var i = 0; i < chunks.length; i++) {
				var chunk = chunks[i];
				if(chunk !== "air" && chunk !== undefined){
					self.postMessage({
						name: "getResponse",
						pid: pid,
						chunk: chunk.getCompressedB64(true, chunkCache),
						position: chunk.pos
					});
				}
			};
			break;
		}
}

var worldToChunk = function(pos, chunkSize, blockScale){
	var clone = pos.clone();
	var scaledBlockSize = chunkSize.multiply(blockScale);

	clone.x = Math.floor(clone.x / scaledBlockSize.x);
	clone.y = Math.floor(clone.y / scaledBlockSize.y);
	clone.z = Math.floor(clone.z / scaledBlockSize.z);

	return clone;
}

var chunkToWorld = function(pos, chunkSize, blockScale){
	var clone = pos.clone();
	var scaledBlockSize = chunkSize.multiply(blockScale);

	clone.x = clone.x * scaledBlockSize.x;
	clone.y = clone.y * scaledBlockSize.y;
	clone.z = clone.z * scaledBlockSize.z;

	return clone;
}
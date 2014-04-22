var chunkCache = {};

Fortia.Net.on("chunk", function(data){
	if(!chunkCache[data.position.x+ "|" + data.position.y + "|" + data.position.z]){
		// Decompress it
		var decompressed = decompressChunk(data.chunk) 

		var posV3 = new Vector3(data.position.x, data.position.y, data.position.z);
		var sizeV3 = new Vector3(data.size.x, data.size.y, data.size.z);

		var worldPos = chunkToWorld( posV3, sizeV3, data.blockScale);
		var scale = new Vector3(data.blockScale, data.blockScale, data.blockScale);


		var blkModel = new BlockModel(decompressed, new Vector3(0, 0, 0), sizeV3, worldPos, scale);
		blkModel.addToWorld();
		chunkCache[data.position.x+ "|" + data.position.y + "|" + data.position.z] = blkModel;
		console.log("added a chunk to world position: ", worldPos)
	}
});
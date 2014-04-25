var chunkCache = {};

Fortia.Net.on("chunk", function(data){
	if(!chunkCache[data.position.x+ "|" + data.position.y + "|" + data.position.z]){
		// Decompress it
		var now = new Date().getTime()
		var decompressed = decompressChunk(data.chunk) ;
		console.log("Time taken deompressing chunk: ", new Date().getTime() - now, "ms");
		
		var posV3 = new Vector3(data.position.x, data.position.y, data.position.z);
		var sizeV3 = new Vector3(data.size.x, data.size.y, data.size.z);
		var scaledSize = new Vector3(data.size.x * data.blockScale, data.size.y * data.blockScale, data.size.z * data.blockScale);

		var worldPos = chunkToWorld( posV3, sizeV3, data.blockScale);
		var scale = new Vector3(data.blockScale, data.blockScale, data.blockScale);

		var blkModel = new BlockModel(decompressed, new Vector3(0, 0, 0), sizeV3, worldPos, scale);
		blkModel.addToWorld();
		chunkCache[data.position.x+ "|" + data.position.y + "|" + data.position.z] = blkModel;
	}
});
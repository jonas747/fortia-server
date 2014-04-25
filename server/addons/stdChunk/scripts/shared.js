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

function compressChunk(chunk){
	var encoded = JSON.stringify(chunk)	
	var compressed = LZString.compressToBase64(encoded);
	return compressed
}

function decompressChunk(chunk){
	var string = LZString.decompressFromBase64(chunk);
	var obj = JSON.parse(string);
	return obj;
}

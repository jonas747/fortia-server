function ChunkCache( wgen ){

	this.chunks = {}
	this.wgen = wgen;
	this.chunkDir = "chunks/"
}

// Loads a chunk from disk
// TODO: Compression
ChunkCache.prototype.loadChunk = function(pos){
	if(!this.chunkIsOnDisk(pos)){
		return
	}

	var file = Fortia.fileLoad(this.chunkDir + ChunkCache.getFileString(pos));
	if(file == "air"){
		this.chunks[ChunkCache.getChunkId(pos)] = "air";
		return
	}

	var rawChunk = JSON.parse(file);
	var chunk = new Chunk(pos, rawChunk, this.wgen.size);


	this.chunks[ChunkCache.getChunkId(pos)] = chunk;
}

// Unloads a chunk
ChunkCache.prototype.unLoadChunk = function(pos){
	if(!this.chunkIsInMemory(pos)){
		return
	}

	this.chunks[ChunkCache.getChunkId(pos)] = null;
}

// Returns wether this chunk exists either on disk or in memory
ChunkCache.prototype.chunkExists = function(pos){
	if(this.chunkIsInMemory(pos) || this.chunkIsOnDisk(pos)){
		return true
	}

	return false
}

// Returns wether the chunk at pos is loaded in memory or not
ChunkCache.prototype.chunkIsInMemory = function(pos){
	if(this.chunks[ChunkCache.getChunkId(pos)]){
		return true
	}
	return false
}

ChunkCache.prototype.chunkIsOnDisk = function(pos){
	if(Fortia.fileExists(this.chunkDir + ChunkCache.getFileString(pos))){
		return true
	}
	return false;
}

// Writes a chunk to disk
// TODO: compression
ChunkCache.prototype.writeChunk = function(pos){
	if(!this.chunkIsInMemory(pos)){
		return
	}

	var	chunk = this.getChunk(pos);
	var serialized = JSON.stringify(chunk.blocks);
	
	// Only air then...
	if(chunk === "air"){
		serialized = "air";
	}
	
	Fortia.fileWrite(this.chunkDir + ChunkCache.getFileString(pos), serialized);
}

// Generates a chunk
ChunkCache.prototype.genChunk = function(pos){
	if(this.chunkExists(pos)){
		return
	}
	var rawChunk = this.wgen.genChunk(pos.x, pos.y, pos.z);
	var chunk = new Chunk(pos, rawChunk, this.wgen.size);
	
	if (rawChunk.length < 1){
		this.chunks[ChunkCache.getChunkId(pos)] = "air";
		return
	}
	this.chunks[ChunkCache.getChunkId(pos)] = chunk;
	console.log("Generated a chunk")
}

// Gets a raw chunk, load it into memory if we have to, if generate is true it will generate the chunk if it does not exist
ChunkCache.prototype.getChunk = function(pos, generate){
	if(this.chunkIsInMemory(pos)){
		return this.chunks[ChunkCache.getChunkId(pos)];
	}else if(this.chunkIsOnDisk(pos)){
		this.loadChunk(pos);
		return this.chunks[ChunkCache.getChunkId(pos)];
	}else if(generate){
		this.genChunk(pos);
		return this.chunks[ChunkCache.getChunkId(pos)];
	}
}

// Returns a chunk, with only the outer blocks, and compressed
ChunkCache.prototype.getClientReadyChunk = function(pos){
	var chunk = this.getChunk(pos);
	if(!chunk){
		return
	}
	
	if(typeof(chunk) === "sting" && chunk === "air"){
		return
	}

	return chunk.getCompressedB64(true);
}

// Checks if a chunk is only consisting of air
ChunkCache.prototype.isAir = function(pos){
	var chunk = this.getChunk(pos);
	if(!chunk){
		return
	}

	if(chunk === "air"){
		return true
	}
	return false
}

// Checks all loaded chunks if there are any players near, unload the chunks that is not near any players
ChunkCache.prototype.checkChunks = function(){

}

// Gets nearby chunks (Loading them id necessary), if generate is true it will generate missing chunks
ChunkCache.prototype.getNearbyChunks = function(pos, radius, generate){
	var radius = radius || new Vector3(3, 3, 3);
	var pos = pos || new Vector3();
	var chunks = new Array();
	for (var x = pos.x - radius.x; x < pos.x + radius.x; x++) {
		for (var y = pos.y - radius.y; y < pos.y + radius.y; y++) {
			for (var z = pos.z - radius.z; z < pos.z + radius.z; z++) {
				var nPos = new Vector3(x, y, z);
				var chunk = this.getChunk(nPos, generate);
				chunks.push(chunk);
			};	
		};	
	};
	return chunks
}

ChunkCache.getFileString = function(pos){
	return pos.x + "." + pos.y + "." + pos.z + ".fortia-chunk";
}

ChunkCache.getChunkId = function(pos){
	return pos.x + "|" + pos.y + "|" + pos.z 
}


function Chunk(pos, blocks, size){
	this.pos = pos || new Vector3();
	this.blocks = blocks || new Array();
	this.size = size || new Vector3();

}

Chunk.prototype.getBlock = function(pos){
	return this.blocks[pos.x + this.size.x * (pos.y + this.size.y * pos.z)];
}
 
// TODO
Chunk.prototype.setBlock = function(pos, block){

}

// Gets only blocks that are visible (plus the chunk edges...)
Chunk.prototype.getOuterBlocks = function(chunkCache){
	processedArr = new Array();
	var num = 0;
	for (var z = 0; z < this.size.z; z++) {
		for (var y = 0; y < this.size.y; y++) {
			InnerLoop:
			for (var x = 0; x < this.size.x; x++) {
				var localPos = new Vector3(x, y, z);
				var block = this.getBlock(localPos);
				// Air..
				if(!block){
					processedArr.push(0);
					continue
				}


				// Check if there are any blocks above, under, next-to etc..
				var above = this.getBlock(new Vector3(localPos.x, localPos.y + 1, localPos.z));
				var below = this.getBlock(new Vector3(localPos.x, localPos.y - 1, localPos.z));
				var x1 = this.getBlock(new Vector3(localPos.x + 1, localPos.y, localPos.z));
				var x_1 = this.getBlock(new Vector3(localPos.x - 1, localPos.y, localPos.z));
				var z1 = this.getBlock(new Vector3(localPos.x, localPos.y, localPos.z + 1));
				var z_1 = this.getBlock(new Vector3(localPos.x, localPos.y, localPos.z - 1));


 				// Check if were at one of the edges, if we are we check the other chunks nearby(if chunkCache is passed as argument)
				if (x === 0 || x === this.size.x - 1 ||
					y === 0 || y === this.size.y - 1 ||
					z === 0 || z === this.size.z - 1){
	 				
					if(chunkCache !== undefined){

						// This is the other chunks were checking, (corners need to check multiple so arrays are used)
						// Redioing this part in the future, very bad code...
		 				var chunkPositions = [];
		 				var blockPositions = [];

		 				if(y === this.size.y - 1){
		 					var nChunkPos = this.pos.clone();
		 					var nBlockPos = localPos.clone();

		 					nChunkPos.y += 1;
		 					nBlockPos.y = 0;

		 					chunkPositions.push(nChunkPos);
		 					blockPositions.push(nBlockPos);

		 					above = true;
		 				}

		 				if(y === 0){
		 					var nChunkPos = this.pos.clone();
		 					var nBlockPos = localPos.clone();
		 				
		 					nChunkPos.y -= 1;
		 					nBlockPos.y = this.size.y -1;
		 				
		 					chunkPositions.push(nChunkPos);
		 					blockPositions.push(nBlockPos);

		 					below = true;
		 				}

		 				if (x === this.size.x - 1) {
		 					var nChunkPos = this.pos.clone();
		 					var nBlockPos = localPos.clone();

		 					nChunkPos.x += 1
		 					nBlockPos.x = 0;

		 					chunkPositions.push(nChunkPos);
		 					blockPositions.push(nBlockPos);
		 					
		 					x1 = true;
		 				}
		 				if (x === 0) {
		 					var nChunkPos = this.pos.clone();
		 					var nBlockPos = localPos.clone();
		 					
		 					nChunkPos.x -= 1;
		 					nBlockPos.x = this.size.x -1;
		 					
		 					chunkPositions.push(nChunkPos);
		 					blockPositions.push(nBlockPos);
		 					x_1 = true;
		 				}
		 				if (z === this.size.x - 1) {
		 					var nChunkPos = this.pos.clone();
		 					var nBlockPos = localPos.clone();
		 					
		 					nChunkPos.z += 1
		 					nBlockPos.z = 0;
		 					
		 					chunkPositions.push(nChunkPos);
		 					blockPositions.push(nBlockPos);
		 					z1 = true;
		 				}
		 				if (z === 0) {
		 					var nChunkPos = this.pos.clone();
		 					var nBlockPos = localPos.clone();
		 					
		 					nChunkPos.z -= 1;
		 					nBlockPos.z = this.size.z -1;
		 					
		 					chunkPositions.push(nChunkPos);
		 					blockPositions.push(nBlockPos);
		 					z_1 = true;
		 				};

		 				for (var i = 0; i < chunkPositions.length; i++) {
		 					var nChunkPos = chunkPositions[i]
		 					var nBlockPos = blockPositions[i]

			 				var nChunk = chunkCache.getChunk(nChunkPos, true);
			 				if(nChunk === "air"){
		 						processedArr.push(block);
		 						num++
		 						continue InnerLoop;
		 					}

		 					if(!nChunk.getBlock(nBlockPos)){
		 						processedArr.push(block);
		 						num++
		 						continue InnerLoop;
			 				}
		 				};

 					}
				}


				// If not we see the block
				if(!above || !below || !x1 || !x_1 || !z1  || !z_1){
					processedArr.push(block);
					num++
					continue
				}

				// Completely covered, we dont see it
				processedArr.push(0);
			};
		};
	};
	console.log("Outer number of blocks: ", num)
	return processedArr;
}

Chunk.prototype.getCompressedB64 = function(outerOnly, chunkCache){
	// Check if its in the cache, if not compress it again
	if((!this.compressedB64Full && !outerOnly) || (!this.compressedB64Outer && outerOnly) || this.isDirty){
		var blocks = this.blocks;
		if(outerOnly){
			blocks = this.getOuterBlocks(chunkCache);
		}

		var encoded = JSON.stringify(blocks);
		var compressed = LZString.compressToBase64(encoded);

		if(outerOnly){
			this.compressedB64Outer = compressed;
		}else{
			this.compressedB64Full = compressed;
		}
		return compressed
	}

	if(outerOnly){
		return this.compressedB64Outer
	}else{
		return this.compressedB64Full
	}
}
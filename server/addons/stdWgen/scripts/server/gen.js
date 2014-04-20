function WorldGen(size, seed, worldHeight, blockScale){
	this.size = size || new Vector3(100, 100, 100);
	this.seed = seed || Math.random()*new Date().getMilliseconds();
	this.worldHeight = worldHeight || 100;
	this.blockScale = blockScale || 1;

	this.smoothNess = 0.1;

	this.noiseGen = new SimplexNoise();

	this.genChunk = function(pX, pY, pZ){

		var chunkWorldPos = this.chunkToWorld(new Vector3(pX, pY, pZ));
		// Check if were above the world hegight
		if(chunkWorldPos.y > this.worldHeight){
			return [];
		}

		var blocks = [];
		var generatedBlock = false;
		for (var x = 0; x < this.size.x; x++) {
			for (var y = 0; y < this.size.y; y++) {
				for (var z = 0; z < this.size.z; z++) {
					// If i dont switch the x and z values i get some issues on the client side, the chunk are rotated 90 dagrees then
					// Im not good at math, dunno why it's happening
					var realPos = new Vector3(z, y, x);
					realPos = realPos.multiply(this.blockScale);
					var worldPos = chunkWorldPos.add(realPos);
					var noisePos = worldPos.multiply(this.smoothNess);
					var noise = this.noiseGen.noise3d(noisePos.x, noisePos.y, noisePos.z);
					var life = 1 - 2*(worldPos.y/this.worldHeight)
					life += noise;
					if(life > 0){
						var btype = Math.floor(Math.random()*4)
						blocks.push(Fortia.blockIds["rock"+btype]);
						generatedBlock = true;
					}else{
						blocks.push(0);
					}
				};
			};
		};
		if(!generatedBlock){
			return [];
		}
		return blocks;
	}

	this.worldToChunk = function(pos){
		return worldToChunk(pos, this.size, this.blockScale);
	}

	this.chunkToWorld = function(pos){
		return chunkToWorld(pos, this.size, this.blockScale);
	}

	return this;
}
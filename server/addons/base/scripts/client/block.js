var Fortia = Fortia || {};

(function(){
	Fortia.blockColors = {};
	Fortia.blockIds = {};
	Fortia.blockIdToColors = [];

	/**
	Registers a new block kind
	@function registerBlockType
	@param {String} name - The name of the block
	@param {Number} color - The color of the block
	**/
	Fortia.registerBlockType = function(name, color){
		Fortia.blockColors[name] = color;
	}
	
	Fortia.Net.on("blockids", function(ids){
		console.log("Block id's", ids)
		Fortia.blockIds = ids;
		for (var key in ids) {
			if (ids.hasOwnProperty(key)) {
				var val = ids[key]
				Fortia.blockIdToColors[val] = Fortia.blockColors[key]	
			}
		}
		// Send the colors to the blockmodel worker
		FortiaEngine.blkWorker.postMessage({
			name: "colors",
			colors: Fortia.blockIdToColors
		});
	});
})();
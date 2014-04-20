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
		id = newId();
		Fortia.blockColors[name] = color;
		Fortia.blockIds[name] = id;
		Fortia.blockIdToColors[id] = color;
	}

	var curId = 1;
	function newId(){
		curId++;
		return curId;
	}

	Fortia.on("playerjoin", function(player){
		// send this player all the block id's
		Fortia.Net.sendUsrMessage("blockids", Fortia.blockIds, player);
	});
})();
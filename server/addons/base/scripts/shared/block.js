var Fortia = Fortia || {};

Fortia.getBlockColor = function(id){
	if(typeof(id) == "string"){
		return Fortia.blockColors[id] || 0;
	}else{
		return Fortia.blockIdToColors[id] || 0;
	}
} 
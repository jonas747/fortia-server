var Fortia = Fortia || {};

// Returns the source of a script
Fortia.getSource = function(path){
	for (var i = 0; i < FortiaEngine.scripts.length; i++) {
		var script = FortiaEngine.scripts[i];
		if(script.Name === path){
			return script.Script;
		}		
	}
}
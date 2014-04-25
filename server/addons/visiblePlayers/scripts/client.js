(function(){
	var PlayerBoxes = [];
	var material = new THREE.MeshPhongMaterial({
		vertexColors: THREE.FaceColors,
		blending: THREE.AdditiveBlending,
		shininess: 80,
		color: new THREE.Color( 0xff0000 ),
	});

	Fortia.Net.on("playerjoin", function(data){
		console.log("A player joined! :D")
		// Add a box
		addBox(data)
	})

	Fortia.Net.on("playerleave", function(data){
		console.log("A player left! :'(")
		removeBox(data)
	})

	Fortia.Net.on("otherPlayers", function(data){
		for (var i = 0; i < data.length; i++) {
			var ply = data[i]
			if(ply.id !== Fortia.GetLocalPlayerId()){
				console.log("Added a player box")	
				addBox(ply);
			}
		};
	})

	Fortia.Net.on("playermove", function(player){
		updateBox(player)
	})

	function addBox(player){
		var geometry = new THREE.BoxGeometry( 1, 1, 1);
		var mesh = new THREE.Mesh( geometry, material );
		mesh.position.x = player.x;
		mesh.position.y = player.y;
		mesh.position.z = player.z;

		PlayerBoxes[player.id] = mesh

		FortiaEngine.scene.add(mesh);
	}

	function removeBox(player){
		var mesh = PlayerBoxes[player.id]
		FortiaEngin.scene.remove(remove);
		PlayerBoxes[player.id] = null;
	}

	function updateBox(player){
		var mesh = PlayerBoxes[player.id];
		mesh.position.x = player.x;
		mesh.position.y = player.y;
		mesh.position.z = player.z;
	}
})()
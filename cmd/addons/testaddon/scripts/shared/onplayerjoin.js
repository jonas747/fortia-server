Fortia.on("playerjoin", function(data){
	console.log("A player joined! ["+data.id+"]");
})
 
Fortia.on("playerleave", function(data){
	console.log("A player left :(! ["+data.id+"]");
})

function extendedEntity(){
	Entity.call(this);
}

console.log(new extendedEntity().name);
function Vector3(x, y, z){
	this.x = x || 0;
	this.y = y || 0;
	this.z = z || 0;	
}

Vector3.prototype.clone = function(){
	return new Vector3(this.x, this.y, this.z);
}

Vector3.prototype.multiply = function(other){
	var clone = this.clone();
	if(typeof(other) === "number"){
		clone.x *= other;
		clone.y *= other;
		clone.z *= other;
	}else{
		// Assume other is vector
		clone.x *= other.x;
		clone.y *= other.y;
		clone.z *= other.z;
	}
	return clone;
}

Vector3.prototype.add = function(other){
	var clone = this.clone();
	if(typeof(other) === "number"){
		clone.x += other;
		clone.y += other;
		clone.z += other;
	}else{
		clone.x += other.x;
		clone.y += other.y;
		clone.z += other.z;
	}
	return clone;
}

//if(_fortiaclient){
//	Vector3 = THREE.Vector3;	
//}
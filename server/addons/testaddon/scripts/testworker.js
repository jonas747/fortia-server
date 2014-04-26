self.onmessage = function(data){
	console.log("Received some data in worker: ", data)
	for (var i = 0; i < 100000000; i++) {
		var frankenstain = i * Math.random();
	};
	console.log("Done!");
	self.postMessage("heyooo!");
	self.postMessage([1, 5, 1000]);
}
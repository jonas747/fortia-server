Fortia._workers = [];
Fortia._onWorkerMsg = function(id, data){
	var worker = Fortia._workers[id];
	if (!worker){
		return;
	}

	worker.onmessage(data);
}


function Worker(path){
	this.id = Fortia._newWorker(path);
	Fortia._workers[this.id] = this;
}

Worker.prototype.postMessage = function(data){
	Fortia._postWorkerMessage(this.id, data);
}

Worker.prototype.onmessage = function(){};
(function(){

	function paragraph(text){
		var pnode = document.createElement("p");
		//var tnode = document.createTextNode(text);
		//pnode.appendChild(tnode)
		pnode.innerHTML = text;
		return pnode;
	}

	var newDiv = document.createElement("div"); 
	newDiv.style.top = "150px";
	newDiv.style.position = "fixed";
	var lines = []
	lines.push(paragraph(""));
	lines.push(paragraph("")); 
	lines.push(paragraph("")); 
	lines.push(paragraph("")); 
	lines.push(paragraph("")); 
	lines.push(paragraph("")); 

	for (var i = 0; i < lines.length; i++) {
		newDiv.appendChild(lines[i]);
	};

	// add the newly created element and its content into the DOM 
	document.body.appendChild(newDiv);

	var oldLog = console.log;
	var oldDebug = console.debug;
	var oldError = console.error;

	console.error = function(){
		log.apply(this, arguments);
		oldError.apply(this, arguments);
	}
	console.debug = function(){
		log.apply(this, arguments)
		oldDebug.apply(this, arguments);
	}
	console.log = function(){
		log.apply(this, arguments)
		oldLog.apply(this, arguments)
	}

	function log(){
		var completeString = "";
		for (var i = 0; i < arguments.length; i++) {
			completeString += JSON.stringify(arguments[i]) + " ";
		};

		var oldLine
		for (var i = lines.length - 1; i >= 0; i--) {
			var temp = oldLine;
			oldLine = lines[i].innerHTML;
			lines[i].innerHTML = temp;
		};
		lines[lines.length-1].innerHTML = completeString;
	}
})()
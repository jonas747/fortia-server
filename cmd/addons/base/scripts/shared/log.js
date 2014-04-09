if(!_fortiaclient){
	console.log = _fortialog;
	console.error = _fortiaerror;
	console.debug = _fortiadebug;
	console.info = console.log;
}
console.log("console.log test");
console.info("console.info test");
console.error("console.error test")
console.debug("console.debug test")
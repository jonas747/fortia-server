function Console(){
}

Console.prototype.log = _logInfo;
Console.prototype.debug = _logDebug;
Console.prototype.error = _logError;

console = new Console();
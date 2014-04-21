console.debug("Does something folder exists?");
console.debug(Fortia.fileExists("something"));

console.debug("Creating directory something");
Fortia.fileCreateDir("something");

console.debug("Does something dir exist now?");
console.debug(Fortia.fileExists("something"))

console.debug("Does something/something.txt exist?");
console.debug(Fortia.fileExists("something/something.txt"));

console.debug("Creating file something/something.txt and filling it with something");
Fortia.fileWrite("something/something.txt", "something is suppoed to\nbe\nin\nhere");

console.debug("Does something/something.txt exist now?");
console.debug(Fortia.fileExists("something/something.txt"));

console.debug("Reading file something/something.txt");
console.debug(Fortia.fileRead("something/something.txt"));
#Fortia game server

Fortia is a multiplayer game im developing. Basically mix garrysmod with minecraft and this is what you get.

It's in very early stages so dont expect anything, and the code is probably very bad.

###Current status:

 - Addons:
	- Scripts
		- Loads and executes the correct scripts
		- Sends scripts to clients asking for it
	- Other resources (blockmodels, sounds and so on)
		- Not implemented yet
 - Physics:
	- Not started on, Need to find a decent physics engine that works for go

 - API:
	- Sending and receiving messages between the server and clients
	- Getting a list of players (Redoing this later)
	- Networking and regular events
	- Console.log console.error console.debug works
	- basic fileapi: Fortia.fileRead fileWrite fileExists fileCreateDir

- Network:
	- Clients can connect
	- Clients can send and receive messages in the protobuf format (Javascript objects will be serialized with json)
	- Networking events works

#Technical information: 
	
##Addons

Addons are structured like so:
- addon.json - Addon information
- recsurces - Folder containing addtional recourses that will be downloaded when you connect
	- resources.json - Information about the resources
	- TBD
- scripts - Folder containing scripts
	- init.js
		- This script will be executed on server startup, use include(), and addClientJsFile() in this file.

More detailed information:

	addon.json:

	{
		"name": "Some addon",
		"version": "alpha1",
		"type": "gamemode",
		"author": "jonas747",
		"license": "bsd"
	}

	resources/resources.json

	{
	}

##BlockModel format:

Everything you see(except the ui) in fortia is created from BlockModels, a blockmodel is essentially an array blocks and animations.

Blockmodels are stored in the protocol buffer format (messages.BlockModel specifically), chunks are an example of using blcokmodels dynamicly.

###Structures
See the protobuf files for blockmodel

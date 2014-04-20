#Fortia game server

Fortia is a multiplayer game im developing. Basically mix garrysmod with minecraft and this is what you get.

It's in very early stages so dont expect anything, and the code is probably very bad.

Current status:
	
	Addons:
		Scripts
			- Loads and executes the corrects scripts
		Other resources (blockmodels, sounds and so on)
			- Not implemented yet
	Physics:
		Not started on, thinking of bullet for the engine

	API:
		DONE:
			- Sending and receiving messages between the server and clients
			- Getting a list of players (Redoing this later)
			- Networking and regular events
			- Console.log and the like


#Technical information: 
	
##Status


Addons are structured like so:
	addon.json - Addon information
	recsurces - Folder containing addtional recourses that will be downloaded when you connect
		resources.json - Information about the resources
		textures - Folder containing textures
			texture.png
		models - Folder containing blockmodels (fortia storing models made out of blocks)
			model.blkmdl
	scripts
		server - Scripts that are executed on the server
		client - Scripts that are executed on the client
		shared - Scripts that are executed on both the the client and the server

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

BlockModel format:

Everything you see(except the ui) in fortia is created from BlockModel, a blockmodel is essentially an array blocks and animations.
blockmodels are stored in protocol buffer(messages.BlockModel specifically), chunks are an example of using blcokmodels dynamicly.

BlockModel
	Represents the blockmodel
	 - Blocks []Block
	 - Animations []AnimactionSequenze
	 - Position vector3f

A Block:
	Blocks have 3(maybe 4, if i add color) fields:
		- Type int - The type of the block, if the type is not found, it defaults to blocknotfound block.
		- Groups []string - The groups this block belongs to, used for animations.
		- Position Vector3(int) - The (local)position of the block
		- (maybe) COlor Vector3(uint8) - The color of the block

	Unsure about:
		Making the size of blocks fixed or not. (leaning towards fixed to make things simpler, atleast in the start)
		Having block types or just colors. (color would then be 4, 1byte channels (rgba))
		 - If i have block types i can add custom handlers/effects for certain blocks in scripts. For exmaple making walking on a certain block deadly.

AnimationStep
	This is a single animations step in an animation sequenze
	the fields:
	 - Group string (The group of of blocks thats being moved)
	 - Time int (The time since start in millisecond that this step is switched to)
	 - NewPos Vector3int ( The new position for that block group)

Animation Sequenze
	An array of animnation steps + The animation name
	 - Steps []AnimationStep ( The steps)
	 - Name string ( The animation name)



My todo list:

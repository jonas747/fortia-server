// Shared
addClientJsFile("shared/event.js");
addClientJsFile("shared/basicBlock.js");
addClientJsFile("shared/entity.js");
addClientJsFile("shared/net.js");
addClientJsFile("shared/block.js");
addClientJsFile("shared/player.js");
addClientJsFile("shared/vector.js");
addClientJsFile("shared/getSource.js");

// Client only
addClientJsFile("client/block.js");
addClientJsFile("client/blockmodel.js");
addClientJsFile("client/net.js");
addClientJsFile("client/playerid.js");

// Server scripts

// Shared
include("shared/event.js");
include("shared/basicBlock.js");
include("shared/entity.js");
include("shared/net.js");
include("shared/block.js");
include("shared/player.js");
include("shared/vector.js");

// Server only
include("server/block.js")
include("server/blockmodel.js")
include("server/getplayers.js")
include("server/net.js")
include("server/player.js")

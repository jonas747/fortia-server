// Shared
addClientJsFile("shared/event.js", true);
addClientJsFile("shared/basicBlock.js", true);
addClientJsFile("shared/entity.js", true);
addClientJsFile("shared/net.js", true);
addClientJsFile("shared/block.js", true);
addClientJsFile("shared/player.js", true);
addClientJsFile("shared/vector.js", true);

// Client only
addClientJsFile("client/block.js", true);
addClientJsFile("client/blockmodel.js", true);
addClientJsFile("client/net.js", true);
addClientJsFile("client/playerid.js", true);
addClientJsFile("client/getSource.js", true);

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

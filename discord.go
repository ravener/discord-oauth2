package discord

import (
	"golang.org/x/oauth2"
)

// All scope constants that can be used.
const (
	ScopeIdentify                   = "identify"
	ScopeBot                        = "bot"
	ScopeEmail                      = "email"
	ScopeGuilds                     = "guilds"
	ScopeGuildsJoin                 = "guilds.join"
	ScopeConnections                = "connections"
	ScopeGroupDMJoin                = "gdm.join"
	ScopeMessagesRead               = "messages.read"
	ScopeRPC                        = "rpc"                    // Whitelist only
	ScopeRPCAPI                     = "rpc.api"                // Whitelist only
	ScopeRPCNotificationsRead       = "rpc.notifications.read" // Whitelist only
	ScopeWebhookIncoming            = "webhook.Incoming"
	ScopeApplicationsBuildsUpload   = "applications.builds.upload" // Whitelist only
	ScopeApplicationsBuildsRead     = "applications.builds.read"
	ScopeApplicationsStoreUpdate    = "applications.store.update"
	ScopeApplicationsEntitlements   = "applications.entitlements"
	ScopeRelationshipsRead          = "relationships.read" // Whitelist only
	ScopeActivitiesRead             = "activities.read"    // Whitelist only
	ScopeActivitiesWrite            = "activities.write"   // Whitelist only
	ScopeApplicationsCommands       = "applications.commands"
	ScopeApplicationsCommandsUpdate = "applications.commands.update"
)

// Endpoint is Discord's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://discord.com/api/oauth2/authorize",
	TokenURL:  "https://discord.com/api/oauth2/token",
	AuthStyle: oauth2.AuthStyleInParams,
}

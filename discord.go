package discord

import (
  "golang.org/x/oauth2"
)

// All scope constants that can be used.
const (
  ScopeIdentify = "identify"
  ScopeBot = "bot"
  ScopeEmail = "email"
  ScopeGuilds = "guilds"
  ScopeGuildsJoin = "guilds.join"
  ScopeConnections = "connections"
  ScopeGroupDMJoin = "gdm.join"
  ScopeMessagesRead = "messages.read"
  ScopeRPC = "rpc"
  ScopeRPCAPI = "rpc.api"
  ScopeRPCNotificationsRead = "rpc.notifications.read"
  ScopeWebhookIncoming = "webhook.Incoming"
)

// Endpoint is Discord's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
  AuthURL:   "https://discordapp.com/api/oauth2/authorize",
  TokenURL:  "https://discordapp.com/api/oauth2/token",
  AuthStyle: oauth2.AuthStyleInParams,
}

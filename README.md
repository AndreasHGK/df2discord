# df2discord

df2discord is a small library that will log the dragonfly chat to a discord webhook.

Using it in your projects is very simple, you only need the following code for basic functionality:
```go
import (
    "github.com/andreashgk/df2discord"
    "github.com/df-mc/dragonfly/server/player/chat"
)
chat.Global.Subscribe(df2discord.NewWebhook("https://discord.com/api/webhooks/CHANNEL/WEBHOOK"))
```
**Be sure to replace the link with your actual webhook link!** Now the server will log the chat history to your discord server. You can also subscribe the same logger to any other chat.
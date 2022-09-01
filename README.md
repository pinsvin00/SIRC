# SIRC
Simple IRC client and server. Message fetching based on long polling.
IRC server can handle the following commands
```
/quit reason - disconnects with IRC room and displays reason of disconnection to all room members 
/colour new_colour - changes the colour of nickname visible in IRC room
/nick new_nickname - changes the nick visible in the IRC room and sends message to all clients that notifies of changing the nick.
```

# Running the app
## Backend
```
If you want to use the server locally, without any runner just run:
go run .

If you want to build the server binary run
go build . 
And then configure your systemd service


If want to change the server messages content, enter the translations.go file and there supply your own message content.
Consider that messages need to have proper printf Formatting. Below are provided the formats of each server message.
- QUIT MESSAGE (%s %s) (nickname, reason)
- NICK CHANGE (%s %s) (old_nickname, new_nickname)
```

## Frontend

Supply frontend files to your HTTP server, if want to use custom server address (which means, address that is different from your HTTP server that serves your frontend), go to fetcher.js and change const ip to match your servers endpoint.

You don't need to compile your frontend, everything is based on blazingly fast vanilla.js

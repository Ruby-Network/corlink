# Corlink Discord Bot

This is the Discord bot for the Corlink project. 
The Bot interacts with our [Licensing Server](../licensing/)
and provides a way to manage the user's API keys.
The specific API routes used are below:

## Installation

- The installation can be found [here](./setup.md)

## Routes and Usage
- `POST /create-user` - Create a new user and API key. Docs [here](../licensing/README.md#post-create-user)
    - Usage: `/admin createuser [username]`. The code can be located [here](./commands/createuser.rb)
--- 
- `POST /delete-user` - Delete a user and their API key. Docs [here](../licensing/README.md#post-delete-user)
    - Usage: `/admin deleteuser [username]`. The code can be located [here](./commands/deleteuser.rb)   
---
- `POST /update-user` - Update a user's API key. Docs [here](../licensing/README.md#post-update-user)
    - Usage: `/admin updateuser [username]`. The code can be located [here](./commands/updateuser.rb)
    - User Usage: `/updateapikey`. The code can be located [here](./commands/updateapikey.rb)
---
- `POST /get-user` - Get a user's API key. Docs [here](../licensing/README.md#post-get-user)
    - User usage: `/getapikey`. The code can be located [here](./commands/getapikey.rb)
---

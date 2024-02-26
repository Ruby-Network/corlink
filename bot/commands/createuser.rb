@bot.application_command(:admin).subcommand(:createuser) do |event|
  #make sure the user is an admin or the owner of the discord server 
  if !event.user.permission?(:administrator) && event.user.id != ENV['OWNER_ID'].to_i 
    event.respond(content: "You don't have permission to use this command.", ephemeral: true)
  else
    username = event.options['username']
    begin
      user_id = event.server.members.find { |member| member.username == username }.id
      user = event.server.members.find { |member| member.username == username }
    rescue 
      event.respond(content: "Error: something went wrong. (The user most likely doesn't exist)", ephemeral: true)
    end
    if user.bot_account?
      event.respond(content: "You can't create an API key for a bot account.")
    else
      url = ENV["LICENSING_SERVER_URL"]
      if url[-1] != "/"
        url += "/"
      end 
      event.respond(content: "Creating user...", ephemeral: true)
      http = HTTParty.post(url + "create-user", headers: { 'Content-Type' => 'application/json', 'Authorization' => "Bearer " + ENV["LICENSING_SERVER_KEY"], "User" => user_id.to_s })
      if http.code == 200
        response = JSON.parse(http.body)
        api_key = response["message"]
        event.send_message(content: "User created! Have the user run `/getapikey` to get their API key.", ephemeral: true)
      else 
        event.send_message(content: "Error: something went wrong.", ephemeral: true)
      end
    end
  end
end
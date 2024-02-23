@bot.application_command(:admin).subcommand(:deleteuser) do |event|
  #make sure the user is an admin or the owner of the discord server
  event.defer
  if !event.user.permission?(:administrator) && event.user.id != ENV['OWNER_ID'].to_i 
    event.send_message(content: "You don't have permission to use this command.", ephemeral: true)
  else
    username = event.options['username']
    begin
      user_id = event.server.members.find { |member| member.username == username }.id
      user = event.server.members.find { |member| member.username == username }
    rescue 
      event.send_message(content: "Error: something went wrong. (The user most likely doesn't exist)", ephemeral: true)
    end
    if user.bot_account?
      event.send_message(content: "You can't delete an API key for a bot account.")
    else
      url = ENV["LICENSING_SERVER_URL"]
      if url[-1] != "/"
        url += "/"
      end 
      http = HTTParty.post(url + "delete-user", headers: { 'Content-Type' => 'application/json', 'Authorization' => "Bearer " + ENV["LICENSING_SERVER_KEY"], "User" => user_id.to_s })
      if http.code == 200
        response = JSON.parse(http.body)
        api_key = response["message"]
        event.send_message(content: "User deleted!", ephemeral: true)
      elsif http.code == 404
        event.send_message(content: "Error: the user doesn't exist.", ephemeral: true)
      else
        event.send_message(content: "Error: something went wrong.", ephemeral: true)
      end
    end
  end
end

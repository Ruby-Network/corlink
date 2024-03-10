@bot.application_command(:admin).subcommand(:generatekey) do |event|
  #make sure the user is an admin or the owner of the discord server
  event.defer
  if !event.user.permission?(:administrator) && event.user.id != ENV['OWNER_ID'].to_i 
    event.send_message(content: "You don't have permission to use this command.", ephemeral: true)
  else
    username = event.options['username']
    url = ENV["LICENSING_SERVER_URL"]
    if url[-1] != "/"
      url += "/"
    end 
    http = HTTParty.post(url + "generate", headers: { 'Content-Type' => 'application/json', 'Authorization' => "Bearer " + ENV["LICENSING_SERVER_KEY"], "User" => user_id.to_s })
    if http.code == 200
      response = JSON.parse(http.body)
      key = response["message"]
      event.send_message(content: "Generated key: " + key, ephemeral: true)
    end
  end
end

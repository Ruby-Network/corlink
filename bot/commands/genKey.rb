@bot.register_application_command(:generatekey, 'Generate an API key', server_id: ENV['GUILD_ID']) do |cmd|
end

@bot.application_command(:generatekey) do |event|
  #make sure the user is an admin or the owner of the discord server
  event.defer
  username = event.options['username']
  url = ENV["LICENSING_SERVER_URL"]
  if url[-1] != "/"
    url += "/"
  end 
  http = HTTParty.post(url + "generate", headers: { 'Content-Type' => 'application/json', 'Authorization' => "Bearer " + ENV["LICENSING_SERVER_KEY"], "User" => username })
  if http.code == 200
    response = JSON.parse(http.body)
    key = response["message"]
    event.send_message(content: "Generated key: " + key, ephemeral: true)
  end
end

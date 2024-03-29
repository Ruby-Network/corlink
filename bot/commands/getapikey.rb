@bot.register_application_command(:getapikey, 'Get your API key!', server_id: ENV['GUILD_ID']) do |cmd|
end

@bot.application_command(:getapikey) do |event|
  event.defer
  user_id = event.user.id
  url = ENV["LICENSING_SERVER_URL"]
  if url[-1] != "/"
    url += "/"
  end 
  http = HTTParty.post(url + "get-user", headers: { 'Content-Type' => 'application/json', 'Authorization' => "Bearer " + ENV["LICENSING_SERVER_KEY"], "User" => user_id.to_s })
  if http.code == 200
    response = JSON.parse(http.body)
    api_key = response["message"]
    event.send_message(content: "Your API key is: `" + api_key + "`", ephemeral: true)
  elsif http.code == 404
    event.send_message(content: "You don't have an API key! Regsiter or open a ticket to get started", ephemeral: true)
  else 
    event.send_message(content: "Error: something went wrong.", ephemeral: true)
  end
end

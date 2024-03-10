@bot.register_application_command(:admin, 'Create a user and return their API key', server_id: ENV['GUILD_ID']) do |cmd|
  cmd.subcommand('createuser', 'Create a user and return their API key') do |sub|
    sub.user('username', 'The user to create', required: true)
  end
  cmd.subcommand('deleteuser', 'Create a user and return their API key') do |sub|
    sub.user('username', 'The username of the user', required: true)
  end  
  cmd.subcommand('updateuser', 'Create a user and return their API key') do |sub|
    sub.user('username', 'The username of the user', required: true)
  end
  cmd.subcommand('generatekey', 'Generate a key') do |sub|
  end
end

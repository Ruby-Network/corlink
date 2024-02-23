require 'discordrb'
require 'colorize'
require 'dotenv'
require 'httparty'
require './utils.rb'
Dotenv.load
@token = ENV['DISCORD_TOKEN']
@bot = Discordrb::Bot.new(token: @token, intents: [:server_members])
@bot.ready do |event|
  @bot.update_status('online', "the API", nil, 0, false, 3)
end
activity = Discordrb::Activity.new("PLAYING", "with Ruby")
@utils = Utils.new
Dir["#{File.dirname(__FILE__)}/commands/*.rb"].each { |file| require file }
@bot.run

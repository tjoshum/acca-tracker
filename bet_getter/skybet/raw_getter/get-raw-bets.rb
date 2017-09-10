require 'rubygems'
require 'highline/import'
require 'mechanize'
require 'nokogiri'

OPEN_BET_PAGE = "https://www.skybet.com/secure/identity/m/history/betting?settled=N"

def get_username
  STDERR.print "Username: "
  $stdin.gets.strip
end

def get_pin
  STDERR.ask("PIN: ") {|q| q.echo = false}
end

class LoginFailureException < RuntimeError

end

def login_and_get_open_bets
  STDERR.puts("Enter SkyBet Credentials")
  username = get_username
  pin = get_pin

  mech_agent = Mechanize.new { | agent |
    agent.follow_meta_refresh = true
  }
  mech_agent.get("https://www.skybet.com/secure/identity/auth?consumer=skybet") do | home_page |
    # Login
    user_home = home_page.form_with(:name => nil) do | form |
      form.username = username
      form.pin = pin
    end.submit
    if (login_failed(user_home))
      puts "Login failed"
      raise LoginFailureException
    end

    raw_bodies = ""
    # Try to get the bet page
    mech_agent.get(OPEN_BET_PAGE) do | bet_page |
      raw_bodies += bet_page.body
      return raw_bodies
    end
  end
end

def login_failed(user_home)
  user_home.body.include?("We couldn't recognise the details you entered. Please try again")
end

puts login_and_get_open_bets

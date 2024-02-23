FROM ruby:3.2.2

RUN bundle config --global frozen 1

WORKDIR /usr/src/app

COPY Gemfile Gemfile.lock ./ 
RUN bundle install

COPY . .

CMD ["bundle", "exec", "ruby", "main.rb"]

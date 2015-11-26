# apidemic [![Build Status](https://travis-ci.org/gernest/apidemic.svg)](https://travis-ci.org/gernest/apidemic)

Apidemic is a service for generating fake JSON API response. You first register the sample JSON response, and apidemic will serve that response with random fake data.

This is experimental, so take it with a grain of salt.

# Motivation
I got bored with hardcoding the sample json api response in tests. I you know golang, you can benefit by using the library, I have included a router that you can use to run disposable servers in your tests.

# Instalation

You can downloaad the binaries for your respective operating system  [Download apidemic](https://github.com/gernest/apidemic/releases/latest)

Then put the downloaaded binary somewhere in your system path.

Alternatively, if you have golang installed

	go get github.com/gernest/apidemic/cmd/apidemic
	

Now you can start the service like this 

	apidemic start
	
This will run a service at localhost default port is 3000, you can change the port by adding a flag `--port=YOUR_PORT_NUMBER`


# How to use
Lets say you expect a response like this

```json
{
  "name": "anton",
  "age": 29,
  "nothing": null,
  "true": true,
  "false": false,
  "list": [
      "first",
      "second"
    ],
  "list2": [
    {
      "street": "Street 42",
      "city": "Stockholm"
      },
    {
      "street": "Street 42",
      "city": "Stockholm"
      }
    ],
  "address": {
    "street": "Street 42",
    "city": "Stockholm"
  },
  "country": {
    "name": "Sweden"
  }
}
```


If you have alreasy started apidemic server you can register that response by making a POST request to the `/register` path. Passing the json body of the form.

```json
{
  "endpoint": "test",
  "payload": {
    "name: first_name": "anton",
    "age: digits_n,max=2": 29,
    "nothing:": null,
    "true": true,
    "false": false,
    "list:word,max=3": [
      "first",
      "second"
    ],
    "list2": [
      {
        "street:street": "Street 42",
        "city:city": "Stockholm"
      },
      {
        "street": "Street 42",
        "city": "Stockholm"
      }
    ],
    "address": {
      "street:street": "Street 42",
      "city:city": "Stockholm"
    },
    "country": {
      "name:ountry": "Sweden"
    }
  }

}
```

See the annotation tags on the payload. Example if I want to generate full name  for field name I will just do `"name:full_name"`.

I f your post request is submitted you are good to ask for the response with fake values. Just do a get request for the endppint you registered

So, every GET call to `/api/test` will return the api response with fake data.

# Tags
Apidemic uses tags to annotate what kind of fake data to generate and also control different requrements of fake data.

You add tags to object keys. For instance if you have a json object `{ "user_name": "gernest"}`. If you want to have fake username  then you can annotate the key by addimg user_name tag like this `{ "user_name:user_name": "gernest"}`.

So,  json keys can be annotated by adding the `:` symbol then followed by comma separated list of tags. The firs entry after `:` is for the tag type, the following entries are in the form `key=value` which will be the extra information to fine tune your fake data. Please see the example above to see how tags are used.

Apidemic comes shiped with a large number of tags, meaning it is capable to geerate a wide range of fake information.

These are different tags that generate different fake data

 Tag | Details( data generated)
------|--------
brand | brand 
 character | character 
 characters | characters 
 characters_n | characters n 
 city | city 
 color | color 
 company | company 
 continent | continent 
 country | country 
 credit_card_num | credit card num 
 currency | currency 
 currency_code | currency code 
 day | day 
 digits | digits 
 digits_n | digits n 
 domain_name | domain name 
 domain_zone | domain zone 
 email_address | email address 
 email_body | email body 
 female_first_name | female first name 
 female_full_name | female full name 
 female_full_name_with_prefix | female full name with prefix 
 female_full_name_with_suffix | female full name with suffix 
 female_last_name | female last name 
 female_last_name_pratronymic | female last name pratronymic 
 first_name | first name 
 full_name | full name 
 full_name_with_prefix | full name with prefix 
 full_name_with_suffix | full name with suffix 
 gender | gender 
 gender_abrev | gender abrev 
 hex_color | hex color 
 hex_color_short | hex color short 
 i_pv_4 | i pv 4 
 industry | industry 
 job_title | job title 
 language | language 
 last_name | last name 
 latitude_degrees | latitude degrees 
 latitude_direction | latitude direction 
 latitude_minutes | latitude minutes 
 latitude_seconds | latitude seconds 
 latitude | latitude 
 longitude | longitude 
 longitude_degrees | longitude degrees 
 longitude_direction | longitude direction 
 longitude_minutes | longitude minutes 
 longitude_seconds | longitude seconds 
 male_first_name | male first name 
 male_full_name_with_prefix | male full name with prefix 
 male_full_name_with_suffix | male full name with suffix 
 male_last_name | male last name 
 male_pratronymic | male pratronymic 
 model | model 
 month | month 
 month_num | month num 
 month_short | month short 
 paragraph | paragraph 
 patagraphs | patagraphs 
 patagraphs_n | patagraphs n 
 password | password 
 patronymic | patronymic 
 phone | phone 
 product | product 
 product_name | product name 
 sentence | sentence 
 sentences | sentences 
 sentences_n | sentences n 
 simple_pass_word | simple pass word 
 state | state 
 state_abbrev | state abbrev 
 street | street 
 street_address | street address 
 title | title 
 top_level_domain | top level domain 
 user_name | user name 
 week_day | week day 
 week_day_short | week day short 
 week_day_num | week day num 
 word | word 
 words | words 
 words_n | words n 
 year | year 
 zip | zip 

# Author
Geofrey Ernest

Twitter  : [@gernesti](https://twitter.com/gernesti)


# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.

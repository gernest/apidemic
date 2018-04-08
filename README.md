# apidemic [![Build Status](https://travis-ci.org/gernest/apidemic.svg)](https://travis-ci.org/gernest/apidemic)

Apidemic is a service for generating fake JSON response. You first register the sample JSON response, and apidemic will serve that response with random fake data.

This is experimental, so take it with a grain of salt.

# Motivation
I got bored with hardcoding the sample json api response in tests. If you know golang, you can benefit by using the library, I have included a router that you can use to run disposable servers in your tests.

# Installation

You can download the binaries for your respective operating system  [Download apidemic](https://github.com/gernest/apidemic/releases/latest)

Then put the downloaded binary somewhere in your system path.

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


If you have already started apidemic server you can register that response by making a POST request to the `/register` path. Passing the json body of the form.

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
      "name:country": "Sweden"
    }
  }
}
```

See the annotation tags on the payload. Example if I want to generate full name for a field name I will just add `"name:full_name"`.

Once your POST request is submitted you are good to ask for the response with fake values. Just make a GET request to the endpoint you registered.

So every GET call to `/api/test` will return the api response with fake data.

# Routes 
Apidemic server has only three http routes

### /
This is the home path. It only renders information about the apidemic server.

### /register
This is where you register endpoints. You POST the annotated sample JSON here. The request body should be a json object of signature.

```json
{
	"endpoint":"my_endpoint",
	"payload": { ANNOTATED__SAMPLE_JSON_GOES_HERE },
}
```

#### /api/{REGISTERED_ENDPOINT_GOES_HERE}
Every GET request on this route will render a fake JSON object for the sample registered in this endpoint.

#### Other HTTP Methods

In case you need to mimic endpoints which respond to requests other than GET then make sure to add a `http_method` key with the required method name into your API description.

```json
{
  "endpoint": "test",
  "http_method": "POST",
  "payload": {
    "name: first_name": "anton"
  }
}
```

Currently supported HTTP methods are: `OPTIONS`, `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, default is `GET`. Please open an issue if you think there should be others added.

# Tags
Apidemic uses tags to annotate what kind of fake data to generate and also control different requrements of fake data.

You add tags to object keys. For instance let's say you have a JSON object `{ "user_name": "gernest"}`. If you want to have a fake username then you can annotate the key by adding user_name tag like this `{ "user_name:user_name": "gernest"}`.

So JSON keys can be annotated by adding the `:` symbol then followed by comma separated list of tags. The first entry after `:` is for the tag type, the following entries are in the form `key=value` which will be the extra information to fine-tune your fake data. Please see the example above to see how tags are used.

Apidemic comes shipped with a large number of tags, meaning it is capable to generate a wide range of fake information.

These are currently available tags to generate different fake data:

 Tag | Details( data generated)
------|--------
brand | brand 
 character | character 
 characters | characters 
 characters_n | characters of maximum length n
 city | city 
 color | color 
 company | company 
 continent | continent 
 country | country 
 credit_card_num | credit card number 
 currency | currency 
 currency_code | currency code 
 day | day 
 digits | digits 
 digits_n | digits of maximum number n
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
 patagraphs_n | patagraphs of maximum n
 password | password 
 patronymic | patronymic 
 phone | phone 
 product | product 
 product_name | product name 
 sentence | sentence 
 sentences | sentences 
 sentences_n | sentences of maximum n
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
 words_n | words of maximum n
 year | year 
 zip | zip 

# Benchmark
This Benchmark uses [boom](https://github.com/rakyll/boom). After registering the sample json above run the following command (Note this is just to check things out, my machine is very slow)

```bash
 boom -n 1000 -c 100 http://localhost:3000/api/test
```

The result
```bash

Summary:
  Total:	0.6442 secs.
  Slowest:	0.1451 secs.
  Fastest:	0.0163 secs.
  Average:	0.0586 secs.
  Requests/sec:	1552.3336
  Total Data Received:	39000 bytes.
  Response Size per Request:	39 bytes.

Status code distribution:
  [200]	1000 responses

Response time histogram:
  0.016 [1]	|
  0.029 [121]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.042 [166]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.055 [192]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.068 [192]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.081 [168]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.094 [69]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.106 [41]	|∎∎∎∎∎∎∎∎
  0.119 [22]	|∎∎∎∎
  0.132 [21]	|∎∎∎∎
  0.145 [7]	|∎

Latency distribution:
  10% in 0.0280 secs.
  25% in 0.0364 secs.
  50% in 0.0560 secs.
  75% in 0.0751 secs.
  90% in 0.0922 secs.
  95% in 0.1066 secs.
  99% in 0.1287 secs.
```

# Contributing

Start with clicking the star button to make the author and his neighbors happy. Then fork the repository and submit a pull request for whatever change you want to be added to this project.

If you have any questions, just open an issue.

# Author
Geofrey Ernest

Twitter  : [@gernesti](https://twitter.com/gernesti)


# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.

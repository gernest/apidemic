# apidemic [![Build Status](https://travis-ci.org/gernest/apidemic.svg)](https://travis-ci.org/gernest/apidemic)

Apidemic is a service for generating fake JSON API response. You first register the sample JSON response, and apidemic will serve that response with random fake data.

This is experimental, so take it with a grain of salt.

# Motivation
I got bored with nardcoding the sample json api response in tests. I you know golang, you can benefit by using the library, I have included a router that you can use to run disposable servers in your tests._

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

# Author
Geofrey Ernest

Twitter  : [@gernesti](https://twitter.com/gernesti)


# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.

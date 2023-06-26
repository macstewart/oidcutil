#### ssorry

A mock oidc server for my personal local oidc testing. 

- It has some use-case specific hardcoding so you'll probably need to fiddle with it
- I made it in an afternoon so pardon my architectural transgressions

##### Usage
```
git clone https://github.com/macstewart/oidcutil
go build
./ssorry server
```

There are some redundant endpoints for other use cases but mainly:
- runs on localhost:3333 by default. I think there's a cli argument to change this but I forget
- `/discovery` basic discovery endpoint that returns the authorize, token, and jwks endpoints
- `/authorize` accepts the redirect_uri and state and stuff, and pops up a web form to fill in fake user data (email + custom claim key + values). Changing these values persists until the web server is restarted
- `/token` returns the signed jwt with the basic IDToken that has the configured  details
- `/keys` returns the signing details for the jwt for key verification


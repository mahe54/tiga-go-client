# tiga-go-client
A client for creating and reading Tiga resources (roles more specifically).

Your machine have to be connected to Telia VPN.
Your shell that you are starting the visual studio code environment from needs to have proxy configured.
Here's an example:
````
export http_proxy=proxy-se-uan.ddc.teliasonera.net:8080
export HTTP_PROXY=proxy-se-uan.ddc.teliasonera.net:8080
export https_proxy=proxy-se-uan.ddc.teliasonera.net:8080
export HTTPS_PROXY=proxy-se-uan.ddc.teliasonera.net:8080
```

Without the proxy, you won't be able to reach SIDM to get your token needed for login towards Tiga.
Without VPN you won't be able to connect to Tiga

In your .bashrc, .zshrc you need to export:
```
export GOPRIVATE="github.com/telia-company/*"
````
Because you won't be able to 'go get' this module otherwise, to be included in other code bases.
also:
```
source data/prodEnvVars.sh
```
To get the environment variables you need.
It is important that you source/export variables from the shell from where you start your vs code env from.
Your vs code environment will inherit these settings.
No, it won't be sufficient to start a terminal inside vs code and set the variable from there because they don't affect the debugging. Those env vars can be set by other means as well, but the easiest is to set them in your launching shell.

Yes, you can in the vs code set env vars and run the main.go file from there and that will work, but you won't be able to debug (which is more fun, as everybody knows).
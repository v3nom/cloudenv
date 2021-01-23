# Cloud Env [![Build Status](https://travis-ci.com/v3nom/cloudenv.svg?branch=master)](https://travis-ci.com/v3nom/cloudenv)

Google App Engine standard Go environment no longer has built-in APIs for accesing internal services and it creates a lot of pain when establishing various client connections (Datastore, Logging, Cloudtasks etc). 

Cloudenv tackles this problem by writing all the boilerplate connection code for you.

# shopify2sqaure
- syncs products, orders, and customers from shopify to square. one way sync will only add (not update / remove)
- to be used when migrating from shopify to square

### Setup:
- install golang, install brew / dep
- create aws account / get iam accesss key
- install serverless / aws cli
- configure aws cli
- configure private app in shopify
- configure private app in ssquare
- set config file
- update makefile for appropriate aws-profile
- run make deploy 

### MakeFile Commands:
> make build

- builds the app and creates binaries in the 'bin' folder

> make test

- runs all unit tests

> make schema

- creates json schema files for serverless documentation and request validation

> make deploy

- performs serverless deployment

>make clean

- removes binaries, dependencies, and generated files

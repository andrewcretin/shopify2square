# shopify2sqaure
- syncs products, orders, and customers from shopify to square. one way sync will only add (not update / remove)
- to be used when migrating from shopify to square

### Setup:
Install Brew:
> https://brew.sh/

Use Brew to Install Dep:
> brew install dep

Use Dep to Install Packages:
> dep ensure

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


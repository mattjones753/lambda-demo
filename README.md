# lambda-demo

This repo contains some simple code for demonstrating lambdas.
They are very simple lambdas that are meant to be shown as part of an AWS API gateway setup. It assumes that you have set up a private s3 bucket with an API Gateway resource sitting in front of it that exposes the GET and PUT  methods.

## Lambdas
The lambda code can be found under the lamdba directory with each lambda being in its own subdirectory. 

### authoriser
Although the S3 bucket is not  publicly exposed, you've set up an API Gateway API with a resource that allows anyone to put objects in it!
This might not be the best idea - With API Gateway you can add custom authorizers so you can restrict access to sensitive endpoints.
Our example here is not meant to be a good example of security but just ensures that anyone making a request has an `Authorization` token with the value `secret`
When configured to use this lambda for the PUT method, that token will need to be in place to avoid a 401 reponse.

### file_validator
We might want to restrict what people put in our S3 bucket, even if they have the correct credentials. In this example, we only want them to upload documents containing valid JSON. When this lambda is configured to be a LAMBDA_PROXY for our PUT request in can be used to validate the file contents as being valid JSON before actually uploading the object to S3

## Building
running the following:
```bash
make
```

will create the lambda binaries in a `bin` directory at the root of the repo

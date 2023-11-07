# DEDIS Go imports

This repository contains the Go lambda hosted on AWS that redirects Go module
requests to the correct repository URL.

## How it works

The domain is redirected to an API Gateway hosted on AWS. This API is configured
so that any request is proxied to a lambda named `dedis-go-imports` also hosted
on AWS.

The Go lambda is using the code of [golang.org](https://github.com/golang/website/blob/master/cmd/golangorg/x.go)
as a base but modified to fit the DEDIS repositories.

## Supported repositories

The domain of the modules is `go.dedis.ch`. The list of repositories supported is in
`main.go`.

## Contribute

First you will need to install the AWS CLI client on your computer. Please use the
following [instructions](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html).
Then you'll need to setup the credentials:

```bash
$ aws configure
AWS Access Key ID [None]: # generate one on AWS
AWS Secret Access Key [None]: # secret of the generated key
Default region name [None]: eu-central-1
Default output format [None]: json
```

You can now deploy your changes using the provided script that will pack the
lambda code in a zip file and upload it to AWS.

```bash
./deploy.sh
```

Note that this will update the **production** lambda.

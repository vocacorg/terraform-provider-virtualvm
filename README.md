Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/vocacorg/terraform-provider-template`

```sh
$ mkdir -p $GOPATH/src/github.com/vocacorg; cd $GOPATH/src/github.com/vocacorg
$ git clone git@github.com:vocacorg/terraform-provider-template
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/vocacorg/terraform-provider-template
$ make build
```

Using the provider
----------------------

```hcl
# Configure the Bitbucket Provider
provider "vocacorg" {
  username = "GobBluthe"
  password = "idoillusions" # you can also use app passwords
}

# Manage your sample
resource "vocacorg_sample" "infrastructure" {
  owner = "myteam"
  name  = "terraform-code"
}
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-template
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Terraform needs TF_ACC env variable set to run acceptance tests

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

About V1 APIs
------------------

This provider will not take any PRs about the v1 apis that dont have v2
equivalents. Please only focus on v2 apis when adding new featues to this
provider.

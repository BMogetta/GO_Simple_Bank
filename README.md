# **goBank is the project exercise of the go backend course by** [techschoolgutu](https://www.youtube.com/c/TECHSCHOOLGURU)

This repository contains the source code and my notes about the course of back-end with go that I toke. It could be improve, it could be safer and I could extend the functionalities but I think that is beyond my original propose of familiarize with golang
It could run either gin server or grpc server. The gin server has more functionality, it wasn't the idea to repeat the code all over again.
___

**[Go](#go) + [Docker](#docker) + [PostGresSQL](postgres/README.md) + [AWS](eks/README.md) + [Kubernetes](eks/README.md) + [Github Actions](#gha) + [Paseto](token/README.md) + [gRPC](gapi/README.md)**

___

<a name="go"></a>
## **What is Go and how to [install](https://go.dev/doc/install) it**


Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. 
```
curl -Lo go.tgz https://go.dev/dl/go1.19.linux-amd64.tar.gz

sudo tar -C /usr/local -xvf go.tgz

go version
```
---
<a name="docker"></a> 
## **This project runs [Docker](https://docs.docker.com/) container**



Docker is a set of platform as a service products that use OS-level virtualization to deliver software in packages called containers. 

```
sudo apt-get update
 
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

---
## **This project uses a dockerize [PostGresSQL](/postgres/README.md)**

<br>

---

## **This project deploys to [AWS](eks/README.md)**

Amazon Web Services (AWS) is the world’s most comprehensive and broadly adopted cloud platform.

---
## **Install [make](https://www.gnu.org/software/make/manual/make.html)**
Make is a build automation tool that automatically builds executable programs and libraries from source code by reading files called Makefiles which specify how to derive the target program. 

Install with:
```
sudo apt install make

make --version
```

---

## **Run the project**
<br>

**Docker-compose:**

```docker-compose up
make test
```

**Manually:**

```
make network
make postgres
make createdb
make migrateup
make server
make test
```
---
## **[Gin Web Framework:](https://github.com/gin-gonic/gin)**

Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to [httprouter](https://github.com/julienschmidt/httprouter).

Install with:

`go get -u github.com/gin-gonic/gin`    

---
## **Handle environment variables with [Viper](https://github.com/spf13/viper):**

Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

Install with:

`go get github.com/spf13/viper`

---
## **Mock-db with [gomock](https://github.com/golang/mock):**

gomock is a mocking framework for the Go programming language. It integrates well with Go's built-in testing package, but can be used in other contexts too.

Install with:

`go install github.com/golang/mock/mockgen@v1.6.0`


---
## **Usefull extensions**


Install [AWS Toolkit](https://marketplace.visualstudio.com/items?itemName=AmazonWebServices.aws-toolkit-vscode) to interact with AWS in VSCode

Install [vscode-dbml](https://marketplace.visualstudio.com/items?itemName=matt-meyers.vscode-dbml) for .DBML language support

Install [proto 3](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3) for protobuf 3 language suppport

In VSCode `Ctrl+p`

`ext install matt-meyers.vscode-dbml`

`ext install zxh404.vscode-proto3`

`ext install AmazonWebServices.aws-toolkit-vscode`

---
<a name="gha"></a> 
# **Automate your workflow from idea to production**


GitHub Actions makes it easy to automate all your software workflows, now with world-class CI/CD. Build, test, and deploy your code right from GitHub. Make code reviews, branch management, and issue triaging work the way you want.

---

## [Search the marketplace](https://github.com/marketplace)

Find apps to use across your development process, from continuous integration to project management and code review. Then start using them without setting up multiple accounts or payment methods.

---

## [About continuous integration](https://docs.github.com/en/actions/automating-builds-and-tests/about-continuous-integration)

You can create custom continuous integration (CI) workflows directly in your GitHub repository with GitHub Actions.

---

## [About monitoring and troubleshooting](https://docs.github.com/en/actions/monitoring-and-troubleshooting-workflows/about-monitoring-and-troubleshooting)

You can use the tools in GitHub Actions to monitor and debug your workflows.

---

## [Deploying with GitHub Actions](https://docs.github.com/en/actions/deployment/about-deployments/deploying-with-github-actions)

Learn how to control deployments with features like environments and concurrency.

---
---
### ***Final clean up***

Once I finish the course I decided to cleanup the commit history -it had a lot of unnecessary test and tries- and when I eventually revisit this project to recap some of the implementations I would like to see a smooth gh page.

The steps that I use:
 First delete all branches -and tags if you have some- besides master. Then unprotect the master branch

```
git checkout --orphan finish
git add -A
git commit -S -m "finish"
git branch -D master
git branch -m master
git push -f origin master
git gc --aggressive --prune=all
```
# Gouth

## Description
An API for a restaurant that operates with Square POS.

## Links

[![Run in Postman](https://run.pstmn.io/button.svg)](https://documenter.getpostman.com/view/31799606/2sA2rCTgtA)

# Objectives

1. Basic DevOps practices
2. Security considerations
3. DevSecOps practices
4. Cloud capabilities

# Details

- Microservice: Auth Service
- Language: Go Lang
- Repo: https://github.com/mrmihi/gouth
- Cloud: GCloud
- Compute Service: Cloud Run
- Build Service: Cloud Build
- Container Registry: Artifact Registry
- Containerization: Docker

## Prerequisites

- [Air](https://github.com/cosmtrek/air) - For live reloading
- [Node](https://nodejs.org/en/) - For setting up git hooks (optional)

## Getting started

- Run `go mod download` to download all dependencies
- Run `go run ./src/main.go` to start the development server without live reloading
- Run `air` to start the development server with hot reloading

- Optionally, run `pnpm install` to set up git hooks
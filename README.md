# go-crud

Golang package for abstractions of basic GCP CRUD operations, API Handlers, JWT, as well as other useful functions.

This is a copy of what was used by [Canary](https://github.com/maxcraig112/FIT3162-Canary), however is written so that it can be more easily used as a Golang package for other projects.

# Repository Structure

| Directory | Description                                                                                                            |
| --------- | ---------------------------------------------------------------------------------------------------------------------- |
| gcp       | Contains basic CRUD functions for Google Storage Buckets, Firestore and Google Secret Manager                          |
| handler   | Defines a handler meant to abstract required context, GCP clients and AuthMW for an API handler                        |
| jwt       | Defines common functions needed to create a JWT HTTP Middleware, validate a JWT, and evaluate a userID from JWT claims |
| password  | Common functions for hashing and comparing bcrypt password hashes                                                      |

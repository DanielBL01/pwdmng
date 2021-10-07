# CLI Password Manager

## Table of Contents

- [About this project](#about-this-project)
    - [Purpose](#purpose)
    - [Built With](#built-with)
    - [Future Improvements](#future-improvements)
- [Getting Started](#getting-started)
    - [Prerequisite](#prerequisite)
    - [Installation](#installation)    
- [Usage](#usage)

## About this project
### Purpose 
For a long time, I would store passwords with the sticky notes app and sometimes even stored them physically on pieces of paper!

This CLI Password Manager is a secure password manager which runs locally on your command line. The CLI Password Manager requires 
an initial login with a password and will automatically log out once the temrinal session is closed.

The single password neccessary to log into your password manager goes through an encryption and stored. 

### Built With
- [Go](https://golang.org)
- [Redis](https://redis.io)
- [Docker](https://www.docker.com)

### Future Improvements
- Use a cli framework for better UI
- Use Redis as a fast key value database
- Use Docker to containerize project

## Getting Started

### Prerequisite

### Installation

## Usage
#### Login 
```
$ pwdmng login
```
#### Add password
```
$ pwdmng add <website> <username/email> <password>
```
#### Get password
```
$ pwdmng get <website>
```
#### Update password
```
$ pwdmng update <website>
```
#### Remove password
```
$ pwdmng remove <website>
```
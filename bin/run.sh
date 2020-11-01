#!/bin/bash

run -d -p 3001:3001 --env-file blochain.env --name golang-blogchain-server qwx1337/blogchain-server:latest
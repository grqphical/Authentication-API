# Authentication API

[![API Tests](https://github.com/grqphical07/Authentication-API/actions/workflows/tests.yaml/badge.svg)](https://github.com/grqphical07/Authentication-API/actions/workflows/tests.yaml)

A simple Authentication REST API made with Go, SQLite and Gin.

It works by having a Refresh JSON Web Token and an Access JSON Web Token. The access token allows you access to autheticated parts of the site whilst the refresh token allows
you to get a new access token if the one you have expires

The backend just uses an SQLite database to store account credentials

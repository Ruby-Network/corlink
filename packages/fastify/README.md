# Corlink Fastify Plugin

This plugin allows you to get started using Corlink with your existing Fastify application easily.

## Usage

- ECMAScript Module (ESM)
```javascript
import fastify from 'fastify';
import path from 'path';
import dotenv from 'dotenv';
dotenv.config();
import corlink from '@rubynetwork/corlink-fastify';

const app = fastify({ logger: true });
app.register(corlink, {
    //the file to serve when a user is denied access
    deniedFilePath: path.join(__dirname, 'denied.html').toString(),
    //any file or route you want to be accessible without a valid Corlink token (e.g. Bare servers)
    unlockedPaths: ['/bare/'],
    //any domains you want corlink to ignore
    whiteListedURLs: [""],
    //the URL of a Colink API server
    corlinkUrl: process.env.CORLINK_URL,
    //the API key for the Corlink API server
    corlinkAPIKey: process.env.CORLINK_API_KEY,
    //use the built in cookie parser (set to false if you are using @fastify/cookie before this plugin)
    builtinCookieParser: true,
});

app.get('/', async (req, res) => {
    return { hello: 'world' };
});

app.listen({ port: 3000 }, (err, address) => {
    if (err) {
        app.log.error(err);
        process.exit(1);
    }
    app.log.info(`server listening on ${address}`);
});
```


- CommonJS Module (CJS)
```javascript
const fastify = require('fastify');
const path = require('path');
const dotenv = require('dotenv');
dotenv.config();
const corlink = require('@rubynetwork/corlink-fastify');

const app = fastify({ logger: true });

app.register(corlink, {
    //the file to serve when a user is denied access
    deniedFilePath: path.join(__dirname, 'denied.html').toString(),
    //any file or route you want to be accessible without a valid Corlink token (e.g. Bare servers)
    unlockedPaths: ['/bare/'],
    //any domains you want corlink to ignore
    whiteListedURLs: [""],
    //the URL of a Colink API server
    corlinkUrl: process.env.CORLINK_URL,
    //the API key for the Corlink API server
    corlinkAPIKey: process.env.CORLINK_API_KEY,
    //use the built in cookie parser (set to false if you are using @fastify/cookie before this plugin)
    builtinCookieParser: true,
});

app.get('/', async (req, res) => {
    return { hello: 'world' };
});

app.listen({ port: 3000 }, (err, address) => {
    if (err) {
        app.log.error(err);
        process.exit(1);
    }
    app.log.info(`server listening on ${address}`);
});
```

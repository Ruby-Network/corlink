# Corlink Express Middleware

This middleware allows you to get started with using Corlink in your existing Express applications easily and simply.

## Usage
> [!IMPORTANT]
> This middleware requires the `cookie-parser` middleware to be used before it.

- ECMAScript 

```javascript
import express from 'express';
import path from 'path';
import cookieParser from 'cookie-parser';
import { crypto } from 'crypto';
import dotenv from 'dotenv';
dotenv.config();
import { corlinkExpress } from '@rubynetwork/corlink-express';

const cookieSecret = crypto.randomBytes(64).toString('hex');
//or:
// const cookieSecret = process.env.COOKIE_SECRET;
const app = express();
app.use(cookieParser(cookieSecret));
app.use(corlinkExpress({ 
    //the page in which the user will be redirected to if the user is not authorized
    deniedFilePath: path.join(__dirname, 'rejected.html'),
    //any endpoints or path's you don't want to be protected by corlink 
    unlockedPaths: ['/bare/'],
    //any urls you don't want to be protected by corlink 
    whiteListedURLs: [],
    //corlink API endpoint
    corlinkUrl: process.env.CORLINK_API_ENDPOINT,
    //corlink API key
    corlinkAPIKey: process.env.CORLINK_API_KEY
}))

app.get('/', (req, res) => {
    res.send('Hello World');
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
```

- CommonJS

```javascript
const express = require('express');
const path = require('path');
const cookieParser = require('cookie-parser');
const crypto = require('crypto');
const dotenv = require('dotenv');
dotenv.config();
const { corlinkExpress } = require('@rubynetwork/corlink-express');

const cookieSecret = crypto.randomBytes(64).toString('hex');
//or:
// const cookieSecret = process.env.COOKIE_SECRET;
const app = express();
app.use(cookieParser(cookieSecret));
app.use(corlinkExpress({
    //the page in which the user will be redirected to if the user is not authorized
    deniedFilePath: path.join(__dirname, 'rejected.html'),
    //any endpoints or path's you don't want to be protected by corlink 
    unlockedPaths: ['/bare/'],
    //any urls you don't want to be protected by corlink 
    whiteListedURLs: [],
    //corlink API endpoint
    corlinkUrl: process.env.CORLINK_API_ENDPOINT,
    //corlink API key
    corlinkAPIKey: process.env.CORLINK_API_KEY
}))

app.get('/', (req, res) => {
    res.send('Hello World');
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
```

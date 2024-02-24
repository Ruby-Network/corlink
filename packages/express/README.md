# Corlink Express Middleware

This middleware allows you to get started with using Corlink in your existing Express applications easily and simply.

## Usage

- ECMAScript 

```javascript
import express from 'express';
import path from 'path';
import cookieParser from 'cookie-parser';
import { crypto } from 'crypto';
import dotenv from 'dotenv';
dotenv.config();
import { corlinkExpress, corlink } from '@rubynetwork/corlink-express';

const corlinkInstance = new corlink(
    // The Page in which the user will be redirected to if the user is not authorized
    path.join(__dirname, 'rejected.html'),
    // Bare server endpoint
    '/bare/',
    // Corlink API server endpoint
    process.env.CORLINK_API_ENDPOINT,
    // Corlink API key
    process.env.CORLINK_API_KEY
);
const cookieSecret = crypto.randomBytes(64).toString('hex');
//or:
// const cookieSecret = process.env.COOKIE_SECRET;
const app = express();
app.use(cookieParser(cookieSecret));
app.use(corlinkExpress({ corlinkInstance }));

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
const { corlinkExpress, corlink } = require('@rubynetwork/corlink-express');

const corlinkInstance = new corlink(
    // The Page in which the user will be redirected to if the user is not authorized
    path.join(__dirname, 'rejected.html'),
    // Bare server endpoint
    '/bare/',
    // Corlink API server endpoint
    process.env.CORLINK_API_ENDPOINT,
    // Corlink API key
    process.env.CORLINK_API_KEY
);

const cookieSecret = crypto.randomBytes(64).toString('hex');
//or:
// const cookieSecret = process.env.COOKIE_SECRET;
const app = express();
app.use(cookieParser(cookieSecret));
app.use(corlinkExpress({ corlinkInstance }));

app.get('/', (req, res) => {
    res.send('Hello World');
});

app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
```

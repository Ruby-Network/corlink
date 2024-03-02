# Corlink SDK Node

This is the official Corlink SDK for Node.js. It provides a simple way to interact with the Corlink API.

## Installation

```bash
npm install @rubynetwork/corlink-sdk-node
```

## Usage

- CommonJS
```javascript
const { Corlink } = require('@rubynetwork/corlink-sdk-node');
//create a new Corlink instance
const corlink = new Corlink({
    corlinkAPIUrl: 'https://corlink.example.com',
    corlinkAPIKey: 'your-api-key',
});

//use any of the methods here ex: createKey
async function createKey() {
    const key = await corlink.createKey();
    console.log(key);
}

createKey();
```

- EcmaScript Modules (ESM)
```javascript
import { Corlink } from '@rubynetwork/corlink-sdk-node';

//create a new Corlink instance
const corlink = new Corlink({
    corlinkAPIUrl: 'https://corlink.example.com',
    corlinkAPIKey: 'your-api-key',
});

//use any of the methods here ex: createKey
async function createKey() {
    const key = await corlink.createKey();
    console.log(key);
}

createKey();
```

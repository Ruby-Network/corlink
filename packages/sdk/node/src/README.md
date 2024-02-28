# Corlink Node SDK

This is the official Corlink Node SDK. It provides a convenient way to access the Corlink API from applications written in server-side JavaScript.

## Usage

The package needs to be installed using npm (or other package managers):

```bash
npm install @rubynetwork/corlink-sdk-node
```

The package can then be included in the application:

- ECMAScript module:

```javascript
import { Corlink } from '@rubynetwork/corlink-sdk-node';

//create a new instance of the Corlink class
const corlink = new Corlink({
    corlinkAPIUrl: process.env.CORLINK_API_URL,
    corlinkAPIKey: process.env.CORLINK_API_KEY
});

//use the instance to call the Corlink API methods
```

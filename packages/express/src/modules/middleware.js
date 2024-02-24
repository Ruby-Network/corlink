import { readFileSync } from 'fs';

/**
    * @function corlink
    * @param {string} deniedFilePath - The path to and html file that will be served to the client if the request is denied 
    * @param {string} bareServerPath - The path to the server that will be used to serve the file 
    * @param {string} corlinkUrl - A URL pointing to a corlink API server
    * @param {string} corlinkAPIKey - A key used to authenticate with the corlink API server
    * @description This function is used to create a new instance of the corlink middleware
**/
function corlink(deniedFilePath, bareServerPath, corlinkUrl, corlinkAPIKey) {
    this.deniedFilePath = deniedFilePath;
    this.bareServerPath = bareServerPath;
    this.corlinkUrl = corlinkUrl;
    this.corlinkAPIKey = corlinkAPIKey;
    return { deniedFilePath, bareServerPath, corlinkUrl, corlinkAPIKey }
}

async function externallyValidateCookies(cookies) {
    try {
        if (cookies === undefined || cookies === null) {
            throw new Error('The cookies are not valid');
        }
        if (cookies.userIsVerified === undefined || cookies.userIsVerified === null) {
            throw new Error('The cookies are not valid');
        }
        if (cookies.userIsVerified === false) {
            throw new Error('The user is not verified');
        }
    }
    catch (e) {
        throw new Error('The cookies are not valid');
    }
}

async function cookieValidator(cookies, fileContent, res) {
    try {
        await externallyValidateCookies(cookies.userIsVerified);
    }
    catch (e) {
        //set the content type to html
        fail(res, fileContent);
        throw new Error('The cookies are not valid');
    }
}

function fail(res, fileContent) {
    res.setHeader('Content-Type', 'text/html');
    res.send(fileContent);
}

async function verifyUser(key, corlinkUrl, corlinkAPIKey) {
    if (corlinkUrl[corlinkUrl.length - 1] !== '/') {
        corlinkUrl += '/';
    }
    try {
        const t = await fetch(corlinkUrl + 'verify', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${corlinkAPIKey}`,
                'Key': key 
            },
        });
        const tt = await t.json();
        if (tt.status !== "ok") {
            throw new Error('The user could not be verified');
        }
        else {
            if (tt.message !== "Authorized") {
                throw new Error('The user could not be verified');
            }
            return;
        }
    }
    catch (e) {
        throw new Error('The user could not be verified');
    }
}

/**
    * @function middleware
    * @param {corlinkInstance} corlinkInstance - An instance of the corlink middlewarw
**/
function middleware(corlinkInstance) {
    corlinkInstance = corlinkInstance || {};
    try {
        const t = corlinkInstance.corlinkInstance.deniedFilePath;
        const p = corlinkInstance.corlinkInstance.bareServerPath;
        const t2 = corlinkInstance.corlinkInstance.corlinkUrl;
        const p2 = corlinkInstance.corlinkInstance.corlinkAPIKey;
        if (t === undefined || p === undefined || t === '' || p === '' || t === null || p === null, t2 === undefined || p2 === undefined || t2 === '' || p2 === '' || t2 === null || p2 === null) {
            throw new Error('The instance is not valid');
        }
    }
    catch (e) {
        throw new Error('The instance is not valid');
    }
    try {
        readFileSync(corlinkInstance.corlinkInstance.deniedFilePath, 'utf8');
    }
    catch (e) {
        throw new Error('The file could not be read');
    }
    return async function (req, res, next) {
        const file = readFileSync(corlinkInstance.corlinkInstance.deniedFilePath, 'utf8');
        const authHeader = req.headers.authorization; 
        const corlinkUrl = corlinkInstance.corlinkInstance.corlinkUrl;
        const corlinkAPIKey = corlinkInstance.corlinkInstance.corlinkAPIKey;
        if (req.signedCookies.userIsVerified) {
            next();
            return;
        }
        //credit: https://github.com/titaniumnetwork-dev/MasqrProject/blob/master/MasqrBackend/index.js#L75 for this trick too
        if (req.cookies.refreshcheck != "true") {
            res.cookie('refreshcheck', 'true', { maxAge: 10000 });
            fail(res, file);
            return;
        }
        if (!authHeader) {
            res.setHeader('WWW-Authenticate', 'Basic');
            res.status(401);
            fail(res, file);
            return;
        }
        const auth = new Buffer.from(authHeader.split(' ')[1], 'base64').toString().split(':');
        const user = auth[0];
        const pass = auth[1];
        try {
            await verifyUser(pass, corlinkUrl, corlinkAPIKey);
        }
        catch (e) {
            res.status(401);
            fail(res, file);
            return;
        }
        // 1 year
        const maxCookieAge = 60 * 60 * 24 * 365;
        //set the cookie to an encrypted version of the pass 
        res.cookie('userIsVerified', pass, { signed: true, maxAge: maxCookieAge, sameSite: 'strict', secure: true });
        //credit: https://github.com/titaniumnetwork-dev/MasqrProject/blob/master/MasqrBackend/index.js#L98 for this trick
        res.send(`<script> window.location.href = window.location.href </script>`);
        return;
    }
}

export { middleware, corlink };

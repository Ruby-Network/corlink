import { readFileSync } from 'fs';
import { Request, Response, NextFunction } from 'express';
import { z } from 'zod';

async function externallyValidateCookies(cookies: any) {
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

async function cookieValidator(cookies: any, fileContent: any, res: Response) {
    try {
        await externallyValidateCookies(cookies.userIsVerified);
    }
    catch (e) {
        //set the content type to html
        fail(res, fileContent);
        throw new Error('The cookies are not valid');
    }
}

function fail(res: Response, fileContent: any) {
    res.setHeader('Content-Type', 'text/html');
    res.send(fileContent);
}

async function verifyUser(key: string, masqrUrl: string, host?: string) {
    try {
        const t = await fetch(masqrUrl + key + `&host=${host}`);
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

function validate(options: any) {
    const schema = z.object({
        deniedFilePath: z.string(),
        unlockedPaths: z.array(z.string()),
        whiteListedURLs: z.array(z.string()),
        masqrUrl: z.string(),
    }).safeParse(options);
    if (!schema.success) {
        if (schema.error.format().masqrUrl) {
            throw new Error('The option corlinkUrl is not a string: ' + schema.error.format().masqrUrl?._errors);
        }
        if (schema.error.format().deniedFilePath) {
            throw new Error('The option deniedFilePath is not a string: ' + schema.error.format().deniedFilePath?._errors);
        }
        if (schema.error.format().unlockedPaths) {
            throw new Error('The option unlockedPaths is not an array: ' + schema.error.format().unlockedPaths?._errors);
        }
        if (schema.error.format().whiteListedURLs) {
            throw new Error('The option whiteListedURLs is not an array: ' + schema.error.format().whiteListedURLs?._errors);
        }
    }
}


/**
    * @function middleware
    * @param {Object} options - The options for the middleware
    * @param {string} options.deniedFilePath - The path to the file that will be sent to the user if they are not verified 
    * @param {string[]} options.unlockedPaths - The path to the server 
    * @param {string[]} options.whiteListedURLs - The white listed URLs 
    * @param {string} options.corlinkUrl - The corlink API URL 
    * @param {string} options.corlinkAPIKey - The corlink API key 
    * @returns {Function} - The middleware function 
    * @description - This middleware function will verify the user using the corlink API 
**/
function masqr(options: any) {
    validate(options);
    try {
        readFileSync(options.deniedFilePath, 'utf8');
    }
    catch (e) {
        throw new Error('The file at the path ' + options.deniedFilePath + ' could not be read');
    }
    return async function (req: Request, res: Response, next: NextFunction) {
        const file = readFileSync(options.deniedFilePath, 'utf8');
        const authHeader = req.headers.authorization; 
        const masqrUrl = options.masqrUrl;
        if (options.whiteListedURLs.includes(req.headers.host)) {
            next();
            return;
        }
        if (options.unlockedPaths.includes(req.path)) {
            next();
            return;
        }
        if (req.signedCookies.userIsVerified) {
            next();
            return;
        } 
        if (req.cookies.refreshcheck != "true") {
            res.cookie('refreshcheck', 'true', { sameSite: 'strict', secure: true });
            fail(res, file);
            return;
        }
        if (!authHeader) {
            res.setHeader('WWW-Authenticate', 'Basic');
            res.status(401);
            fail(res, file);
            return;
        }
        //@ts-expect-error buffer is not defined
        const auth = new Buffer.from(authHeader.split(' ')[1], 'base64').toString().split(':');
        const user = auth[0];
        const pass = auth[1];
        try {
            await verifyUser(pass, masqrUrl, req.headers.host)
        }
        catch (e) {
            console.log('User not verified');
            res.status(401);
            fail(res, file);
            return;
        }
        const maxCookieAge = 60 * 60 * 24 * 365;
        res.cookie('userIsVerified', pass, { signed: true, maxAge: maxCookieAge, sameSite: 'strict', secure: true });
        //credit: https://github.com/titaniumnetwork-dev/MasqrProject/blob/master/MasqrBackend/index.js#L98 for this trick
        //console.log('User verified');
        res.send(`<script> window.location.href = window.location.href </script>`);
        return;
    }
}

export { masqr };

import fp from 'fastify-plugin';
import { z } from 'zod';
import { readFileSync } from 'fs';
import fastifyCookie from '@fastify/cookie';

function verify(opts) {
    const schema = z.object({
        deniedFilePath: z.string(),
        unlockedPaths: z.array(z.string()),
        whiteListedURLs: z.array(z.string()),
        corlinkUrl: z.string(),
        corlinkAPIKey: z.string()
    }).safeParse(opts);
    if (!schema.success) {
        if (schema.error.format().corlinkUrl) {
            throw new Error('The option corlinkUrl is not a string: ' + schema.error.format().corlinkUrl?._errors);
        }
        if (schema.error.format().corlinkAPIKey) {
            throw new Error('The option corlinkAPIKey is not a string: ' + schema.error.format().corlinkAPIKey?._errors);
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

function fail(reply, fileContent) {
    reply.header('Content-Type', 'text/html');
    reply.send(fileContent);
}

function turnToHostname(url) {
    try {
        return new URL(url).hostname;
    }
    catch (error) {
        return url;
    }
}

async function verifyUser(pass, corlinkUrl, corlinkAPIKey) {
    if (corlinkUrl[corlinkUrl.length - 1] !== '/') {
        corlinkUrl += '/';
    }
    try {
        const t = await fetch(corlinkUrl + 'verify', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${corlinkAPIKey}`,
                'Key': pass
            },
        });
        const tt = await t.json();
        if (!tt.status != "ok") {
            throw new Error('The user is not verified');
        }
        else {
            if (tt.message != "Authorized") {
                throw new Error('The user is not verified');
            }
            return true;
        }
    }
    catch (error) {
        return false;
    }
}


const plugin = (fastify, opts, done) => {
    verify(opts);
    try {
        readFileSync(opts.deniedFilePath);
    } catch (error) {
        throw new Error('The deniedFilePath does not exist');
    }
    fastify.register(fastifyCookie, {
        secret: opts.corlinkAPIKey, 
        parseOptions: {}
    });
    fastify.addHook('onRequest', function (req, reply, next) {
        const file = readFileSync(opts.deniedFilePath, 'utf8');
        const authHeader = req.headers.authorization;
        const corlinkUrl = opts.corlinkUrl;
        const corlinkAPIKey = opts.corlinkAPIKey;
        const whiteListedURLs = opts.whiteListedURLs.map(turnToHostname);
        if (whiteListedURLs.includes(req.hostname)) {
            next();
            return;
        }
        if (opts.unlockedPaths.includes(req.url)) {
            next();
            return;
        }
        //get a userIfVerified cookie
        if (req.cookies.userIfVerified) {
            next();
            return;
        }
        if (req.cookies.refreshcheck != 'true') {
            reply.setCookie('refreshcheck', 'true', { path: '/', sameSite: 'strict', secure: true, maxAge: 10000 }).type('text/html');
            fail(reply, file);
            return;
        }
        if (!authHeader) {
            reply.code(401).header('WWW-Authenticate', 'Basic');
            fail(reply, file);
            return;
        }
        const auth = new Buffer.from(authHeader.split(' ')[1], 'base64').toString().split(':');
        const user = auth[0];
        const pass = auth[1];
        const isVerified = verifyUser(pass, corlinkUrl, corlinkAPIKey);
        if (!isVerified) {
            reply.status(401).header('WWW-Authenticate', 'Basic').type('text/html');
            fail(reply, file);
            return;
        }
        else {
            reply.setCookie('userIfVerified', 'true', { path: '/', sameSite: 'strict', secure: true, maxAge: 10000 }).type('text/html').send('<script>window.location.href = window.location.href</script>');
            return;
        }
    });
    done();
};

/**
    * @typedef {Object} CorlinkOptions
    * @property {string} deniedFilePath - The path to the file that contains the denied URLs
    * @property {string[]} unlockedPaths - The paths that are not going to be checked by corlink 
    * @property {string[]} whiteListedURLs - The URLs that are not going to be checked by corlink 
    * @property {string} corlinkUrl - The URL of the corlink API 
    * @property {string} corlinkAPIKey - The API key of the corlink API 
**/
const corlink = fp(plugin, {
    fastify: '4.x',
    name: '@rubynetwork/corlink-fastify',
});
export default corlink;

import { schema } from './sdk.js';
import { z } from 'zod';

function error(message: string) {
    console.error(message);
    return;
}

async function createkey(opts: z.infer<typeof schema>) {
    if (opts.corlinkAPIUrl.slice(-1) !== '/') {
        opts.corlinkAPIUrl += '/';
    }
    try {
        const resp = await fetch(opts.corlinkAPIUrl + 'generate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + opts.corlinkAPIKey,
            },
        });
        const data = await resp.json();
        if (data.status !== "ok") {
            error('Error: ' + data.message);
            return;
        }
        return data;
    }
    catch (e) {
        error('Error: ' + e);
        return;
    }
}

async function deletekey(opts: z.infer<typeof schema>, apiKey: string) {
    if (opts.corlinkAPIUrl.slice(-1) !== '/') {
        opts.corlinkAPIUrl += '/';
    }
    try {
        const resp = await fetch(opts.corlinkAPIUrl + 'delete', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + opts.corlinkAPIKey,
                'Key': apiKey,
            },
        });
        const data = await resp.json();
        if (data.status !== "ok") {
            error('Error: ' + data.message);
            return;
        }
        if (data.message !== "Deleted") {
            error('Error: ' + data.message);
            return;
        }
        return data;
    }
    catch (e) {
        error('Error: ' + e);
        return;
    }
}

async function verifykey(opts: z.infer<typeof schema>, key: string) {
    if (opts.corlinkAPIUrl.slice(-1) !== '/') {
        opts.corlinkAPIUrl += '/';
    }
    try {
        const resp = await fetch(opts.corlinkAPIUrl + 'verify', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + opts.corlinkAPIKey,
                'Key': key,
            },
        });
        const data = await resp.json();
        if (data.status !== "ok") {
            error('Error: ' + data.message);
            return;
        }
        if (data.message !== "Authorized") {
            error('Error: ' + data.message);
            return;
        }
        return data;
    }
    catch (e) {
        error('Error: ' + e);
        return;
    }
}

export { createkey, deletekey, verifykey };

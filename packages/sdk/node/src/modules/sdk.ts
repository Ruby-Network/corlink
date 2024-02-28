import { z } from 'zod';
import { createuser } from './module.js'

const schema = z.object({
    corlinkAPIUrl: z.string(),
    corlinkAPIKey: z.string(),
});

function verify(opts: z.infer<typeof schema>) {
    const verified = schema.safeParse(opts);
    if (!verified.success) {
        if (verified.error.format().corlinkAPIUrl) {
            throw new Error('Error: corlinkAPIUrl is invalid: ' + verified.error.format().corlinkAPIUrl?._errors);
        }
        if (verified.error.format().corlinkAPIKey) {
            throw new Error('Error: corlinkAPIKey is invalid: ' + verified.error.format().corlinkAPIKey?._errors);
        }
        else {
            throw new Error('Error: ' + verified.error.format()?._errors);
        }
    }
}

class Corlink {
    opts: z.infer<typeof schema>;
    constructor(opts: z.infer<typeof schema>) {
        verify(opts);
        this.opts = opts;
    }
    createUser(user: string) {
        createuser(this.opts, user);
    }
}

export { Corlink, verify, schema };

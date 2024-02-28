import { verify, schema } from './sdk.js';
import { z } from 'zod';

/**
    * @function createUser
    * @description - A function to create a user using the Corlink API
    * @param {Object} opts - Options from the corlink function 
    * @param {string} user - The user to create 
    * @returns {void}
    * @throws {Error} - If the options are invalid
    * @example 
    * const opts = new corlink({ corlinkAPIUrl: 'https://corlinkapi.com', corlinkAPIKey: '12345' });
    * createUser(opts, 'user');
**/
function createuser(opts: z.infer<typeof schema>, user: string) {
    console.log('Creating user: ' + user);
}

export { createuser };

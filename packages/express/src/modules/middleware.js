import { readFileSync } from 'fs';

/**
    * @function corlink
    * @param {string} deniedFilePath - The path to and html file that will be served to the client if the request is denied 
    * @param {string} bareServerPath - The path to the server that will be used to serve the file 
    * @description This function is used to create a new instance of the corlink middleware
**/
function corlink(deniedFilePath, bareServerPath) {
    this.deniedFilePath = deniedFilePath;
    this.bareServerPath = bareServerPath;
    return { deniedFilePath, bareServerPath };
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
        if (t === undefined || p === undefined) {
            throw new Error('The instance is not valid');
        }
    }
    catch (e) {
        throw new Error('The instance is not valid');
    }
    return function (req, res, next) {
        console.log('middleware');
        next();
    }
}

export { middleware, corlink };

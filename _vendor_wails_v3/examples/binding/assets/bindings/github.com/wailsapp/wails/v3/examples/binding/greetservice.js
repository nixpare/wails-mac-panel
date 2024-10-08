// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * GreetService is a service that greets people
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "/wails/runtime.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as data$0 from "./data/models.js";

/**
 * GetPerson returns a person with the given name.
 * @param {string} name
 * @returns {Promise<data$0.Person> & { cancel(): void }}
 */
export function GetPerson(name) {
    let $resultPromise = /** @type {any} */($Call.ByID(2952413357, name));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        return $$createType0($result);
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

/**
 * Greet greets a person
 * @param {string} name
 * @param {number[]} counts
 * @returns {Promise<string> & { cancel(): void }}
 */
export function Greet(name, ...counts) {
    let $resultPromise = /** @type {any} */($Call.ByID(1411160069, name, counts));
    return $resultPromise;
}

/**
 * GreetPerson greets a person
 * @param {data$0.Person} person
 * @returns {Promise<string> & { cancel(): void }}
 */
export function GreetPerson(person) {
    let $resultPromise = /** @type {any} */($Call.ByID(4021313248, person));
    return $resultPromise;
}

// Private type creation functions
const $$createType0 = data$0.Person.createFrom;

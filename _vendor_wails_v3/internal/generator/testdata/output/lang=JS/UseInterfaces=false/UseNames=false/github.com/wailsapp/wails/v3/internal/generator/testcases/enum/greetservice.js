// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * GreetService is great
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "/wails/runtime.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

/**
 * Greet does XYZ
 * @param {string} name
 * @param {$models.Title} title
 * @returns {Promise<string> & { cancel(): void }}
 */
export function Greet(name, title) {
    let $resultPromise = /** @type {any} */($Call.ByID(1411160069, name, title));
    return $resultPromise;
}

/**
 * NewPerson creates a new person
 * @param {string} name
 * @returns {Promise<$models.Person | null> & { cancel(): void }}
 */
export function NewPerson(name) {
    let $resultPromise = /** @type {any} */($Call.ByID(1661412647, name));
    let $typingPromise = /** @type {any} */($resultPromise.then(($result) => {
        return $$createType1($result);
    }));
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

// Private type creation functions
const $$createType0 = $models.Person.createFrom;
const $$createType1 = $Create.Nullable($$createType0);

// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Create as $Create} from "/wails/runtime.js";

/**
 * Person represents a person
 */
export class Person {
    /**
     * Creates a new Person instance.
     * @param {Partial<Person>} [$$source = {}] - The source object to create the Person.
     */
    constructor($$source = {}) {
        if (!("Title" in $$source)) {
            /**
             * @member
             * @type {Title}
             */
            this["Title"] = (/** @type {Title} */(""));
        }
        if (!("Name" in $$source)) {
            /**
             * @member
             * @type {string}
             */
            this["Name"] = "";
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Person instance from a string or object.
     * @param {any} [$$source = {}]
     * @returns {Person}
     */
    static createFrom($$source = {}) {
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        return new Person(/** @type {Partial<Person>} */($$parsedSource));
    }
}

/**
 * Title is a title
 * @readonly
 * @enum {string}
 */
export const Title = {
    /**
     * The Go zero value for the underlying type of the enum.
     */
    $zero: "",

    /**
     * Mister is a title
     */
    Mister: "Mr",
    Miss: "Miss",
    Ms: "Ms",
    Mrs: "Mrs",
    Dr: "Dr",
};

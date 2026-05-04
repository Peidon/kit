const test = require("node:test");
const assert = require("node:assert/strict");
const fs = require("node:fs");
const path = require("node:path");
const vm = require("node:vm");

function extractFunctionSource(source, functionName) {
    const start = source.indexOf(`function ${functionName}(`);
    if (start === -1) {
        throw new Error(`Function ${functionName} not found`);
    }

    const bodyStart = source.indexOf("{", start);
    let depth = 0;

    for (let index = bodyStart; index < source.length; index++) {
        const char = source[index];
        if (char === "{") {
            depth++;
        } else if (char === "}") {
            depth--;
            if (depth === 0) {
                return source.slice(start, index + 1);
            }
        }
    }

    throw new Error(`Function ${functionName} is not balanced`);
}

function loadNormalize() {
    const source = fs.readFileSync(path.join(__dirname, "content.js"), "utf8");
    const normalizeSource = extractFunctionSource(source, "normalize");
    const context = {};

    vm.createContext(context);
    vm.runInContext(`
${normalizeSource}
this.normalize = normalize;
`, context);

    return context.normalize;
}

function loadSplitIntoPhrases() {
    const source = fs.readFileSync(path.join(__dirname, "content.js"), "utf8");
    const splitSource = extractFunctionSource(source, "splitIntoPhrases");
    const context = {};

    vm.createContext(context);
    vm.runInContext(`
${splitSource}
this.splitIntoPhrases = splitIntoPhrases;
`, context);

    return context.splitIntoPhrases;
}

const splitIntoPhrases = loadSplitIntoPhrases();

test("splitIntoPhrases splits camelCase words into separate tokens", () => {
    assert.equal(splitIntoPhrases("workExperience").join(" "), "work experience");
});

test("splitIntoPhrases splits PascalCase words into separate tokens", () => {
    assert.equal(splitIntoPhrases("WorkExperience").join(" "), "work experience");
});

test("splitIntoPhrases keeps consecutive uppercase letters together", () => {
    assert.equal(splitIntoPhrases("URLAddress").join(" "), "url address");
    assert.equal(splitIntoPhrases("MyURLAddress").join(" "), "my url address");
});

test("splitIntoPhrases lowercases text and replaces symbols with spaces", () => {
    assert.equal(
        splitIntoPhrases("workExperience-6--endDate-dateSectionYear-input").join(","),
        "work experience,end date,date section year,input"
    );
});

const normalize = loadNormalize();

test("normalize splits camelCase words into separate tokens", () => {
    assert.equal(normalize("workExperience"), "work experience");
});

test("normalize splits PascalCase words into separate tokens", () => {
    assert.equal(normalize("WorkExperience"), "work experience");
});

test("normalize keeps consecutive uppercase letters together", () => {
    assert.equal(normalize("URLAddress"), "url address");
    assert.equal(normalize("MyURLAddress"), "my url address");
});

test("normalize lowercases text and replaces symbols with spaces", () => {
    assert.equal(
        normalize("workExperience-6--endDate-dateSectionYear-input"),
        "work experience 6  end date date section year input"
    );
});

test("normalize returns an empty string for empty input", () => {
    assert.equal(normalize(""), "");
});

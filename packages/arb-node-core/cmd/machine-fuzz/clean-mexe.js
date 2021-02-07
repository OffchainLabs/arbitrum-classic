#!/usr/bin/env node
const fs = require("fs");

const path = process.argv[2] || "/dev/stdin";
const program = JSON.parse(fs.readFileSync(path));

for (let instruction of program.code) {
    if (typeof instruction.opcode === "number") {
        instruction.opcode = instruction.opcode & 0xff;
    }
    for (let key of Object.keys(instruction)) {
        if (key.length > 16) {
            delete instruction[key];
        }
    }
}

// Format JSON nicely
const code = program.code;
delete program.code;
let out = JSON.stringify(program);
out = out.slice(0, out.length - 1);
out += ",\"code\":["
let first = true;
for (let instruction of code) {
    if (first) {
        first = false;
    } else {
        out += ",";
    }
    out += "\n  ";
    out += JSON.stringify(instruction);
}
out += "\n]}";
console.log(out);

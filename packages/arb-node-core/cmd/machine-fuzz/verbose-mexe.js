#!/usr/bin/env node
const fs = require("fs");

const path = process.argv[2] || "/dev/stdin";
const program = JSON.parse(fs.readFileSync(path));

for (let instruction of program.code) {
    if (typeof instruction.opcode === "number") {
        instruction.opcode = {AVMOpcode: instruction.opcode};
    }
    instruction.debug_info = {attributes: {breakpoint: false, inline: false}};
}
program.file_name_chart = {};

console.log(JSON.stringify(program));

// Simple JavaScript examples to explain common operators
// Save as /Users/mohsin/MCA/staffproof/index.js and run with `node index.js`

'use strict';

// Arithmetic operators
console.log('Arithmetic:');
console.log('5 + 2 =', 5 + 2);
console.log('5 - 2 =', 5 - 2);
console.log('5 * 2 =', 5 * 2);
console.log('5 / 2 =', 5 / 2);
console.log('5 % 2 =', 5 % 2); // remainder
console.log('2 ** 3 =', 2 ** 3); // exponentiation

// Increment / Decrement (pre vs post)
let a = 5;
console.log('\nIncrement/Decrement:');
console.log('a (start) =', a);
console.log('++a (pre-increment) =', ++a); // increments, then returns
console.log('a++ (post-increment) =', a++); // returns, then increments
console.log('a (after) =', a);

// Assignment operators
let b = 10;
console.log('\nAssignment:');
console.log('b =', b);
b += 5; // b = b + 5
console.log('b += 5 ->', b);
b *= 2; // b = b * 2
console.log('b *= 2 ->', b);

// Comparison operators
console.log('\nComparison:');
console.log("5 == 5 ->", 5 == 5);   // loose equality (coerces types)
console.log("'5' == 5 ->", '5' == 5); // strict equality (no coercion)
console.log('3 > 2 ->', 3 > 2);
console.log('3 >= 3 ->', 3 >= 3);
console.log('3 < 2 ->', 3 < 2);

// Logical operators and short-circuiting
console.log('\nLogical operators:');
const truthy = true;
const falsy = false;
function sideEffect() {
    console.log('sideEffect() called');
    return 'result';
}
console.log('true && sideEffect() ->', truthy && sideEffect()); // calls sideEffect
console.log('false && sideEffect() ->', falsy && sideEffect()); // short-circuits, not called
console.log('false || sideEffect() ->', falsy || sideEffect()); // calls sideEffect because left is falsy

// Ternary operator
const age = 17;
console.log('\nTernary:');
console.log('age >= 18 ? "adult" : "minor" ->', age >= 18 ? 'adult' : 'minor');

// Unary operators
console.log('\nUnary operators:');
console.log("typeof 'hello' ->", typeof 'hello');
console.log("+'42' ->", +'42'); // unary plus converts string to number
console.log("!true ->", !true); // logical NOT

// Example combining operators
console.log('\nCombined example:');
const x = 4;
const y = '4';
const result = (x === +y && x % 2 === 0) ? 'even and equal' : 'not both';
console.log('x === +y && x % 2 === 0 ->', result);

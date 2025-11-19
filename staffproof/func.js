// let i = 12
// if (i % 2 === 0) {
//     console.log("Even")
// } else {
//     console.log("Odd")
// }

// let j= 245; 
// if (j % 2 === 0) {
//     console.log("Even")
// } else {
//     console.log("Odd")
// }

// let k = 1234567890;
// if (k % 2 === 0) {
//     console.log("Even")
// } else {
//     console.log("Odd")
// }

//Function for odd and even
// function abc(num) {
//     console.log(num)
//     if (num % 2 === 0) {
//         return num/2;
//     } else {
//         return num*3;
//     }
// }



// console.log(abc(12) + " is the result")
// console.log(abc(245) + " is the result")
// console.log(abc(1234567890) + " is the result")

// Reverse a string - Create a function that reverses a given string.

// let a = 12;
// let b = 13;

// //function to add two nos
// function add(a, b) {
//     console.log(a, b + "is from function")
//     return a - b;
// }

// console.log(a, b + "is from  main")
// console.log(add(b, a) + " is the result")
// console.log(add(14,17))

// let abc= "hello";


// function add(y, x, z){
//     let adding = x+y+z;
//     return adding;
// }

// let abc = 40;
// let adding = add(13,13,14);
// console.log(adding)

//Create a function that removes consecutive duplicate characters from a string while preserving order.

let str = "aaabbbccddd";
// let result = "abcd";


    let game = "";
    for (let i = 0; i < str.length; i++) {
        if (str[i] !== str[i+1]) {
            game += str[i];
        }
    }



console.log(game);


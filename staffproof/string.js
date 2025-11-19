// let text = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
// console.log(text.length)

// for (let i = 0; i < text.length ; i++) {
//     if (i % 2 === 0) {
//         console.log(text[i])
//     } 
// }

// let firstName = "John";
// let lastName = "Doe";
// let full = firstName + " " + lastName
// let fullName = lastName.concat(" ",firstName)
// console.log(fullName)
// console.log(full)


// let text = "Hello World";
// console.log(text.slice(0,5))
// console.log(text.slice(6))
// console.log(text.slice(-8))
// console.log(text.slice(-5, 1))
// console.log(text.slice(-5, -1))


// negetive indexes
// 0 is starting from left to right
// -1 is starting from right to left
// -2 is starting from right to left
// -3 is starting from right to left
// -4 is starting from right to left
// -5 is starting from right to left
// -6 is starting from right to left
// -7 is starting from right to left
// -8 is starting from right to left
// -9 is starting from right to left
// -10 is starting from right to left

// let text = "apple banana cherry date elderberry fig grape honeydew";
// console.log(text)
// console.log(text.split(" ")) // split the string into an array of words
// console.log(text.split(" ", 10)) // split the string into an array of words and limit the number of words to 3


// let username = "John Doe123123123";
// if (username.length > 10) {
//     console.log("Username is too long")
// } else {
//     console.log("Username is too short")
// }

// let fruit1 = "kiwi";
// let fruit2 = "banana";
// let fruit3 = "cherry";
// let fruitbasket = fruit1 + " " + fruit2 + " " + fruit3;
// console.log(fruitbasket)
// if (fruitbasket.includes("apple")) {
//     console.log("Apple is in the fruit basket")
// } else {
//     console.log("Apple is not in the fruit basket")
// }

// let str = "aeiouAEIOU";
// let count = 0;
// for (let i = 0; i < str.length; i++) {
//     if("a".includes(str[i])) {
//         count++;
//     }
//     if("e".includes(str[i])) {
//         count++;
//     }
//     if("i".includes(str[i])) {
//         count++;
//     }
//     if("o".includes(str[i])) {
//         count++;
//     }
//     if("u".includes(str[i])) {
//         count++;
//     }
// }
// console.log("The number of vowels in the string is " + count);


let hi = "Mohsin";
let hello = hi.length;
let newer = "";
console.log(hello); 

for (let i = hi.length-1; i >= 0; i--) {
    newer += hi[i];
}
console.log(newer.toLowerCase());
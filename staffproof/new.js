// Program to find prime numbers using nested for loops and if statements






for (let num = 2; num <= 100; num++) {
    let isPrime = true;
    
    // Nested for loop to check if num is prime
    for (let divisor = 2; divisor < num; divisor++) {
        if (num % divisor === 0) {
            isPrime = false;
            break; // Not prime, exit inner loop
        }
    }
    
    // If statement to print prime numbers
    if (isPrime) {
        console.log(num);
    }
}

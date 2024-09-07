// Define a function
function greet(name) {
  return `Hello, ${name}!`;
}

// Use the function
console.log(greet("Goja"));

// Perform some calculations
let a = 5;
let b = 3;
console.log(`${a} + ${b} = ${a + b}`);
console.log(`${a} * ${b} = ${a * b}`);

// Create an object
let person = {
  name: "Alice",
  age: 30,
  sayHi: function () {
    console.log(`Hi, I'm ${this.name} and I'm ${this.age} years old.`);
  },
};

// Use the object
person.sayHi();

// Demonstrate a loop
for (let i = 0; i < 5; i++) {
  console.log(`Iteration ${i + 1}`);
}

// Use a built-in JavaScript method
let fruits = ["apple", "banana", "cherry"];
console.log("Fruits:", fruits.join(", "));

// Demonstrate error handling
try {
  throw new Error("This is a test error");
} catch (error) {
  console.log("Caught an error:", error.message);
}

console.log("Script execution completed!");

const readline = require('readline');

var r1 = readline.createInterface({
  input: process.stdin,
//  output: process.stdout
});

function sum2020(x, y) {
  //console.log(x + " " + y + " = " + x+y);
  return (x+y) == 2020;
}

let numbers = [];

r1.on('line', function(line) {
    // Put the number in an array
    num = parseInt(line);
    for (var i = numbers.length - 1; i>=0; i--) {
      if (sum2020(num, numbers[i])) {
        console.log(num + " " + numbers[i]);
        console.log(num * numbers[i]);
      }
    }
    numbers.push(num);
});



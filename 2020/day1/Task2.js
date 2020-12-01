const readline = require('readline');

var r1 = readline.createInterface({
  input: process.stdin,
//  output: process.stdout
});

function sum2020(x, y, z) {
  //console.log(x + " " + y + " = " + x+y);
  return (x+y+z) == 2020;
}

let numbers = [];

r1.on('line', function(line) {
    // Put the number in an array
    num = parseInt(line);
    for (var i = numbers.length - 1; i>=0; i--) {
      for (var j = i; j>0; j--) {
        if (sum2020(num, numbers[i], numbers[j])) {
          console.log(num + " " + numbers[i] + " " + numbers[j]);
          console.log(num * numbers[i] * numbers[j]);
        }
      }
    }
    numbers.push(num);
});



const readline = require("readline");

const rl = readline.createInterface({
  input: process.stdin,

  output: process.stdout,
});

fact = (number) => {
  if (number == 0) {
    return 1;
  } else {
    return number * fact(number - 1);
  }
};

rl.question("", (number) => {
  console.log(fact(+number));

  rl.close();
});

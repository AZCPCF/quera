const readline = require("readline");

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

rl.question("", (answer) => {
  const textArr = Array.from(answer);
  let count = 1;
  let isBad = false;

  for (let i = 1; i < textArr.length; i++) {
    if (textArr[i] === textArr[i - 1]) {
      count++;
    } else {
      if (count % 2 !== 0) {
        isBad = true;
        break;
      }
      count = 1;
    }
  }

  // Check the last sequence
  if (count % 2 !== 0) {
    isBad = true;
  }

  console.log(isBad ? "bad" : "khoob");
  rl.close();
});

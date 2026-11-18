const CHANGE_TYPE = {
  UP: "UP",
  DOWN: "DOWN",
};
const ERROR_TYPE = {
  NOT_FOUND: "NOT_FOUND",
  NOT_POSSIBLE: "NOT_POSSIBLE",
  INVALID_INPUT: "INVALID_INPUT",
};
let numbers = [4, 6, 10, 23, 0, 24, 30, 2];

const numbersContainer = document.getElementById("numbers-container");
const errorContainer = document.getElementById("error-container");
const itemInput = document.getElementById("item-input");
const countInput = document.getElementById("count-input");
const submitBtn = document.getElementById("submit-btn");

function renderNumbers() {
  numbersContainer.innerHTML = "";
  numbers.forEach((num) => {
    const span = document.createElement("span");
    span.textContent = num;
    numbersContainer.appendChild(span);
  });
}

function clearError() {
  const old = document.getElementById("error");
  if (old) old.remove();
}

function showError(type) {
  clearError();
  const p = document.createElement("p");
  p.id = "error";
  p.textContent = type;
  errorContainer.appendChild(p);
}

function handleSubmit() {
  clearError();

  const itemValue = itemInput.value;
  const countValue = countInput.value;

  // 1) INVALID_INPUT
  if (itemValue === "" || countValue === "") {
    showError(ERROR_TYPE.INVALID_INPUT);
    return;
  }

  const item = Number(itemValue);
  const count = Number(countValue);

  // 2) NOT_FOUND
  const index = numbers.indexOf(item);
  if (index === -1) {
    showError(ERROR_TYPE.NOT_FOUND);
    return;
  }

  // 3) NOT_POSSIBLE
  const type = document.querySelector('input[name="type"]:checked').value;
  const newIndex = type === CHANGE_TYPE.UP ? index + count : index - count;

  if (newIndex < 0 || newIndex > numbers.length - 1) {
    showError(ERROR_TYPE.NOT_POSSIBLE);
    return;
  }

  numbers.splice(index, 1);
  numbers.splice(newIndex, 0, item);
  renderNumbers();
}

renderNumbers();
submitBtn.addEventListener("click", handleSubmit);
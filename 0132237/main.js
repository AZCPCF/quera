const leftSide = document.querySelector(".left-side");
const rightSide = document.querySelector(".right-side");
const midSide = document.querySelector(".mid-side");
const [allToRight, checkedToRight, checkedToLeft, allToLeft] = midSide.children;

let leftList = [
  { id: "item1", checked: false, title: "PHP" },
  { id: "item2", checked: false, title: "Python" },
  { id: "item3", checked: false, title: "Ruby" },
  { id: "item4", checked: false, title: "C++" },
];
let rightList = [
  { id: "item5", checked: false, title: "HTML" },
  { id: "item6", checked: false, title: "Css" },
  { id: "item7", checked: false, title: "JavaScript" },
  { id: "item8", checked: false, title: "Java" },
];

renderDom(leftList, rightList);

// Render Dom
function renderDom(leftListToRender, rightListToRender) {
  leftListToRender.forEach((item) => {
    leftSide.innerHTML += `<div class="box">
        <input type="checkbox" class="input-box" id="${item.id}" />
        <label for="${item.id}">${item.title}</label>
        </div>`;
  });

  rightListToRender.forEach((item) => {
    rightSide.innerHTML += `<div class="box">
          <input type="checkbox" class="input-box" id="${item.id}" />
          <label for="${item.id}">${item.title}</label>
          </div>`;
  });

  registerEvents();
}

// Clear Dom
function clearDom() {
  document.querySelectorAll(".side").forEach((el) => {
    el.innerHTML = "";
  });
}

function registerEvents() {
  [...leftList, ...rightList].forEach((item) => {
    const input = document.getElementById(item.id);
    input.checked = item.checked;
    input.onchange = () => {
      item.checked = input.checked;
    };
  });

  const leftEmpty = leftList.length === 0;
  const rightEmpty = rightList.length === 0;
  allToRight.classList.toggle("disabled", leftEmpty);
  checkedToRight.classList.toggle("disabled", leftEmpty);
  allToLeft.classList.toggle("disabled", rightEmpty);
  checkedToLeft.classList.toggle("disabled", rightEmpty);

  allToRight.onclick = () => {
    if (allToRight.classList.contains("disabled")) return;
    move("right", false);
  };
  checkedToRight.onclick = () => {
    if (checkedToRight.classList.contains("disabled")) return;
    move("right", true);
  };
  allToLeft.onclick = () => {
    if (allToLeft.classList.contains("disabled")) return;
    move("left", false);
  };
  checkedToLeft.onclick = () => {
    if (checkedToLeft.classList.contains("disabled")) return;
    move("left", true);
  };
}

function move(direction, onlyChecked) {
  const from = direction === "right" ? leftList : rightList;
  const to = direction === "right" ? rightList : leftList;

  const moving = onlyChecked ? from.filter((i) => i.checked) : [...from];
  const staying = onlyChecked ? from.filter((i) => !i.checked) : [];

  moving.forEach((i) => (i.checked = false));
  to.push(...moving);

  if (direction === "right") {
    leftList = staying;
    rightList = to;
  } else {
    rightList = staying;
    leftList = to;
  }

  clearDom();
  renderDom(leftList, rightList);
}
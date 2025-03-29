let toggle_btn_img = () => {};
document.addEventListener("DOMContentLoaded", () => {
  const button = document.getElementById("btn");
  const avatar = document.getElementById("avatar");
  toggle_btn_img = () => {
    if (button.innerText === "Show!") {
      avatar.classList.remove("w3-hide");
      avatar.classList.add("w3-show");
      button.innerText = "Hide!";
      return;
    }
    avatar.classList.remove("w3-show");
    avatar.classList.add("w3-hide");
    button.innerText = "Show!";
  };
});

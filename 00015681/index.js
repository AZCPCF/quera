document.addEventListener("DOMContentLoaded", function () {
  fetch("count.json")
    .then((response) => {
      return response.json();
    })
    .then((data) => {
      const likesCount = document.getElementById("likes");
      likesCount.textContent = data.count;
    })
    .catch((error) => {
      console.error("Error:", error);
    });
});

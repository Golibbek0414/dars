window.addEventListener('DOMContentLoaded', (_) => {
    form.addEventListener("submit", function (event) {
      event.preventDefault();
      let username = document.getElementById("input-username");
      let text = document.getElementById("input-text");
      text.value = "";
    });
   });
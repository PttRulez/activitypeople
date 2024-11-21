let modal = document.getElementById("daisymodal");

if (modal) {
  modal.addEventListener("click", function (e) {
    if (e.target === modal) {
      e.target.close();
    }
  });
}

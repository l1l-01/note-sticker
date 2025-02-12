function select_one(selector) {
  return document.querySelector(selector);
}
function addClass(ele, cssClass) {
  ele.classList.add(cssClass);
}
function removeClass(ele, cssClass) {
  ele.classList.remove(cssClass);
}
const closeBtn = select_one(".cancel"),
  form = select_one("form"),
  addBtn = select_one(".add");

closeBtn.addEventListener("click", () => {
  removeClass(form, "active");
});

addBtn.addEventListener("click", () => {
  addClass(form, "active");
});

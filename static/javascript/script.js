function selectOne(selector) {
  return document.querySelector(selector);
}
function selectAll(selector) {
  return document.querySelectorAll(selector);
}
function addClass(ele, cssClass) {
  ele.classList.add(cssClass);
}
function removeClass(ele, cssClass) {
  ele.classList.remove(cssClass);
}

function addEvent(ele, event, handler) {
  ele.addEventListener(event, handler);
}

function removeEvent(element, event, handler) {
  element.removeEventListener(event, handler);
}

const closeBtn = selectOne(".cancel"),
  form = selectOne("form"),
  addBtn = selectOne(".add");

addEvent(addBtn, "click", () => {
  addClass(form, "active");
});

addEvent(closeBtn, "click", () => {
  removeClass(form, "active");
});

const move_eles = selectAll(".move"),
  notes = selectAll(".note"),
  moveImg = selectOne(".move img");

move_eles.forEach((ele, key) => {
  let isMoving = false;

  const moveHandler = (event) => {
    notes[key].style.left = event.clientX - 100 + "px";
    notes[key].style.top = event.clientY - 25 + "px";
  };

  addEvent(ele, "click", () => {
    isMoving = !isMoving;

    if (isMoving) {
      addEvent(document, "mousemove", moveHandler);
    } else {
      removeEvent(document, "mousemove", moveHandler);
    }
  });
});

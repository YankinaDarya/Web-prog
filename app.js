var tag_ul = document.createElement("ul");

add_task("Сделать задание #3 по web-программированию");

function del() {
    this.parentNode.remove();
}

document.getElementById('root').appendChild(tag_ul);
var input = document.createElement("input");
input.id = "add_task_input";
var but2 = document.createElement("button");
but2.id = "add_task";
document.getElementById('root').appendChild(input);
document.getElementById('root').appendChild(but2);
but2.innerHTML = "Добавить";
but2.addEventListener('click', function() {add_task(input.value)});
function add_task(string_input) {
    var tag_li = document.createElement("li");

    var span = document.createElement("span");
    span.innerHTML = string_input;

    var but = document.createElement("button");
    but.innerHTML = "Удалить";
    tag_ul.appendChild(tag_li);
    tag_li.appendChild(span);
    tag_li.appendChild(but);
    but.addEventListener('click', del);
}
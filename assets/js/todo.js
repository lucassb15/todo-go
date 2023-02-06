// cria um novo elemento(task) na lista
add = () => {
    let li = document.createElement('LI');
    let input_value = document.todo_form.task.value;
    let input_text = document.createTextNode(input_value);

    document.querySelector('ul').addEventListener('click', (e) => {
        if (e.target.tagName === 'LI')
            e.target.classList.toggle('done');
    });

    li.appendChild(input_text);
    document.querySelector('ul').appendChild(li);
    document.todo_form.task.value = "";

    
}   
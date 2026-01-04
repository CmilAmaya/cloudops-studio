const form = document.getElementById("todo-form");
const input = document.getElementById("todo-input");
const list = document.getElementById("todo-list");

let todos = [];
let nextId = 1;

form.addEventListener("submit", (e) => {
  e.preventDefault();

  const text = input.value.trim();
  if (!text) return;

  addTodo(text);
  input.value = "";
});

function addTodo(text) {
  const todo = {
    id: nextId++,
    text,
    completed: false,
  };

  todos.push(todo);
  renderTodos();
}

function toggleTodo(id) {
  todos = todos.map(todo =>
    todo.id === id ? { ...todo, completed: !todo.completed } : todo
  );
  renderTodos();
}

function deleteTodo(id) {
  todos = todos.filter(todo => todo.id !== id);
  renderTodos();
}

function renderTodos() {
  list.innerHTML = "";

  todos.forEach(todo => {
    const item = document.createElement("div");
    item.className = `todo-item ${todo.completed ? "completed" : ""}`;

    item.innerHTML = `
      <button class="check">✓</button>
      <span>${todo.text}</span>
      <button class="delete">🗑</button>
    `;

    item.querySelector(".check").onclick = () => toggleTodo(todo.id);
    item.querySelector(".delete").onclick = () => deleteTodo(todo.id);

    list.appendChild(item);
  });
}

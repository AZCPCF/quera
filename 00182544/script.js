const ul = document.querySelector("#post-list");
const getData = async () => {
  const data = await (await fetch("http://localhost:3000/posts")).json();
  data.forEach((item) => {
    const li = document.createElement("li");
    const h3 = document.createElement("h3");
    h3.innerHTML = item.title;
    const p = document.createElement("p");
    p.innerHTML = item.body;
    const em = document.createElement("em");
    em.innerHTML = `شماره  ${item.id}`;
    li.appendChild(h3);
    li.appendChild(p);
    li.appendChild(em);
    ul.appendChild(li);
  });
};
getData();

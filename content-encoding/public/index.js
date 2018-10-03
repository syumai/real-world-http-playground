const button = document.createElement('button');
button.textContent = 'click me!';
button.addEventListener('click', () => console.log('hello, world!'));
document.body.appendChild(button);

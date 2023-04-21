'use strict';

const toggleContainer = document.getElementById('toggle');
const toggleButton = document.getElementById('toggle-button');

toggleButton.addEventListener('click', () => {
    toggleContainer.classList.toggle('end-justify');
})
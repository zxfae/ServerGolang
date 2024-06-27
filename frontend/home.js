import { loadCategories } from './categories.js';
import { loadPosts } from './posts.js';

let cachedUsers = null;

export function loadHome() {
    console.log('Loading home...');
    if (cachedUsers) {
        renderHome();
    } else {
        fetch('/api/users')
            .then(response => response.json())
            .then(data => {
                console.log('Users loaded:', data);
                cachedUsers = data;
                renderHome();
            })
            .catch(error => console.error('Error fetching user data:', error));
    }
}

function renderHome() {
    const content = document.getElementById('content');
    content.innerHTML = `
        <div class="home-message">
            <!-- home content -->
        </div>
        <p class="white-text">Categories</p>
        <div class="categories-container"></div>
        <p class="white-text">Lasts Posts</p>
        <div class="posts-container"></div>
    `;
    loadCategories();
    loadPosts();
}

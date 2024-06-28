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
    console.log('Rendering home...');
    const content = document.getElementById('content');
    if (!content) {
        console.error('Content element not found');
        return;
    }

    content.innerHTML = `
        <div class="home-message">
            <div class="full-height-box form-container">
                <h1 class="white-text">Welcome To <span>IFHK</span> Forum</h1>
                <p class="white-text">This is your HomeBoard. Enjoy your stay!</p>
            </div>
        </div>
        <p class="white-text">Categories</p>
        <div class="categories-container"></div>
        <div class="PostForm-container">
        <p class="white-text">Create your post</p>
            <div class="textArea">
                <form id="postForm">
                    <div class="Title">
                        <label for="category" class="white-text">Category:</label>
                            <select name="category" id="cat-select" required>
                            <option value="">--Please choose an category--</option>
                            <option value="Presentation">Presentation</option>
                            <option value="Programmation">Programmation</option>
                            <option value="Linux">Linux</option>
                            <option value="TechNews">TechNews</option>
                            <option value="Artificial-intelligence">Artificial Intelligence</option>
                            <option value="Mathematical">Mathematical</option>
                            <option value="Tutorials">Tutorials</option>
                            <option value="CyberSecurity">CyberSecurity</option>
                            <option value="Exercises">Exercises</option>
                            <option value="Moderation">Moderation</option>
                            <option value="Coffeebar">CoffeeBar</option>
                            <option value="Others">Others</option>
                        </select>
                        <label for="title" class="white-text">Title:</label>
                        <input type="text" id="title" name="title" required><br><br>
                        <label for="description" class="white-text">Description:</label>
                        <textarea class="textarea" id="description" name="description" required></textarea>
                        <input type="hidden" id="category" name="category">
                        <div class="Send-button">
                        <button type="submit">Send</button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
        <p class="white-text">Lasts Posts</p>
        <div class="posts-container"></div>
    `;
    console.log('Home content set');
    document.getElementById('postForm').addEventListener('submit', submitForm);
    loadCategories();
    loadPosts();

}

function submitForm(event) {
    event.preventDefault();
    const form = event.target;
    const formData = new FormData(form);

    fetch('/created', {
        method: 'POST',
        body: new URLSearchParams(formData)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        } else if(response.ok){
            window.location.hash = '#home';
            loadHome();
        }
        return response.json();
    })
    .catch(error => {
        console.error('Error:', error);
    });
}
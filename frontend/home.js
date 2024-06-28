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
            <!-- home content -->
        </div>
        <p class="white-text">Categories</p>
        <div class="categories-container"></div>
        <div class="PostForm-container">
        <p class="white-text">Create your post</p>
            <div class="textArea">
                <form>
                    <div class="Title">
                        <label for="category" class="white-text">Category:</label>
                            <select name="category" id="cat-select" required>
                            <option value="">--Please choose an category--</option>
                            <option value="presentation">Presentation</option>
                            <option value="programmation">Programmation</option>
                            <option value="linux">Linux</option>
                            <option value="techNews">TechNews</option>
                            <option value="artificial-intelligence">Artificial Intelligence</option>
                            <option value="mathematical">Mathematical</option>
                            <option value="tutorials">Tutorials</option>
                            <option value="cybersecurity">CyberSecurity</option>
                            <option value="exercises">Exercises</option>
                            <option value="moderation">Moderation</option>
                            <option value="coffeebar">CoffeeBar</option>
                            <option value="others">Others</option>
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

    loadCategories();
    loadPosts();
}

document.addEventListener("DOMContentLoaded", () => {
    checkUserStatus();
    window.addEventListener('hashchange', handleRoute);
});

function checkUserStatus() {
    fetch('/api/checkAuth')
        .then(response => {
            if (response.ok) {
                window.isUserAuthenticated = true;
                if (!window.location.hash || window.location.hash === '#login' || window.location.hash === '#register') {
                    window.location.hash = '#home';
                }
            } else {
                window.isUserAuthenticated = false;
                if (!window.location.hash || window.location.hash === '#home') {
                    window.location.hash = '#login';
                }
            }
            handleRoute();
        })
        .catch(error => {
            console.error('Error checking user status:', error);
            window.isUserAuthenticated = false;
            handleRoute();
        });
}

function handleRoute() {
    const hash = window.location.hash || '#login';
    switch (hash) {
        case '#home':
            if (window.isUserAuthenticated) {
                loadHome();
            } else {
                window.location.hash = '#login';
            }
            break;
        case '#register':
            loadRegister();
            break;
        case '#login':
            loadLogin();
            break;
        case '#logout':
            logoutUser();
            break;
        default:
            window.location.hash = '#home';
            break;
    }
}

function navigate(event, route) {
    event.preventDefault();
    window.location.hash = route;
}

// Client Handling
function loadRegister() {
    document.getElementById('content').innerHTML = `
        <form id="registerForm" class="form-container">
            <label for="nickname" class="white-text">Nickname:</label>
            <input type="text" id="nickname" name="nickname" required><br><br>

            <label for="age" class="white-text">Age:</label>
            <input type="number" id="age" name="age" min="18" value="18"><br><br>

            <label for="gender" class="white-text">Gender:</label>
            <input type="text" id="gender" name="gender"><br><br>

            <label for="firstname" class="white-text">First Name:</label>
            <input type="text" id="firstname" name="firstname"><br><br>

            <label for="lastname" class="white-text">Last Name:</label>
            <input type="text" id="lastname" name="lastname"><br><br>

            <label for="email" class="white-text">Email:</label>
            <input type="email" id="email" name="email" required><br><br>

            <label for="password" class="white-text">Password:</label>
            <input type="password" id="password" name="password" required><br><br>

            <button type="submit">Register</button>
            <a href="#login" onclick="navigate(event, '#login')" class="logLink">Your account already exist? Click here</a>
        </form>
    `;
    document.getElementById('registerForm').addEventListener('submit', registerUser);
}

function loadLogin() {
    document.getElementById('content').innerHTML = `
        <form id="loginForm" class="form-container">
            <label for="nickname" class="white-text">Nickname or Email:</label>
            <input type="text" id="nickname" name="nickname" required><br><br>

            <label for="password" class="white-text">Password:</label>
            <input type="password" id="password" name="password" required><br><br>

            <button type="submit">Login</button>
            <a href="#register" onclick="navigate(event, '#register')" class="logLink">You don't have an account? Click here</a>
        </form>
    `;
    document.getElementById('loginForm').addEventListener('submit', loginUser);
}

function logoutUser() {
    fetch('/logout', { method: 'POST' })
        .then(response => {
            if (response.ok) {
                window.isUserAuthenticated = false;
                window.location.hash = '#login';
            } else {
                throw new Error('Logout failed');
            }
        })
        .catch(error => console.error('Error:', error));
}

function registerUser(event) {
    event.preventDefault();
    const form = event.target;
    const formData = new FormData(form);

    fetch('/register', {
        method: 'POST',
        body: new URLSearchParams(formData),
    })
    .then(response => response.text())
    .then(data => {
        alert(data);
        const loginData = new URLSearchParams({
            nickname: formData.get('nickname'),
            password: formData.get('password')
        });
        fetch('/login', {
            method: 'POST',
            body: loginData,
        })
        .then(response => {
            if (response.ok) {
                window.isUserAuthenticated = true;
                window.location.hash = '#home';
            } else {
                alert('Auto-login failed after registration');
                window.location.hash = '#login';
            }
        })
        .catch(error => console.error('Error during auto-login:', error));
    })
    .catch(error => console.error('Error:', error));
}

function loginUser(event) {
    event.preventDefault();
    const form = event.target;
    const formData = new FormData(form);

    fetch('/login', {
        method: 'POST',
        body: new URLSearchParams(formData),
    })
    .then(response => {
        if (response.ok) {
            window.isUserAuthenticated = true;
            window.location.hash = '#home';
        } else {
            alert('Login failed');
        }
    })
    .catch(error => console.error('Error:', error));
}

let cachedUsers = null;
let cachedCategories = null;
let cachedPosts = null;

function loadHome() {
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
            <div class="full-height-box form-container">
                <h1 class="white-text">Welcome To <span>IFHK</span> Forum</h1>
                <p class="white-text">This is your HomeBoard. Enjoy your stay!</p>
            </div>
        </div>
        <p class="white-text">Categories</p>
        <div class="categories-container"></div>
        <p class="white-text">Last Posts</p>
        <div class="posts-container"></div>
    `;
    loadCategories();
    loadPosts();
}

function loadCategories() {
    console.log('Fetching categories...');
    fetch('/api/categories')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(categories => {
            console.log('Categories loaded:', categories);
            const categoriesContainer = document.querySelector('.categories-container');
            categoriesContainer.innerHTML = '';
            const limitedCategories = categories.slice(0, 12);
            limitedCategories.forEach(category => {
                const categoryElement = document.createElement('div');
                categoryElement.className = 'category';
                categoryElement.innerHTML = `
                    <h2><a class="username">category</a><span class="termux">@IFHK</span>:<span class="SpanPost"><span class="blue">~</span><span class="white">$</span>${category.name}</h2>
                    <p>${category.description}</p>
                `;
                categoriesContainer.appendChild(categoryElement);
            });
        })
        .catch(error => console.error('Error loading categories:', error));
}

function loadPosts() {
    console.log('Fetching posts...');
    fetch('/api/posts')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(posts => {
            console.log('Posts loaded:', posts);
            const postContainer = document.querySelector('.posts-container');
            postContainer.innerHTML = '';
            const limitedPosts = posts.slice(0, 12);
            limitedPosts.forEach(post => {
                const postElement = document.createElement('div');
                postElement.className = 'posts';
                let truncatedTitle = post.title.length > 60 ? `${post.title.substring(0, 60)}...` : post.title;
                let truncatedDescription = post.description.length > 80 ? `${post.description.substring(0, 80)}...` : post.description;
                postElement.innerHTML = `
                    <h2><a class="username">${post.username}</a><span class="Termux">@IFHK</span>:<span class="SpanPost"><span class="blue">~</span><span class="white">$/post/</span>${truncatedTitle}</span></h2>
                    <div class="Weigth">
                    <p><a class="username">category</a>:<span class="SpanPost"><span class="blue">~</span><span class="white">$</span>${post.categoryname}</p>
                    </div>
                    <p>${truncatedDescription}</p>
                    <p><span class="time">${post.created}</span></p>
                `;
                postContainer.appendChild(postElement);
            });
        })
        .catch(error => console.error('Error loading posts:', error));
}

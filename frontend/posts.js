export function loadPosts() {
    fetch('/api/posts')
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(posts => {
            const postContainer = document.querySelector('.posts-container');
            postContainer.innerHTML = '';
            const limitedPosts = posts.slice(0, 12);
            limitedPosts.forEach(post => {
                const postElement = document.createElement('div');
                postElement.className = 'posts';
                let truncatedTitle = post.title.length > 60 ? `${post.title.substring(0, 60)}...` : post.title;
                let truncatedDescription = post.description.length > 80 ? `${post.description.substring(0, 80)}...` : post.description;
                postElement.innerHTML = `
                    <h2>
                        <a class="username">${post.username}</a>
                        <span class="Termux">@IFHK</span>:<span class="SpanPost"><span class="blue">~$</span><a href="#posts/${post.title}">${truncatedTitle}</a></span>
                    </h2>
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
export function loadPostsByPosts(postsName) {
    fetch(`/api/posts/posts/${postsName}`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(posts => {
            const content = document.getElementById('content');
            content.innerHTML = `
                <div class="posts-page">
                    <h2 class="LinkCat">Posts: ${postsName}</h2>
                    <div class="posts-container"></div>
                </div>
            `;
            const postContainer = document.querySelector('.posts-container');
            postContainer.innerHTML = '';
            posts.forEach(post => {
                const postElement = document.createElement('div');
                postElement.className = 'posts';

                postElement.innerHTML = `
                    <div class="scrolling">
                        <h2><a class="username">${post.username}</a><span class="Termux">@IFHK</span>:<span class="SpanPost"><span class="blue">~$</span><a href="#posts/${post.title}">${post.title}</a></span></h2>
                        <div class="Weigth">
                            <p><a class="username">category</a>:<span class="SpanPost"><span class="blue">~</span><span class="white">$</span>${post.categoryname}</p>
                        </div>
                        <p>${post.description}</p>
                        <p><span class="time">${post.created}</span></p>
                    </div>
                `;
                postContainer.appendChild(postElement);
            });
        })
        .catch(error => console.error('Error loading posts:', error));
}

export function loadPostsByCategory(categoryName) {
    fetch(`/api/posts/category/${categoryName}`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(posts => {
            const content = document.getElementById('content');
            content.innerHTML = `
            <div class="Cat">
            <h1><p><a class="username">category</a>:<span class="SpanPost"><span class="blue">~</span><span class="white">$</span>${categoryName}</p></h1>
            </div>
            <div class="textArea">
                <form>
                    <div class="Title">
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
                <div class="PostsBycat-wrapper">
                <div class="PostsBycat">
                    <div class="posts-container"></div>
                </div>
            </div>
            `;
            const postContainer = document.querySelector('.posts-container');
            postContainer.innerHTML = '';
            posts.forEach(post => {
                const postElement = document.createElement('div');
                postElement.className = 'posts';
                let truncatedTitle = post.title.length > 60 ? `${post.title.substring(0, 60)}...` : post.title;
                let truncatedDescription = post.description.length > 80 ? `${post.description.substring(0, 80)}...` : post.description;
                postElement.innerHTML = `
                    <div class="scrolling">
                        <h2><a class="username">${post.username}</a><span class="Termux">@IFHK</span>:<span class="SpanPost"><span class="blue">~</span><span class="white">$/post/</span><a href="#posts/${post.title}">${truncatedTitle}</a></span></h2>
                        <p>${truncatedDescription}</p>
                        <p><span class="time">${post.created}</span></p>
                    </div>
                `;
                postContainer.appendChild(postElement);
            });
            document.addEventListener('DOMContentLoaded', (event) => {
                const form = document.getElementById('postForm');
                const categoryInput = document.getElementById('category');
                categoryInput.value = categoryName;

                form.addEventListener('submit', (e) => {
                    e.preventDefault();
                    const formData = new FormData(form);
                    fetch('/api/posts', {
                        method: 'POST',
                        body: formData
                    })
                    .then(response => response.json())
                    .then(data => {
                        console.log('Success:', data);
                    })
                    .catch((error) => {
                        console.error('Error:', error);
                    });
                });
            });
        })
        .catch(error => console.error('Error loading posts:', error));
}


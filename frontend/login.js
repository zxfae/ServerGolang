export function loadLogin() {
    document.getElementById('content').innerHTML = `
        <form id="loginForm" class="form-container">
            <!-- form content -->
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

export function loginUser(event) {
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

export function logoutUser() {
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

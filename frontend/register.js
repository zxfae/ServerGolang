export function loadRegister() {
    document.getElementById('content').innerHTML = `
        <form id="registerForm" class="form-container">
            <!-- form content -->
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

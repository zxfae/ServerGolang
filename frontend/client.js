export function LogoutUser() {
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

export function RegisterUser(event) {
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

export function LoginUser(event) {
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
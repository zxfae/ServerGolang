class Router {
    constructor() {
        this.routes = {};
        this.init();
    }

    init() {
        // Check Session history
        window.addEventListener('popstate', () => this.route());
        window.addEventListener('DOMContentLoaded', () => this.route());
    }

    add(route, action) {
        this.routes[route] = action;
    }

    route() {
        const path = window.location.pathname;
        const action = this.routes[path];

        if (action) {
            action();
        } else {
            this.routes['/not-found']();
        }
    }
}

const router = new Router();

router.add('/', function() {
    document.getElementById('DOM').innerHTML = `
                     <h2>Signup</h2>
                     <div class="logSignForm">
                         <p><label>Nickname:</label>
                             <input type="text" id="nickname" placeholder='Type your nickname' required/>
                         </p>
                         <p><label>Age:</label>
                             <input type="number" id="age" min="18" placeholder='Type your age' required/>
                         </p>
                         <p><label>Gender:</label>
                             <input type="text" id="gender" placeholder='Type your gender : male, female or non binary' required/>
                         </p>
                         <p><label>Firstname:</label>
                             <input type="text" id="firstname" placeholder='Type your firstname' required/>
                         </p>
                         <p><label>Lastname:</label>
                             <input type="text" id="lastname" placeholder='Type your lastname' required/>
                         </p>
                         <p><label>Email:</label>
                             <input type="email" id="email" placeholder='Type your email' required/>
                         </p>
                         <p><label>Password:</label>
                             <input type="password" id="password" placeholder='Type your password' required/>
                         </p>
                     <p>
                         <button type="submit">Submit</button>
                     </p>
                     <a href="/login" onclick="event.preventDefault(); navigateTo('/login');">Already registered ? log in here</a>
                     </div>
                 `;
});
router.add('/login', function() {
    document.getElementById('DOM').innerHTML = `
                     <h2>Login</h2>
                     <div class="logSignForm"> 
                     <p>
                         <label>Nickname or Email:</label>
                         <input type="text" id="nickname" placeholder='Type your nickname or your email' required/>
                     </p>
                     <p>
                         <label>Password:</label>
                         <input type="password" id="password" placeholder='Type your password' required/>
                     </p>
                     <p>
                         <button type="submit">Submit</button>
                     </p>
                     <a href="/" onclick="event.preventDefault(); navigateTo('/');">You do not have an account? register here</a>
                     </div>
                 `;
});

router.add('/not-found', function() {
    document.getElementById('DOM').innerHTML = '<h1>404</h1><p>Page non trouvée.</p>';
});

function navigateTo(path) {
    window.history.pushState({}, path, window.location.origin + path);
    router.route();
}



// document.addEventListener('DOMContentLoaded', function() {
//     const socket = new WebSocket('ws://localhost:8080/ws');
//     socket.onopen = function(event) {
//         console.log('Connection established');
//     };

//     socket.onmessage = function(event) {
//         console.log('Message received:', event.data);
//     };

//     socket.onclose = function(event) {
//         console.log('Connection closed');
//     };

//     socket.onerror = function(error) {
//         console.error('WebSocket Error:', error);
//     };
// });

// function submitFormRegister(){
// const data = {
//     nickname : "nickname",
//     age : "age",
//     gender : "gender",
//     firstname : "firstname",
//     lastname : "lastname",
//     email : "email",
//     password : "password"
// }

// fetch("/register", {
//     method: "POST",
//     headers: {
//         "Content-Type":"application/json",
//     },
//     body:JSON.stringify(data)
// })
// .then(response => response.json())
//         .then(data => {
//             console.log('Success:', data);
//         })
//         .catch((error) => {
//             console.error('Error:', error);
//         });
// }

// function submitFormRegister(){
//     function submitFormRegister() {
//         const nickname = document.getElementById('nickname').value;
//         const age = document.getElementById('age').value;
//         const genders = document.getElementsByName('gender').value;
//         const firstname = document.getElementById('firstname').value;
//         const lastname = document.getElementById('lastname').value;
//         const email = document.getElementById('email').value;
//         const password = document.getElementById('password').value;
    
//         const userData = {
//             nickname,
//             age,
//             gender,
//             firstname,
//             lastname,
//             email,
//             password
//         };

//         // Now send the data to the backend
//         sendDataToBackend(userData);
//     }
// }

// function sendDataToBackend(userData) {
//     fetch('http://localhost:8080/register', {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json'
//         },
//         body: JSON.stringify(userData)
//     })
//     .then(response => response.json())
//     .then(data => {
//         console.log('Success:', data);
//     })
//     .catch((error) => {
//         console.error('Error:', error);
//     });
// }
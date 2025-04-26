package templates


func NotFound() string{
    var html = (
        `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>404 - Page Not Found</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            color: #333;
            margin: 0;
            padding: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            text-align: center;
        }
        .container {
            max-width: 600px;
            padding: 20px;
            background-color: #fff;
            border-radius: 10px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        h1 {
            font-size: 100px;
            margin: 0;
            color: #e74c3c;
        }
        p {
            font-size: 18px;
            margin: 10px 0;
            color: #555;
        }
        a {
            font-size: 16px;
            color: #3498db;
            text-decoration: none;
            font-weight: bold;
        }
        a:hover {
            text-decoration: underline;
        }
        .button {
            display: inline-block;
            padding: 10px 20px;
            margin-top: 20px;
            background-color: #3498db;
            color: #fff;
            text-decoration: none;
            border-radius: 5px;
            font-weight: bold;
        }
        .button:hover {
            background-color: #2980b9;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>404</h1>
        <p>Oops! The page you are looking for does not exist.</p>
        <p>It seems that the page you requested was moved or doesn't exist anymore.</p>
        <a href="/" class="button">Back to Home</a>
    </div>
</body>
</html>`)
    return (html)
}

func Login(action_url string) string {
	var html = (`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Login</title>
</head>
<style>
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

body {
    height: 100vh;
    background: white;
    display: flex;
    justify-content: center;
    align-items: center;
}

.login-container {
    background-color: #fff;
    padding: 2rem 2.5rem;
    border-radius: 12px;
    box-shadow: 0 8px 16px rgba(0,0,0,0.2);
    width: 100%;
    max-width: 400px;
}

h2 {
    text-align: center;
    margin-bottom: 1.5rem;
    color: #333;
}

.input-group {
    margin-bottom: 1.2rem;
}

.input-group label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 600;
    color: #555;
}

.input-group input {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #ccc;
    border-radius: 8px;
    outline: none;
    font-size: 1rem;
    transition: border 0.3s;
}

.input-group input:focus {
    border-color: rgb(31, 128, 218);
}

button {
    width: 100%;
    padding: 0.75rem;
    border: none;
    border-radius: 8px;
    background:rgb(31, 128, 218);
    color: #fff;
    font-size: 1rem;
    font-weight: bold;
    cursor: pointer;
    transition: background 0.3s;
}

button:hover {
    background: #57c7b6;
}

.error-message {
    margin-top: 1rem;
    color: red;
    text-align: center;
    display: none;
}

</style>
<body>
    <div class="login-container">
        <h2>Login</h2>
        <form id="loginForm">
            <div class="input-group">
                <label for="username">Username</label>
                <input type="text" id="username" name="username" placeholder="Enter your email" required>
            </div>
            <div class="input-group">
                <label for="password">Password</label>
                <input type="password" id="password" name="password" placeholder="Enter your password" required>
            </div>
            <button type="submit">Login</button>
            <p class="error-message" id="errorMessage"></p>
        </form>
    </div>

    <script>
    const form = document.getElementById('loginForm');
    const errorMessage = document.getElementById('errorMessage');

    form.addEventListener('submit', function(e) {
        e.preventDefault(); // Prevent default form submission

        const username = document.getElementById('username').value.trim();
        const password = document.getElementById('password').value.trim();

        if (username === '' || password === '') {
            showError('Please fill in all fields.');
            return;
        }

        const formData = {
            username: username,
            password: password
        };

        fetch('`+action_url+`', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData)
        })
        .then(response => {
            if (response.status === 200) {
                alert('Login successful!');
                window.location.href = '/'; // Redirect to dashboard
            } else {
                return response.json();
            }
        })
        .then(data => {
            if (data && data.error) {
                showError(data.error); // Display error message if any
            }
        })
        .catch(error => {
            showError('An error occurred. Please try again.');
            console.error('Error during login:', error);
        });
    });

    function showError(message) {
        errorMessage.style.display = 'block';
        errorMessage.textContent = message;
    }
    </script>
</body>
</html>
	`)

	return html
}

func Register(action_url string) string {
	var html = (`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Register</title>
</head>
<style>
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

body {
    height: 100vh;
    background: white;
    display: flex;
    justify-content: center;
    align-items: center;
}

.login-container {
    background-color: #fff;
    padding: 2rem 2.5rem;
    border-radius: 12px;
    box-shadow: 0 8px 16px rgba(0,0,0,0.2);
    width: 100%;
    max-width: 400px;
}

h2 {
    text-align: center;
    margin-bottom: 1.5rem;
    color: #333;
}

.input-group {
    margin-bottom: 1.2rem;
}

.input-group label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 600;
    color: #555;
}

.input-group input {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #ccc;
    border-radius: 8px;
    outline: none;
    font-size: 1rem;
    transition: border 0.3s;
}

.input-group input:focus {
    border-color: rgb(31, 128, 218);
}

button {
    width: 100%;
    padding: 0.75rem;
    border: none;
    border-radius: 8px;
    background:rgb(31, 128, 218);
    color: #fff;
    font-size: 1rem;
    font-weight: bold;
    cursor: pointer;
    transition: background 0.3s;
}

button:hover {
    background: #57c7b6;
}

.error-message {
    margin-top: 1rem;
    color: red;
    text-align: center;
    display: none;
}

</style>
<body>
    <div class="login-container">
        <h2>Register</h2>
        <form id="loginForm">
            <div class="input-group">
                <label for="username">Username</label>
                <input type="text" id="username" name="username" placeholder="Enter your email" required>
            </div>
            <div class="input-group">
                <label for="password">Password</label>
                <input type="password" id="password" name="password" placeholder="Enter your password" required>
            </div>
            <button type="submit">Login</button>
            <p class="error-message" id="errorMessage"></p>
        </form>
    </div>

    <script>
    const form = document.getElementById('loginForm');
    const errorMessage = document.getElementById('errorMessage');

    form.addEventListener('submit', function(e) {
        e.preventDefault(); // Prevent default form submission

        const username = document.getElementById('username').value.trim();
        const password = document.getElementById('password').value.trim();

        if (username === '' || password === '') {
            showError('Please fill in all fields.');
            return;
        }

        const formData = {
            username: username,
            password: password
        };

        fetch('`+action_url+`', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData)
        })
        .then(response => {
            if (response.status === 200) {
                alert('Login successful!');
                window.location.href = '/'; // Redirect to dashboard
            } else {
                return response.json();
            }
        })
        .then(data => {
            if (data && data.error) {
                showError(data.error); // Display error message if any
            }
        })
        .catch(error => {
            showError('An error occurred. Please try again.');
            console.error('Error during login:', error);
        });
    });

    function showError(message) {
        errorMessage.style.display = 'block';
        errorMessage.textContent = message;
    }
    </script>
</body>
</html>
	`)

	return html
}

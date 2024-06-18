import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';
import { useEffect, useState } from 'react';

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [username, setUsername] = useState("");

    const handleLogin = async (username: string, password: string) => {
        try {
            const response = await fetch('http://localhost:3000/api/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ username, password }),
                credentials: 'include'  // This is to include cookies in the request
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            console.log('Login successful', data);
            setIsLoggedIn(true);
            setUsername(username);
        } catch (error) {
            console.error('Login failed', error);
        }
    };

    const handleLogout = async () => {
        try {
            const response = await fetch('http://localhost:3000/api/auth/logout', {
                method: 'GET',
                credentials: 'include'  // This is to include cookies in the request
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            console.log('Logout successful', data);
            setIsLoggedIn(false);
            setUsername("");
        } catch (error) {
            console.error('Logout failed', error);
        }
    };

    useEffect(() => {
        const checkLoginStatus = async () => {
            try {
                const response = await fetch('http://localhost:8080/', {
                    credentials: 'include'  // This is to include cookies in the request
                });
                const data = await response.json();
                if (data.isLoggedIn) {
                    setIsLoggedIn(true);
                    setUsername(data.username);
                }
            } catch (error) {
                console.error('Failed to check login status', error);
            }
        };
        checkLoginStatus();
    }, []);

    return (
        <>
            <div>
                <a href="https://vitejs.dev" target="_blank">
                    <img src={viteLogo} className="logo" alt="Vite logo" />
                </a>
                <a href="https://react.dev" target="_blank">
                    <img src={reactLogo} className="logo react" alt="React logo" />
                </a>
            </div>
            <h1>Vite + React</h1>
            <div>
                {isLoggedIn ? (
                    <>
                        <h2>Welcome back, {username}!</h2>
                        <button onClick={handleLogout}>Logout</button>
                    </>
                ) : (
                    <LoginForm onLogin={handleLogin} />
                )}
            </div>
            <p className="read-the-docs">
                Click on the Vite and React logos to learn more
            </p>
        </>
    );
}

const LoginForm = ({ onLogin }: { onLogin: (username: string, password: string) => void }) => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        onLogin(username, password);
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>
                    Username:
                    <input
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </label>
            </div>
            <div>
                <label>
                    Password:
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </label>
            </div>
            <button type="submit">Login</button>
        </form>
    );
};

export default App;

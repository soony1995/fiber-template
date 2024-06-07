import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';

function App() {
    const handleLogin = (provider: string) => {
        let url = "";
        switch(provider) {
            case 'google':
                url = "http://localhost:3000/api/oauth/google/login";
                break;
            case 'kakao':
                url = "http://localhost:3000/api/oauth/kakao/login";
                break;
            default:
                console.error('Unsupported provider');
                return;
        }
        window.location.href = url;
    };

    const handleLogout = () => {
        fetch("http://localhost:3000/api/oauth/google/logout", {
            credentials: 'include'  // This is to include cookies in the request
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                console.log('Logout successful', data);
                window.location.href = '/';
            })
            .catch(error => console.error('Logout failed', error));
    };

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
                <button onClick={() => handleLogin('google')}>Login with Google</button>
                <button onClick={handleLogout}>Logout</button>
            </div>
            <div>
                <button onClick={() => handleLogin('kakao')}>Login with Kakao</button>
                <button onClick={handleLogout}>Logout</button>
            </div>
            <p className="read-the-docs">
                Click on the Vite and React logos to learn more
            </p>
        </>
    );
}

export default App;

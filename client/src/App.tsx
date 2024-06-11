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

    const handleLogout = async (provider: string) => {
        try {
            const baseUrl = "http://localhost:3000/api/oauth/";
            const supportedProviders = ['google', 'kakao'];

            if (!supportedProviders.includes(provider)) {
                console.error('Unsupported provider');
                return;
            }

            const url = `${baseUrl}${provider}/logout`;

            const response = await fetch(url, {
                credentials: 'include'  // This is to include cookies in the request
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            console.log('Logout successful', data);
            window.location.href = '/';
        } catch (error) {
            console.error('Logout failed', error);
        }
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
                <button onClick={() => handleLogout('google')}>Logout</button>
            </div>
            <div>
                <button onClick={() => handleLogin('kakao')}>Login with Kakao</button>
                <button onClick={() => handleLogout('kakao')}>Logout</button>
            </div>
            <p className="read-the-docs">
                Click on the Vite and React logos to learn more
            </p>
        </>
    );
}

export default App;

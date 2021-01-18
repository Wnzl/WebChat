import logo from './logo.svg';
import './App.css';
import PingComponent from './PingComponent';
import Button from '@material-ui/core/Button';
import Login from "./components/Login/Login";
import LoginForm from "./components/LoginForm/LoginForm";


function App() {
  return (
    <div className="App">
    <LoginForm />
        <PingComponent />
    </div>
  );
}

export default App;

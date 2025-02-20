import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
  useLocation, // Importe o useLocation
} from "react-router-dom";
import logo from "./logo.svg";
import "./App.css";
import LoginForm from "./Components/LoginForm";
import Dashboard from "./Components/Dashboard";
import CreateUserForm from "./Components/CreateUserForm";

function App() {
  const location = useLocation();

  return (
    <div className="App">
      <header className="App-header">
        {/* Renderiza o logo apenas se a rota não for /home */}
        {location.pathname !== "/home" && (
          <img src={logo} className="App-logo" alt="logo" />
        )}
        <Routes>
          {/* Rota para a página de login */}
          <Route path="/login" element={<LoginForm />} />

          {/* Rota para criar um novo usuário */}
          <Route path="/create-account" element={<CreateUserForm />} />

          {/* Página inicial com redirecionamento baseado na autenticação */}
          <Route path="/" element={<Navigate to={"/login"} />} />

          {/* Rota para a página principal (protegida) */}
          <Route path="/home" element={<Dashboard />} />
        </Routes>
      </header>
    </div>
  );
}

function Root() {
  return (
    <Router>
      <App />
    </Router>
  );
}

export default Root;
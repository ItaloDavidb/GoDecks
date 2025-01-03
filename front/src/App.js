import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import logo from "./logo.svg";
import "./App.css";
import LoginForm from "./Components/LoginForm";
import HomePage from "./Components/HomePage";
import CreateUserForm from "./Components/CreateUserForm";
function App() {
  const isAuthenticated = localStorage.getItem("token");

  return (
    <Router>
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <Routes>
            {/* Rota para a página de login */}
            <Route path="/login" element={<LoginForm />} />

            {/* Rota para criar um novo usuário */}
            <Route path="/create-account" element={<CreateUserForm />} />

            {/* Página inicial com redirecionamento baseado na autenticação */}
            <Route
              path="/"
              element={<Navigate to={isAuthenticated ? "/home" : "/login"} />}
            />

            {/* Rota para a página principal (protegida) */}
            <Route
              path="/home"
              element={
                isAuthenticated ? <HomePage /> : <Navigate to="/login" />
              }
            />
          </Routes>
        </header>
      </div>
    </Router>
  );
}

export default App;

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

            {/* Rota para a página principal (após login), protegida */}
            <Route
              path="/home"
              element={
                isAuthenticated ? <HomePage /> : <Navigate to="/login" />
              }
            />

            {/* Página inicial (ou redirecionamento para login se não autenticado) */}
            <Route
              path="/home"
              element={<Navigate to={isAuthenticated ? "/home" : "/login"} />}
            />
            {/* Página inicial */}
            <Route
              path="/"
              element={<Navigate to={isAuthenticated ? "/home" : "/login"} />}
            />
          </Routes>
        </header>
      </div>
    </Router>
  );
}

export default App;

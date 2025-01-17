import React, { useState } from "react";
import "./LoginForm.css";
import { useNavigate } from "react-router-dom";
function LoginForm() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const handleCreateAccountClick = () => {
    navigate("/create-account");
  };
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    setLoading(true);
    try {
      const response = await fetch("http://localhost:8080/api/Login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },

        body: JSON.stringify({ username, password }),
      });

      if (response.ok) {
        const data = await response.json();
        localStorage.setItem("token", data.token);
        alert("Login bem-sucedido!");
        navigate("/home");
      } else {
        setError("Nome de usuário ou senha incorretos");
      }
    } catch (error) {
      setError("Erro ao tentar realizar o login");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-form">
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="username">Nome de usuário:</label>
          <input
            type="text"
            id="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="Digite seu nome de usuário"
          />
        </div>
        <div>
          <label htmlFor="password">Senha:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Digite sua senha"
          />
        </div>
        {error && <div style={{ color: "red" }}>{error}</div>}{" "}
        <div class="login-form-button">
          <button type="submit" disabled={loading} className="login">
            {loading ? "Carregando..." : "Entrar"}
          </button>
          <button
            type="button"
            className="createAccount"
            onClick={handleCreateAccountClick}
          >
            Criar Usuário
          </button>
        </div>
      </form>
    </div>
  );
}

export default LoginForm;

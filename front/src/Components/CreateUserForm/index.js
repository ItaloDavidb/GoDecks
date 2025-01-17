import React, { useState } from "react";
import "./CreateUserForm.css";
import { useNavigate } from "react-router-dom";
function CreateUserForm() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const handleTurnBack = () => {
    navigate("/login");
  };
  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    setLoading(true);
    try {
      const response = await fetch("http://localhost:8080/api/Users", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },

        body: JSON.stringify({ username, password, email }),
      });

      if (response.ok) {
        const data = await response.json();
        localStorage.setItem("token", data.token);
        alert("Usuário criado com sucesso!");
        navigate("/login");
      } else {
        setError("Houve algum problema ao criar este usuário tente novamente!");
      }
    } catch (error) {
      setError("Erro ao tentar criar usuario");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="createUser-form">
      <h2>Criar Usuário</h2>
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
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Digite seu Email"
          />
        </div>
        {error && <div style={{ color: "red" }}>{error}</div>}{" "}
        <div class="createUser-form-button">
          <button type="submit" disabled={loading} className="createUser">
            {loading ? "Carregando..." : "Criar Usuário"}
          </button>
          <button type="button" className="turnBack" onClick={handleTurnBack}>
            Voltar
          </button>
        </div>
      </form>
    </div>
  );
}

export default CreateUserForm;

import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const useAuth = () => {
  const navigate = useNavigate();
  const baseUrl = "http://localhost:8080";

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      console.log("Token não encontrado, redirecionando para login...");
      navigate("/login");
    } else {
      console.log("Token encontrado, validando...");

      fetch(`${baseUrl}/api/Auth`, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
        .then((response) => {
          if (!response.ok) {
            console.log(
              "Token inválido ou expirado, redirecionando para login..."
            );
            navigate("/login");
          } else {
            console.log("Token válido, usuário autenticado.");
          }
        })
        .catch((error) => {
          console.log("Erro na requisição de autenticação:", error);
          navigate("/login");
        });
    }
  }, [navigate]);
};

export default useAuth;

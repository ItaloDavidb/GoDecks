import React, { useEffect } from "react";

const HomePage = () => {
  useEffect(() => {
    const token = localStorage.getItem("token");

    if (!token) {
      window.location.href = "/";
    }
  }, []);

  return (
    <div>
      <h2>Bem-vindo à Home Page</h2>
    </div>
  );
};

export default HomePage;

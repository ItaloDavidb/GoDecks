import React from "react";
import useAuth from "../../Hooks/useAuth";

const HomePage = () => {
  useAuth();
  return (
    <div>
      <h2>Bem-vindo à Home Page</h2>
    </div>
  );
};

export default HomePage;

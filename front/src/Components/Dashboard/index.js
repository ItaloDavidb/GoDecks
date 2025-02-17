import React from "react";
import useAuth from "../../Hooks/useAuth";
import './Dashboard.css';

const Dashboard = () => {
  return (
    <div className="dashboard">
      <h1>Dashboard</h1> {/* Título igual à imagem */}
      
      {/* Seção Decks */}
      <div className="dashboard-section">
        <h2>My Decks</h2>
        <p>You have 8 decks.</p>
        <button>View Decks</button>
      </div>

      {/* Seção Cards */}
      <div className="dashboard-section">
        <h2>My Cards</h2>
        <p>You have 120 cards.</p>
        <button>View Cards</button>
      </div>

      {/* Seção Profile */}
      <div className="dashboard-section">
        <h2>Profile</h2>
        <p>Username: JohnDoe</p>
        <button>View Profile</button>
      </div>

      {/* Barra de Pesquisa */}
      <div className="search-section">
        <input type="text" placeholder="Search for cards and decks" />
      </div>
    </div>
  );
};

export default Dashboard;
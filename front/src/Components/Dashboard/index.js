import React from "react";
import useAuth from "../../Hooks/useAuth";
import "./Dashboard.css";
import { ImSearch } from "react-icons/im";
import "chart.js/auto";

const Dashboard = () => {
  const cardCount = 120;
  const deckCount = 8;

  return (
    <div className="dashboard">
      <h1>Dashboard</h1>
      <div className="search-section">
        <input
          type="text"
          placeholder="Search Cards or Decks..."
          className="search-input"
        />
        <button className="search-button decks">
          <ImSearch />
        </button>
      </div>
      <div className="dashboard-actions">
        <div className="dashboard-actions-left">
          <button className="dashboard-button">Cadastrar Carta</button>
          <button className="dashboard-button">Cadastrar Deck</button>
        </div>
        <div className="dashboard-actions-right">
          <button className="dashboard-button">Excluir Deck</button>
          <button className="dashboard-button">Excluir Carta</button>
        </div>
      </div>

      <div className="dashboard-row">
        <div className="dashboard-section" id="decks">
          <h2>My Decks</h2>
          <p>You have {deckCount} decks.</p>
          <button>View Decks</button>
        </div>

        <div className="dashboard-section" id="cards">
          <h2>My Cards</h2>
          <p>You have {cardCount} cards.</p>
          <button>View Cards</button>
        </div>

        <div className="dashboard-section" id="profile">
          <h2>Profile</h2>
          <p>Username: JohnDoe</p>
          <button>View Profile</button>
        </div>
      </div>

      <div className="dashboard-lists">
        <div className="list-section">
          <h3>Últimos Decks Criados</h3>
          <ul>
            <li>Deck 1 - Grass Power</li>
            <li>Deck 2 - Fire Fury</li>
            <li>Deck 3 - Psychic Control</li>
          </ul>
        </div>

        <div className="list-section">
          <h3>Cartas Mais Usadas</h3>
          <ul>
            <li>Rowlet (x5)</li>
            <li>Rare Candy (x4)</li>
            <li>Ultra Ball (x3)</li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;

import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import Professores from './pages/professores';
import Turmas from './pages/turmas';

function App() {
  return (
    <Router>
      <div className="App">
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
          <div className="container-fluid">
            <Link className="navbar-brand" to="/">Cadastro-Escola</Link>
            <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
              <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse" id="navbarNav">
              <ul className="navbar-nav">
                <li className="nav-item">
                  <Link className="nav-link" to="/professores">Professores</Link>
                </li>
                <li className="nav-item">
                  <Link className="nav-link" to="/turmas">Turmas</Link>
                </li>
              </ul>
            </div>
          </div>
        </nav>

        <div className="container mt-4">
          <Routes>
            <Route path="/professores" element={<Professores />} />
            <Route path="/turmas" element={<Turmas />} />
            <Route path="/" element={<h2>Bem-vindo ao Cadastro-Escola</h2>} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;
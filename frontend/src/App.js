import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Professores from './pages/Professores';
import Turmas from './pages/Turmas';

function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/professores" element={<Professores />} />
          <Route path="/turmas" element={<Turmas />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;